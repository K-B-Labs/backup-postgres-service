package cron

import (
	"fmt"
	"os"
	"postgres-backup-service/db"
	"postgres-backup-service/env"

	"github.com/adhocore/gronx"
	"github.com/robfig/cron"
)

func StartCron() {
	if !gronx.New().IsValid(env.ENVIRONMENT.BACKUP_CRON) {
		fmt.Printf("Invalid cron expression: %s\n", env.ENVIRONMENT.BACKUP_CRON)
		os.Exit(1)
	}

	c := cron.New()
	c.AddFunc(env.ENVIRONMENT.BACKUP_CRON, db.Dump)
	c.Start()
}
