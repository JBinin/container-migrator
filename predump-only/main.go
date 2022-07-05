package predump_only

import (
	"github.com/JBinin/container-migrator/client"
	"log"
	"os"
	"os/exec"
	"time"
)

func TestDump(containerID string, checkpointPath string, channel *chan int) error {
	defer killContainer(containerID)

	log.Println(checkpointPath)
	os.RemoveAll(checkpointPath)
	os.MkdirAll(checkpointPath, os.ModePerm)

	oldPath, _ := os.Getwd()
	os.Chdir(checkpointPath)
	defer os.Chdir(oldPath)

	timeInv := 100
	maxIteration := 1 * 1000 / 100
	dumpTime := make([]float64, maxIteration)
	defer printPreTime(dumpTime)

	for i := 0; i < maxIteration; i += 1 {
		preTime, _ := client.PreDump(containerID, i)
		dumpTime[i] = preTime
		log.Println("Checkpoint dump index: ", i)
		time.Sleep(time.Duration(timeInv) * time.Millisecond)
	}
	*channel <- 1
	return nil
}

func killContainer(containerID string) error {
	cmd := exec.Command("runc", "kill", containerID, "9")
	log.Println(cmd.String())
	return cmd.Start()
}

func printPreTime(preTime []float64) {
	for i, t := range preTime {
		log.Println(i, ":\t", t, "s")
	}
}