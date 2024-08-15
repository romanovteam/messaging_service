package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
	"strings"
)

var Pool *pgxpool.Pool

func InitDB(connString string) {
	var err error
	Pool, err = pgxpool.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Запускаем все миграции
	RunMigrations("internal/db/migrations")
}

func CloseDB() {
	Pool.Close()
}

func RunMigrations(migrationsDir string) {
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		log.Fatalf("Failed to read migrations directory: %v\n", err)
	}

	var migrationFiles []string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") && file.Name() != ".migrations" {
			migrationFiles = append(migrationFiles, file.Name())
		}
	}

	sort.Strings(migrationFiles)

	for _, fileName := range migrationFiles {
		version := strings.TrimSuffix(fileName, filepath.Ext(fileName))
		filePath := filepath.Join(migrationsDir, fileName)
		MigrateDB(version, filePath)
	}
}
