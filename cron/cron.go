package cron

import (
	"fmt"
	"postgres-backup-service/db"
	"postgres-backup-service/env"

	"github.com/robfig/cron"
)

func StartDumpCronJob() {
	fmt.Println("Starting cron service")

	c := cron.New()
	c.AddFunc(env.ENVIRONMENT.BACKUP_CRON, db.Dump)
	c.Start()
}
