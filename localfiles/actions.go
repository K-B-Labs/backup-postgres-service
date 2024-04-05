package localfiles

import (
	"os"
	"postgres-backup-service/env"
)

func DeleteOldestFile() {
	oldestFile, err := findOldestFile()

	if err != nil {
		log.Error("Error finding oldest file")
	}

	log.Info("Deleting oldest file: " + env.ENVIRONMENT.BACKUP_DIR + "/" + oldestFile.Name())
	err = os.Remove(env.ENVIRONMENT.BACKUP_DIR + "/" + oldestFile.Name())

	if err != nil {
		log.Error("Error deleting oldest file")
	}
}
