package server

import (
	"log"
	"net"
	"os/exec"
	"path"
	"strings"
)

func ListenAndServe(migratedContainerDir string) {
	conn, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
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
	var buf [512]byte
	n, err := c.Read(buf[:])
	if err != nil {
		log.Fatal(err)
	}
	receive := string(buf[:n])
	if receive == "DestPath" {
		c.Write([]byte(migratedContainerDir))
	}
	if strings.HasPrefix(receive, "restore") {
		cmd := strings.Split(receive, ":")
		imagePath := path.Join(migratedContainerDir, "image")
		args := []string{cmd[0], "--image-path", imagePath, cmd[1]}
		output, err := exec.Command("runc", args...).Output()
		if err != nil {
			log.Println(output)
			log.Fatal(err)
		}
	}
}
