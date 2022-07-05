package predump_only

import (
	"github.com/JBinin/container-migrator/client"
	"os"
	"time"
)

func TestDump(containerID string, checkpointPath string) error {
	os.RemoveAll(checkpointPath)
	os.MkdirAll(checkpointPath, os.ModePerm)

	oldPath, _ := os.Getwd()
	defer os.Chdir(oldPath)

	timeInv := 100
	maxIteration := 1 * 1000 / 100
	for i := 0; i < 10; i += maxIteration {
		client.PreDump(containerID, i)
		time.Sleep(time.Duration(timeInv) * time.Millisecond)
	}
	return nil
}
