package client

import (
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"time"
)

func preDump(containerId string) error {
	start := time.Now()
	args := []string{"checkpoint", "--pre-dump", "--image-path", "parent", containerId}
	_, err := exec.Command("runc", args...).Output()
	if err != nil {
		log.Fatal(err)
		return err
	}
	elapsed := time.Since(start)
	log.Println("The pre dump time is ", elapsed)
	return nil
}

func dump(containerID string) error {
	start := time.Now()
	args := []string{"checkpoint", "--image-path", "image", "--parent-path", "../parent", containerID}
	_, err := exec.Command("runc", args...).Output()
	if err != nil {
		log.Fatal(err)
		return err
	}
	elapsed := time.Since(start)
	log.Println("The dump time is ", elapsed)
	return nil
}

func transfer(sourcePath string, destination string, destPath string, info string) error {
	start := time.Now()
	if output, err := exec.Command("du", "-hs", sourcePath).Output(); err != nil {
		log.Fatal(err)
		return err
	} else {
		log.Println(info)
		log.Println("Transfer size: ", string(output))
	}
	rsyncOpts := "-aqz"
	dest := destination + ":" + destPath
	if _, err3 := exec.Command("rsync", rsyncOpts, sourcePath, dest).Output(); err3 != nil {
		log.Fatal(err3)
		return err3
	}
	elapsed := time.Since(start)
	log.Println("The transfer time is ", elapsed)
	return nil
}

func PreCopy(containerID string, destination string, othersPath string) error {
	oldDir, _ := os.Getwd()
	basePath := path.Join(oldDir, containerID)
	os.RemoveAll(basePath)
	imagePath := path.Join(basePath, "image")
	parentPath := path.Join(basePath, "parent")
	os.MkdirAll(imagePath, os.ModePerm)
	os.MkdirAll(parentPath, os.ModePerm)

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
	if err4 := preDump(containerID); err4 != nil {
		log.Fatal(err4)
		return err4
	} else {
		transfer(parentPath, destination, destPath, "preDump data")
	}

	start := time.Now()
	if err5 := dump(containerID); err2 != nil {
		log.Fatal(err5)
		return err5
	} else {
		transfer(imagePath, destination, destPath, "dump data")
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
