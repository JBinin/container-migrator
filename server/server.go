package server

import (
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"strings"
)

func ListenAndServe(migratedContainerDir string) {
	os.RemoveAll(migratedContainerDir)
	os.MkdirAll(migratedContainerDir, os.ModePerm)

	conn, err := net.Listen("tcp", ":8001")
	defer conn.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
	for {
		acc, err1 := conn.Accept()
		if err1 != nil {
			log.Println(err1.Error())
			break
		}
		go handleConn(acc, migratedContainerDir)
	}
}

func handleConn(c net.Conn, migratedContainerDir string) {
	defer c.Close()

	f, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	defer f.Close()
	log.SetOutput(f)

	var buf [512]byte
	n, err1 := c.Read(buf[:])
	if err1 != nil {
		log.Fatal(err1)
	}
	receive := string(buf[:n])
	log.Println(receive)
	if receive == "DestPath" {
		c.Write([]byte(migratedContainerDir))
	}
	n, err1 = c.Read(buf[:])
	if err1 != nil {
		log.Fatal(err1)
	}
	receive = string(buf[:n])
	log.Println(receive)
	if strings.HasPrefix(receive, "restore") {
		cmd := strings.Split(receive, ":")
		imagePath := path.Join(migratedContainerDir, "image")
		args := []string{cmd[0], "--image-path", imagePath, cmd[1]}

		oldDir, _ := os.Getwd()
		os.Chdir(migratedContainerDir)
		exec.Command("runc", args...).Output()
		os.Chdir(oldDir)
		c.Write([]byte("started"))
	}
	log.Println("Handle finished.")
}
