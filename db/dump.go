package db

import (
	"fmt"
	"os"
	"os/exec"
	"postgres-backup-service/env"
	"postgres-backup-service/logger"
	"time"
)

var log = logger.GetLogger()

func Dump() {
	dumpFile := fmt.Sprintf("%s/dump_%s.sql", env.ENVIRONMENT.BACKUP_DIR, time.Now().Format("2006-01-02_15-04-05"))
	log.Info("Starting backup job for file: " + dumpFile)

	cmd := exec.Command("pg_dump", env.ENVIRONMENT.POSTGRES_URL, "-f", dumpFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Error("Error running pg_dump command for file: " + dumpFile)
		os.Exit(1)
	}

	log.Info("Backup job completed for file: " + dumpFile)
}
