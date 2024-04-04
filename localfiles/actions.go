package localfiles

import (
	"os"
	"postgres-backup-service/env"
	"strconv"
)

func DeleteOldestFile() {
	numOfBackups, err := countFiles()

	if err != nil {
		log.Error("Error counting backups")
		return
	}

	if numOfBackups < env.ENVIRONMENT.MAX_BACKUPS {
		log.Info("No need to delete old backups yet as threshold is not yet met. Current number of files: " + strconv.Itoa(numOfBackups) + "; Threshold: " + strconv.Itoa(env.ENVIRONMENT.MAX_BACKUPS))
		return
	}

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
