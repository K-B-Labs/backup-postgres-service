package cron

import (
	"postgres-backup-service/env"
	"postgres-backup-service/logger"

	"github.com/robfig/cron"
)

var log = logger.GetLogger()

func StartDumpCronJob() {
	log.Info("Starting cron service")

	c := cron.New()
	c.AddFunc(env.ENVIRONMENT.BACKUP_CRON, func() { dumpLocal() })
	c.Start()
}
