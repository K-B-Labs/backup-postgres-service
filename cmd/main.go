package main

import (
	"postgres-backup-service/db"
	"postgres-backup-service/env"
	"time"
)

func main() {
	env.LoadEnvironment()
	db.Dump()

	for {
		time.Sleep(time.Second)
	}
}
