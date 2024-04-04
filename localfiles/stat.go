package localfiles

import (
	"os"
	"postgres-backup-service/env"
	"postgres-backup-service/logger"
	"time"
)

var log = logger.GetLogger()

func countFiles() (int, error) {
	files, err := os.ReadDir(env.ENVIRONMENT.BACKUP_DIR)

	if err != nil {
		return 0, err
	}

	return len(files), nil
}

func findOldestFile() (oldestFile os.FileInfo, err error) {
	files, err := os.ReadDir(env.ENVIRONMENT.BACKUP_DIR)

	if err != nil {
		return nil, err
	}

	t := time.Now()

	for _, file := range files {
		if !file.IsDir() {
			fileInf, err := file.Info()

			if err != nil {
				return nil, err
			}

			modTime := fileInf.ModTime()
			if modTime.Before(t) {
				t = modTime
				oldestFile = fileInf
			}
		}
	}

	return oldestFile, nil
}
