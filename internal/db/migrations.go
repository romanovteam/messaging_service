package db

import (
	"bufio"
	"context"
	"log"
	"os"
	"path/filepath"
)

func MigrateDB(version string, migrationFilePath string) {
	migrationsFilePath := filepath.Join("internal", "db", "migrations", ".migrations")

	// Проверяем, существует ли файл миграций
	if _, err := os.Stat(migrationsFilePath); os.IsNotExist(err) {
		_, err := os.Create(migrationsFilePath)
		if err != nil {
			log.Fatalf("Failed to create migrations file: %v\n", err)
		}
	}

	// Проверяем, была ли уже выполнена миграция
	done, err := checkMigrationDone(migrationsFilePath, version)
	if err != nil {
		log.Fatalf("Failed to check migration file: %v\n", err)
	}

	if done {
		log.Println("Migration already executed:", version)
		return
	}

	// Читаем и выполняем миграцию
	migration, err := os.ReadFile(filepath.Clean(migrationFilePath))
	if err != nil {
		log.Fatalf("Failed to read migration file: %v\n", err)
	}

	_, err = Pool.Exec(context.Background(), string(migration))
	if err != nil {
		log.Fatalf("Failed to execute migration: %v\n", err)
	}

	// Записываем факт выполнения миграции
	recordMigrationDone(migrationsFilePath, version)

	log.Println("Migration executed successfully:", version)
}

func checkMigrationDone(migrationsFilePath, version string) (bool, error) {
	file, err := os.Open(migrationsFilePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Проверка, выполнена ли уже миграция
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == version {
			return true, nil
		}
	}

	return false, scanner.Err()
}

func recordMigrationDone(migrationsFilePath, version string) {
	file, err := os.OpenFile(migrationsFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open migrations file: %v\n", err)
	}
	defer file.Close()

	if _, err := file.WriteString(version + "\n"); err != nil {
		log.Fatalf("Failed to record migration version: %v\n", err)
	}
}
