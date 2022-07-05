package predump_only

import (
	"github.com/JBinin/container-migrator/client"
	"log"
	"os"
	"time"
)

func TestDump(containerID string, checkpointPath string) error {
	oldPath, _ := os.Getwd()
	defer os.Chdir(oldPath)

	if err := os.Chdir(checkpointPath); err != nil {
		log.Println("Failed to chdir to ", checkpointPath)
		log.Println("Use the current work dir")
	}

	timeInv := 100
	maxIteration := 1 * 1000 / 100
	for i := 0; i < 10; i += maxIteration {
		client.PreDump(containerID, i)
		time.Sleep(time.Duration(timeInv) * time.Millisecond)
	}
	return nil
}
