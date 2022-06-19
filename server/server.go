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
	f, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return
	}
	defer func() {
		f.Close()
	}()
	log.SetOutput(f)

	defer c.Close()
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
		output, err := exec.Command("runc", args...).Output()
		if err != nil {
			log.Println(output)
			log.Fatal(err)
		}
	}
	log.Println("Handle finished.")
}
