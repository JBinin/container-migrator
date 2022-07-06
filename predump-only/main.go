package predump_only

import (
	"fmt"
	"github.com/JBinin/container-migrator/client"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

func getSize(sourcePath string) (int, error) {
	if output, err := exec.Command("du", "-s", sourcePath).Output(); err != nil {
		log.Println(output)
		return 0, err
	} else {
		size, _ := strconv.Atoi(strings.Split(string(output), "\t")[0])
		return size, nil
	}
}

func TestDump(containerID string, checkpointPath string, channel *chan int) error {
	defer killContainer(containerID)

	log.Println(checkpointPath)
	os.RemoveAll(checkpointPath)
	os.MkdirAll(checkpointPath, os.ModePerm)

	oldPath, _ := os.Getwd()
	os.Chdir(checkpointPath)
	defer os.Chdir(oldPath)

	netSpeed := 1e9
	maxIteration := 10
	dumpTime := make([]float64, maxIteration)
	dumpSize := make([]int, maxIteration)
	xferTime := make([]float64, maxIteration)
	dedupFactor := make([]float64, maxIteration, 1)
	defer printPreInfo(dumpTime, dumpSize, xferTime)

	for i := 0; i < maxIteration; i += 1 {
		preTime, _ := client.PreDump(containerID, i)
		dumpTime[i] = preTime
		size, _ := getSize(path.Join(checkpointPath, fmt.Sprintf("checkpoint%03d", i)))
		dumpSize[i] = size
		log.Printf("Checkpoint dump index: %03d", i)
		timeSleep := float64(size*8*1024*1000) / netSpeed
		timeSleep = timeSleep * dedupFactor[i]
		xferTime[i] = timeSleep
		time.Sleep(time.Duration(int64(timeSleep)) * time.Millisecond)
	}
	*channel <- 1
	return nil
}

func killContainer(containerID string) error {
	cmd := exec.Command("runc", "kill", containerID, "9")
	log.Println(cmd.String())
	return cmd.Start()
}

func printPreInfo(preTime []float64, preSize []int, xferTime []float64) {
	for i, t := range preTime {
		log.Println(i, ":\t", t, "s\t", preSize[i], "KB\t", xferTime[i])
	}
}
