package client

import (
	"log"
	"net"
	"os"
	"os/exec"
	"path"
)

func preDump(containerId string) error {
	args := []string{"checkpoint", "--pre-dump", "--image-path", "parent", containerId}
	_, err := exec.Command("runc", args...).Output()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func dump(containerID string) error {
	args := []string{"checkpoint", "--image-path", "image", "--parent-path", "../parent", containerID}
	_, err := exec.Command("runc", args...).Output()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func transfer(sourcePath string, destination string, destPath string, info string) error {
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
	return nil
}

func PostCopy(containerID string, destination string) error {
	// todo
	return nil
}
