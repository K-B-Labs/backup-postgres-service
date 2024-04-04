package db

import (
	"fmt"
	"os"
	"os/exec"
	"postgres-backup-service/env"
	"time"
)

func Dump() {
	dumpFile := fmt.Sprintf("backups/dump_%s.sql", time.Now().Format("2006-01-02_15-04-05"))

	cmd := exec.Command("pg_dump", env.ENVIRONMENT.POSTGRES_URL, "-f", dumpFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("Backup created: %s\n", dumpFile)
}
