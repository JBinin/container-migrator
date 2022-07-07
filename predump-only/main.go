package predump_only

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

func DumpRuning(containerID string, index int) (dumpTime float64, err error) {
	start := time.Now()
	args := []string{
		"checkpoint",
		"--tcp-established",
		"--leave-running",
		"--image-path",
		fmt.Sprintf("checkpoint%03d", index),
	}
	//if index != 0 {
	//	args = append(args, "--parent-path", fmt.Sprintf("../checkpoint%03d", index-1))
	//}
	args = append(args, containerID)
	cmd := exec.Command("runc", args...)
	var b bytes.Buffer
	cmd.Stderr = &b
	if output, err := cmd.Output(); err != nil {
		log.Println(output)
		log.Println(b.String())
		log.Println(cmd.String())
		return 0, err
	}
	elapsed := time.Since(start)
	//log.Println("The dump time is ", elapsed)
	return elapsed.Seconds(), nil
}

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
	maxIteration := 2
	dumpTime := make([]float64, maxIteration)
	dumpSize := make([]int, maxIteration)
	xferTime := make([]float64, maxIteration)
	dedupFactor := make([]float64, maxIteration)
	for i := 0; i < maxIteration; i += 1 {
		dedupFactor[i] = 1
	}
	//dedupFactor[0] = 0.72
	defer printPreInfo(dumpTime, dumpSize, xferTime, dedupFactor)
	last := false
	for i := 0; i < maxIteration; i += 1 {
		if (i != 0 && dumpTime[i-1]+xferTime[i-1] < 1) || i == maxIteration-1 {
			dumptime, _ := DumpRuning(containerID, i)
			dumpTime[i] = dumptime
			size, _ := getSize(path.Join(checkpointPath, fmt.Sprintf("checkpoint%03d", i)))
			dumpSize[i] = size
			last = true
		} else {
			preTime, _ := DumpRuning(containerID, i)
			dumpTime[i] = preTime
			size, _ := getSize(path.Join(checkpointPath, fmt.Sprintf("checkpoint%03d", i)))
			dumpSize[i] = size
		}

		log.Printf("Checkpoint dump index: %03d", i)
		timeSleep := float64(dumpSize[i]*8*1024*1000) / netSpeed
		xferTime[i] = timeSleep / 1000
		time.Sleep(time.Duration(int64(timeSleep*dedupFactor[i])) * time.Millisecond)
		if last {
			break
		}
	}
	*channel <- 1
	return nil
}

func killContainer(containerID string) error {
	cmd := exec.Command("runc", "kill", containerID, "9")
	log.Println(cmd.String())
	return cmd.Start()
}

func printPreInfo(preTime []float64, preSize []int, xferTime []float64, dedepFactor []float64) {
	for i, t := range preTime {
		log.Println(i, ":\t", t, "s\t", float64(preSize[i])/1024, "MB\t", xferTime[i], "s\t", xferTime[i]*dedepFactor[i], "s")
	}
}
