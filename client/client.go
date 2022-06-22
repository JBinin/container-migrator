package client

import (
	"errors"
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

func transfer(sourcePath string, destIP string, destPath string) (float64, int, error) {
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
	rsyncOpts := "-aqz --bwlimit=500000"
	dest := destIP + ":" + destPath
	if output, err := exec.Command("rsync", rsyncOpts, sourcePath, dest).Output(); err != nil {
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
			if _, size, err := transfer(preDumpPath, destIP, destPath); err != nil {
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

	log.Println("Transfer config.json")
	if _, _, err := transfer(path.Join(othersPath, "config.json"), destIP, destPath); err != nil {
		log.Println("Transfer config failed")
		return err
	}
	log.Println("Transfer rootfs")
	if _, _, err := transfer(path.Join(othersPath, "rootfs"), destIP, destPath); err != nil {
		log.Println("Transfer rootfs failed")
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
			if _, _, err := transfer(dumpPath, destIP, destPath); err != nil {
				log.Println("Transfer dump data failed")
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
