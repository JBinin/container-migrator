package predump_only

import (
	"github.com/JBinin/container-migrator/client"
	"log"
	"os"
	"os/exec"
	"time"
)

func TestDump(containerID string, checkpointPath string) error {
	defer killContainer(containerID)

	os.RemoveAll(checkpointPath)
	os.MkdirAll(checkpointPath, os.ModePerm)

	oldPath, _ := os.Getwd()
	defer os.Chdir(oldPath)

	timeInv := 100
	maxIteration := 1 * 1000 / 100
	for i := 0; i < 10; i += maxIteration {
		client.PreDump(containerID, i)
		log.Println("Checkpoint dump index: ", i)
		time.Sleep(time.Duration(timeInv) * time.Millisecond)
	}
	return nil
}

func killContainer(containerID string) error {
	cmd := exec.Command("runc", "kill", containerID, "9")
	log.Println(cmd.String())
	return cmd.Start()
}
