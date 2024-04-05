package cron

import (
	"postgres-backup-service/db"
	"postgres-backup-service/env"
	"postgres-backup-service/localfiles"
	"strconv"
)

func dumpLocal() {
	db.Dump()

	numOfBackups, err := localfiles.CountFiles()

	if err != nil {
		log.Error("Error counting backups")
		return
	}

	if numOfBackups < env.ENVIRONMENT.MAX_BACKUPS {
		log.Info("Not deleting old backups. Current number of backups: " + strconv.Itoa(numOfBackups) + "; Threshold: " + strconv.Itoa(env.ENVIRONMENT.MAX_BACKUPS))
		return
	}

	localfiles.DeleteOldestFile()
}
