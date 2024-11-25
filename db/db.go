package db

import (
	"database/sql"
	"fmt"
	"github.com/GDG-on-Campus-KHU/SC7_BE/config"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init(cfg *config.Config) {
	// 데이터베이스 연결 문자열 생성
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	// 재시도 로직 추가
	var err error
	for i := 0; i < 10; i++ { // 최대 10번 재시도
		DB, err = sql.Open("mysql", dsn)
		if err == nil && DB.Ping() == nil {
			log.Println("Database connection established successfully!")
			return
		}
		log.Printf("Failed to connect to database (attempt %d): %v", i+1, err)
		time.Sleep(2 * time.Second) // 2초 대기 후 재시도
	}

	log.Fatalf("Database connection failed after retries: %v", err)
}

func Close() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Printf("Error closing the database connection: %v", err)
		} else {
			log.Println("Database connection closed successfully.")
		}
	}
}
