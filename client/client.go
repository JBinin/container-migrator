package client

import (
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

func transfer(sourcePath string, destination string, destPath string, info string) (speed float64, size int, err error) {
	start := time.Now()

	if output, err := exec.Command("du", "-s", sourcePath).Output(); err != nil {
		log.Fatal(err)
		return 0, 0, err
	} else {
		size, _ = strconv.Atoi(strings.Split(string(output), "\t")[0])
		log.Println(info)
		log.Println("Transfer size: ", string(output))
	}
	rsyncOpts := "-aqz"
	dest := destination + ":" + destPath
	if _, err3 := exec.Command("rsync", rsyncOpts, sourcePath, dest).Output(); err3 != nil {
		log.Fatal(err3)
		return 0, 0, err3
	}
	elapsed := time.Since(start)
	log.Println("The transfer time is ", elapsed)
	return float64(size) / elapsed.Seconds(), size, nil
}

func PreCopy(containerID string, destination string, othersPath string) error {
	oldDir, _ := os.Getwd()
	basePath := path.Join(oldDir, containerID)
	_ = os.RemoveAll(basePath)
	_ = os.MkdirAll(basePath, os.ModePerm)

	var conn net.Conn
	var err error
	conn, err = net.Dial("tcp", destination+":8001")
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if _, err1 := conn.Write([]byte("DestPath")); err1 != nil {
		log.Fatal(err1)
		return err1
	}
	buf := [512]byte{}
	n, err2 := conn.Read(buf[:])
	if err2 != nil {
		log.Fatal(err2)
		return err2
	}
	destPath := string(buf[:n])

	totalStart := time.Now()

	transfer(path.Join(othersPath, "config.json"), destination, destPath, "config.json")
	transfer(path.Join(othersPath, "rootfs"), destination, destPath, "rootfs")

	err3 := os.Chdir(basePath)
	defer os.Chdir(oldDir)
	if err3 != nil {
		log.Fatal(err3)
		return err3
	}
	var index int
	for i := 0; i < 10; i += 1 {
		log.Println("The ", index, " iteration")
		index = i
		if err4 := preDump(containerID, i); err4 != nil {
			log.Fatal(err4)
			return err4
		} else {
			var D float64
			D = 128 * 1024
			preDumpPath := path.Join(basePath, "checkpoint"+strconv.Itoa(index))
			speed, size, err5 := transfer(preDumpPath, destination, destPath, "preDump data")
			if err5 != nil {
				log.Fatal(err5)
				return err5
			}
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

	start := time.Now()
	if err5 := dump(containerID, index); err2 != nil {
		log.Fatal(err5)
		return err5
	} else {
		dumpPath := path.Join(basePath, "checkpoint")
		transfer(dumpPath, destination, destPath, "dump data")
	}
	_, err6 := conn.Write([]byte("restore:" + containerID))
	if err6 != nil {
		log.Fatal(err6)
		return err6
	}
	var err7 error
	if n, err7 = conn.Read(buf[:]); err7 == nil {
		if string(buf[:n]) == "started" {
			elapsed := time.Since(start)
			log.Println("The downtime is ", elapsed)

			totalElapsed := time.Since(totalStart)
			log.Println("The total migration time is ", totalElapsed)

			return nil
		}
	}
	return err7
}

func PostCopy(containerID string, destination string) error {
	// todo
	return nil
}
