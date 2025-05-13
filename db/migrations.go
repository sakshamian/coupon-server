package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

func RunMigrations(db *sql.DB, migrationsDir string) error {
	// Get list of migration files
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to read migrations directory: %w", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" && !file.IsDir() {
			// Read the SQL file
			filePath := filepath.Join(migrationsDir, file.Name())
			sqlScript, err := os.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", filePath, err)
			}

			// Execute the SQL script
			if _, err := db.Exec(string(sqlScript)); err != nil {
				return fmt.Errorf("failed to execute script %s: %w", filePath, err)
			}

			fmt.Printf("Successfully executed migration: %s\n", file.Name())
		}
	}

	return nil
}
