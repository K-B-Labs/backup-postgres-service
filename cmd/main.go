package main

import (
	"postgres-backup-service/cron"
	"postgres-backup-service/env"
	"time"
)

func main() {
	env.LoadEnvironment()
	cron.StartCron()

	for {
		time.Sleep(time.Second)
	}
}
