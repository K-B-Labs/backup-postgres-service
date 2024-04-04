package cron

import (
	"postgres-backup-service/db"
	"postgres-backup-service/localfiles"
)

func dumpLocal() {
	db.Dump()
	localfiles.DeleteOldestFile()
}
