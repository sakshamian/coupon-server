package db

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func Connect() {
// 	var err error
// 	dsn := os.Getenv("DB_DSN")
// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal("Error connecting to the database:", err)
// 	}

// 	// DB.AutoMigrate(&models.Message{})
// 	fmt.Println("Successfully connected to the database!")
// }

// package sdb

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DBUser             string
	DBPassword         string
	DBName             string
	DBHost             string
	DBPort             string
	MaxOpenConnections int
	MaxIdleConnections int
}

var (
	DB *gorm.DB
)

func Connect() {
	var err error

	config := LoadDBConfig()
	// Initialize MySQL client with GORM
	dbParams := "?parseTime=true&charset=utf8mb4&timeout=5s&rejectReadOnly=true"
	fmt.Println(config, ":check")
	dsn := config.DBUser + ":" + config.DBPassword + "@tcp(" + config.DBHost + ":" + config.DBPort + ")/" + config.DBName + dbParams
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB from GORM: %v", err)
	}

	sqlDB.SetMaxOpenConns(config.MaxOpenConnections)
	sqlDB.SetMaxIdleConns(config.MaxIdleConnections)

	cwd, _ := os.Getwd()
	migrationsPath := filepath.Join(cwd, "migrations")
	err = RunMigrations(sqlDB, migrationsPath)
	if err != nil {
		fmt.Println("Failed to migrate tables:", err)
	}
	log.Println("connected to MySQL")
}

func LoadDBConfig() Config {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or there was an error loading it: %v", err)
	}

	maxOpenConns, err := strconv.Atoi(getEnv("MAX_OPEN_CONNECTIONS", "100"))
	if err != nil {
		maxOpenConns = 100
	}
	maxIdleConns, err := strconv.Atoi(getEnv("MAX_IDLE_CONNECTIONS", "100"))
	if err != nil {
		maxIdleConns = 10
	}

	return Config{
		DBUser:             getEnv("DB_USER", "root"),
		DBPassword:         getEnv("DB_PASSWORD", ""),
		DBName:             getEnv("DB_NAME", "mydatabase"),
		DBHost:             getEnv("DB_HOST", "localhost"),
		DBPort:             getEnv("DB_PORT", "3306"),
		MaxOpenConnections: maxOpenConns,
		MaxIdleConnections: maxIdleConns,
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
