package main

import (
	"github.com/amankumarsingh77/cmr/config"
	"github.com/amankumarsingh77/cmr/internal/server"
	"github.com/amankumarsingh77/cmr/pkg/db/aws"
	"github.com/amankumarsingh77/cmr/pkg/db/postgres"
	"github.com/amankumarsingh77/cmr/pkg/db/redis"
	"github.com/amankumarsingh77/cmr/pkg/logger"
	"log"
)

func main() {
	log.Println("Starting api server")
	configPath := "config.yml"
	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)
	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	} else {
		appLogger.Infof("Postgres connected, Status: %#v", psqlDB.Stats())
	}
	defer psqlDB.Close()

	redisClient := redis.NewRedisClient(cfg)
	defer redisClient.Close()
	appLogger.Info("Redis connected")
	s3Client, err := aws.NewAwsClient(cfg.S3.Endpoint, cfg.S3.Region, cfg.S3.AccessKey, cfg.S3.SecretKey)
	if err != nil {
		appLogger.Fatalf("Failed to create S3 client: %v", err)
	}
	appLogger.Info("S3 client connected")
	s := server.NewServer(cfg, psqlDB, redisClient, s3Client, appLogger)
	if err := s.Run(); err != nil {
		appLogger.Fatal(err)
	}
}
