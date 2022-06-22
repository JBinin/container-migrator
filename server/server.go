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
	if err := os.RemoveAll(migratedContainerDir); err != nil {
		log.Println("Failed to remove ", migratedContainerDir)
	}
	if err := os.MkdirAll(migratedContainerDir, os.ModePerm); err != nil {
		log.Println("Failed to mkdir ", migratedContainerDir)
	}

	conn, err := net.Listen("tcp", ":8001")
	defer conn.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}
	for {
		if acc, err := conn.Accept(); err != nil {
			log.Println(err.Error())
			break
		} else {
			go handleConn(acc, migratedContainerDir)
		}
	}
}

func handleConn(c net.Conn, migratedContainerDir string) {
	defer c.Close()
	if f, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm); err != nil {
		log.Println("Failed to open the server.log")
	} else {
		defer f.Close()
		log.SetOutput(f)
	}

	var buf [512]byte
	if n, err := c.Read(buf[:]); err != nil {
		log.Println("Failed to receive DestPath cmd")
		return
	} else {
		receive := string(buf[:n])
		log.Println(receive)
		if receive == "DestPath" {
			if _, err := c.Write([]byte(migratedContainerDir)); err != nil {
				log.Println("Failed to send the migratedContainerDir to client")
				return
			}
		}
	}

	if n, err := c.Read(buf[:]); err != nil {
		log.Println("Failed to receive restore cmd")
		return
	} else {
		receive := string(buf[:n])
		log.Println(receive)
		if strings.HasPrefix(receive, "restore") {
			cmd := strings.Split(receive, ":")
			imagePath := path.Join(migratedContainerDir, "checkpoint")
			args := []string{"restore", "-d", "--image-path", imagePath, cmd[1]}

			oldDir, _ := os.Getwd()
			os.Chdir(migratedContainerDir)
			defer os.Chdir(oldDir)

			if _, err := exec.Command("runc", args...).Output(); err != nil {
				log.Println("Failed to restore the contaier")
				log.Println(err.Error())
				return
			} else {
				if _, err := c.Write([]byte("started")); err != nil {
					log.Println("Failed to send the started message to client ")
					return
				}
			}
		}
	}
	log.Println("Handle finished.")
}
