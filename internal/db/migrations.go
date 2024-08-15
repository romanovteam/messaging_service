package db

import (
	"context"
	"log"
	"os"
	"path/filepath"
)

func MigrateDB(migrationFilePath string) {
	migration, err := os.ReadFile(filepath.Clean(migrationFilePath))
	if err != nil {
		log.Fatalf("Failed to read migration file: %v\n", err)
	}

	_, err = Pool.Exec(context.Background(), string(migration))
	if err != nil {
		log.Fatalf("Failed to execute migration: %v\n", err)
	} else {
		log.Println("Migration executed successfully")
	}
}
