package client

import (
	"errors"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

// T the max expected time of downtime(s)
var T float64 = 1

type transferInfo struct {
	index        int
	data         float64
	preTime      float64
	transferTime float64
}

var Info []transferInfo

func PrintInfo() {
	log.Println("-------------PrintInfo---------------------------------------------")
	log.Println("index\t", "data size(KB)\t", "pre-time(s)\t", "transfer-time(s)\t")
	for _, f := range Info {
		log.Println(f.index, "\t", f.data, "\t", f.preTime)
	}
	log.Println("--------------------------------------------------------------------")
}

func preDump(containerId string, index int) (preTime float64, err error) {
	start := time.Now()
	args := []string{
		"checkpoint",
		"--pre-dump",
		"--tcp-established",
		"--image-path",
		"checkpoint" + strconv.Itoa(index),
	}
	if index != 0 {
		args = append(args, "--parent-path", "../checkpoint"+strconv.Itoa(index-1))
	}
	args = append(args, containerId)
	if output, err := exec.Command("runc", args...).Output(); err != nil {
		log.Println(output)
		return 0, err
	}
	elapsed := time.Since(start)
	//log.Println("The pre-dump index is ", index, " . The pre-dump time is ", elapsed)
	return elapsed.Seconds(), nil
}

func dump(containerID string, index int) (dumpTime float64, err error) {
	start := time.Now()
	args := []string{
		"checkpoint",
		"--tcp-established",
		"--image-path",
		"checkpoint",
		"--parent-path",
		"../checkpoint" + strconv.Itoa(index),
		containerID,
	}
	if output, err := exec.Command("runc", args...).Output(); err != nil {
		log.Println(output)
		return 0, err
	}
	elapsed := time.Since(start)
	//log.Println("The dump time is ", elapsed)
	return elapsed.Seconds(), nil
}

func transfer(sourcePath string, destIP string, destPath string, otherOpts []string) (transferTime float64, size int, err error) {
	if output, err := exec.Command("du", "-s", sourcePath).Output(); err != nil {
		log.Println(output)
		return 0, 0, err
	} else {
		size, _ = strconv.Atoi(strings.Split(string(output), "\t")[0])
		//log.Println("Transfer size: ", size, " KB")
	}
	dest := destIP + ":" + destPath
	rsyncOpts := []string{"-aqz", "--bwlimit=500000", sourcePath, dest}
	if otherOpts != nil {
		rsyncOpts = append(otherOpts, rsyncOpts...)
	}
	start := time.Now()
	if output, err := exec.Command("rsync", rsyncOpts...).Output(); err != nil {
		log.Println(output)
		return 0, size, err
	}
	elapsed := time.Since(start)
	//log.Println("The transfer time is ", elapsed)

	return elapsed.Seconds(), size, nil
}

func iterator(containerID string, basePath string, destIP string, destPath string) (int, error) {
	var index int
	for i := 0; i < 10; i += 1 {
		index = i
		if preTime, err := preDump(containerID, i); err != nil {
			log.Println("The ", index, "iteration pre dump failed ")
			return index, err
		} else {
			D := 128 * 1024.0
			N := 1.25e5
			preDumpPath := path.Join(basePath, "checkpoint"+strconv.Itoa(index))
			if transferTime, size, err := transfer(preDumpPath, destIP, destPath, nil); err != nil {
				log.Println("The ", index, "iteration transfer pre data failed")
				return index, err
			} else {
				log.Println("Disk IO : ", D, " KB/s")
				log.Println("Net speed: ", N, " KB/s")
				S := T * (D * N / (2*N + D))
				log.Println("Expect memory size: ", S)
				log.Println("Real memory size: ", size)
				Info = append(Info, transferInfo{
					index:        index,
					data:         float64(size),
					preTime:      preTime,
					transferTime: transferTime,
				})
				if float64(size) < S {
					break
				}
			}
		}
	}
	return index, nil
}

func syncDir(destPath string, destIP string, othersPath string, otherOpts []string) error {
	dir, err := ioutil.ReadDir(othersPath)
	if err != nil {
		log.Println("Open ", othersPath, " failed")
		return err
	}
	start := time.Now()
	for _, fi := range dir {
		absPath := path.Join(othersPath, fi.Name())
		if _, _, err := transfer(absPath, destIP, destPath, otherOpts); err != nil {
			log.Println("Failed to transfer ", absPath)
			return err
		}
	}
	elapsed := time.Since(start)
	log.Println("Sync time is ", elapsed)
	return nil
}

func PreCopy(containerID string, destIP string, othersPath string) error {
	defer PrintInfo()
	oldDir, _ := os.Getwd()
	basePath := path.Join(oldDir, containerID)
	if err := os.RemoveAll(basePath); err != nil {
		log.Println("Remove ", basePath, " failed")
		return err
	}
	if err := os.MkdirAll(basePath, os.ModePerm); err != nil {
		log.Println("Mkdir ", basePath, " failed")
	}

	var conn net.Conn
	var conErr error
	conn, conErr = net.Dial("tcp", destIP+":8001")
	defer conn.Close()
	if conErr != nil {
		log.Println("Tcp connect failed")
		return conErr
	}

	if _, err := conn.Write([]byte("DestPath")); err != nil {
		log.Println("Get DestPath failed")
		return err
	}

	buf := [512]byte{}
	var destPath string
	if n, err := conn.Read(buf[:]); err != nil {
		log.Println("Get DestPath failed")
		return err
	} else {
		destPath = string(buf[:n])
	}

	totalStart := time.Now()

	if err := syncDir(destPath, destIP, othersPath, nil); err != nil {
		log.Println("Sync dir failed")
		return err
	}

	if err := os.Chdir(basePath); err != nil {
		log.Println("Failed to change the work directory")
		return err
	}
	defer os.Chdir(oldDir)

	if index, err := iterator(containerID, basePath, destIP, destPath); err != nil {
		log.Println("Iterator transfer failed")
		return err
	} else {
		start := time.Now()
		if dumpTime, err := dump(containerID, index); err != nil {
			log.Println("Dump data failed")
			return err
		} else {
			dumpPath := path.Join(basePath, "checkpoint")
			if transferTime, size, err := transfer(dumpPath, destIP, destPath, nil); err != nil {
				log.Println("Transfer dump data failed")
				return err
			} else {
				otherOpts := []string{"--exclude", "rootfs/", "--exclude", "config.json"}
				if err := syncDir(destPath, destIP, othersPath, otherOpts); err != nil {
					log.Println("Sync dir failed")
					return err
				}
				log.Println("---------------------dump------------------------")
				log.Println("dumpTime(s)\t", "data size(KB)\t", "transfer time(s)")
				log.Println(dumpTime, "\t", size, "\t", transferTime, "\t")
				log.Println("-------------------------------------------------")
			}
		}
		if _, err := conn.Write([]byte("restore:" + containerID)); err != nil {
			log.Println("Send restore cmd failed")
			return err
		}
		if n, err := conn.Read(buf[:]); err != nil {
			log.Println("Waiting for restore container in another machine")
			return err
		} else {
			if string(buf[:n]) == "started" {
				elapsed := time.Since(start)
				log.Println("The downtime is ", elapsed)
			} else {
				log.Println("Restore error in remote machine")
				return errors.New("Restore failed ")
			}
		}
	}
	totalElapsed := time.Since(totalStart)
	log.Println("The total migration time is ", totalElapsed)
	return nil
}

func PostCopy(containerID string, destination string) error {
	// todo
	return nil
}
