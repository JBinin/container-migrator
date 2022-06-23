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

func preDump(containerId string, index int) error {
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
		return err
	}
	elapsed := time.Since(start)
	log.Println("The pre-dump index is ", index, " . The pre-dump time is ", elapsed)
	return nil
}

func dump(containerID string, index int) error {
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
		return err
	}
	elapsed := time.Since(start)
	log.Println("The dump time is ", elapsed)
	return nil
}

func transfer(sourcePath string, destIP string, destPath string, otherOpts []string) (float64, int, error) {
	start := time.Now()
	speed := 0.0
	size := 0
	if output, err := exec.Command("du", "-s", sourcePath).Output(); err != nil {
		log.Println(output)
		return speed, size, err
	} else {
		size, _ = strconv.Atoi(strings.Split(string(output), "\t")[0])
		log.Println("Transfer size: ", size, " KB")
	}
	dest := destIP + ":" + destPath
	rsyncOpts := []string{"-aqz", "--bwlimit=500000", sourcePath, dest}
	if otherOpts != nil {
		rsyncOpts = append(otherOpts, rsyncOpts...)
	}
	if output, err := exec.Command("rsync", rsyncOpts...).Output(); err != nil {
		log.Println(output)
		return speed, size, err
	}
	elapsed := time.Since(start)
	log.Println("The transfer time is ", elapsed)
	speed = float64(size) / elapsed.Seconds()
	return speed, size, nil
}

func iterator(containerID string, basePath string, destIP string, destPath string) (int, error) {
	var index int
	for i := 0; i < 10; i += 1 {
		log.Println("The ", index, " iteration")
		index = i
		if err := preDump(containerID, i); err != nil {
			log.Println("Pre dump failed ")
			return index, err
		} else {
			D := 128 * 1024.0
			speed := 500000.0
			preDumpPath := path.Join(basePath, "checkpoint"+strconv.Itoa(index))
			log.Println("Pre dump")
			if _, size, err := transfer(preDumpPath, destIP, destPath, nil); err != nil {
				log.Println("Transfer pre data failed")
				return index, err
			} else {
				log.Println("Disk IO : ", D, " KB/s")
				log.Println("Net speed: ", speed, " KB/s")
				S := T * (D * speed / (2*speed + D))
				log.Println("Expect memory size: ", S)
				log.Println("Real memory size: ", size)
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
	for _, fi := range dir {
		absPath := path.Join(othersPath, fi.Name())
		if _, _, err := transfer(absPath, destIP, destPath, otherOpts); err != nil {
			log.Println("Failed to transfer ", absPath)
			return err
		}
	}
	return nil
}

func PreCopy(containerID string, destIP string, othersPath string) error {
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
		if err := dump(containerID, index); err != nil {
			log.Println("Dump data failed")
			return err
		} else {
			dumpPath := path.Join(basePath, "checkpoint")
			log.Println("Dump data")
			if _, _, err := transfer(dumpPath, destIP, destPath, nil); err != nil {
				log.Println("Transfer dump data failed")
				return err
			}
			otherOpts := []string{"--exclude", "rootfs/", "--exclude", "config.json"}
			if err := syncDir(destPath, destIP, othersPath, otherOpts); err != nil {
				log.Println("Sync dir failed")
				return err
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
