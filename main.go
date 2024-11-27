package main

import (
	"github.com/GDG-on-Campus-KHU/SC7_BE/config"
	"github.com/GDG-on-Campus-KHU/SC7_BE/db"
	"github.com/GDG-on-Campus-KHU/SC7_BE/routes"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 환경 변수 로드 및 설정 초기화
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// DB 연결 초기화
	db.Init(cfg)
	defer db.Close()

	// 라우팅 설정
	router := routes.InitRoutes()

	// 서버 종료 신호 처리
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		log.Println("Shutting down server...")
		os.Exit(0)
	}()

	// 서버 실행
	log.Printf("Server is running on port: %d", cfg.Port)
	if err := router.Run("0.0.0.0:8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
