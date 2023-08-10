package main

import (
	"comment-backend/model"
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func getServiceProperties() model.ServiceProperties {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	return getEnv()
}

func getEnv() model.ServiceProperties {
	fmt.Println("Starting Load Config " + time.Now().Format(time.RFC3339Nano))

	// SERVICE
	svcPort, _ := strconv.Atoi(os.Getenv("SERVICE_PORT"))
	svcPoolSizeConnection, _ := strconv.Atoi(os.Getenv("SERVICE_POOL_SIZE_CONNECTION"))
	svcTimezone, _ := time.LoadLocation(os.Getenv("SERVICE_TIMEZONE"))

	svcProperties := model.ServiceProperties{
		ServiceName:               os.Getenv("SERVICE_NAME"),
		ServicePort:               svcPort,
		ServicePoolSizeConnection: svcPoolSizeConnection,
		ServiceTimezone:           svcTimezone,
		Database: model.DBConfig{
			IP:       os.Getenv("DB_IP"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		CustomUserID: os.Getenv("CUSTOM_USER_ID"),
	}
	if err := validator.New().Struct(svcProperties); err != nil {
		panic(err)
	}
	fmt.Println("Finish Load Config " + time.Now().Format(time.RFC3339Nano))
	return svcProperties
}

func databaseConnect(svcName string, dbConfig model.DBConfig) *bun.DB {
	addr := fmt.Sprintf("%s:%s", dbConfig.IP, dbConfig.Port)
	fmt.Printf("Connecting to Database : '%v'\n", addr)
	openDB := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(addr),
		pgdriver.WithUser(dbConfig.User),
		pgdriver.WithPassword(dbConfig.Password),
		pgdriver.WithDatabase(dbConfig.Name),
		pgdriver.WithApplicationName(svcName),
		pgdriver.WithTLSConfig(nil),
	))
	db := bun.NewDB(openDB, pgdialect.New())
	if err := db.PingContext(context.Background()); err != nil {
		panic(err)
	}
	fmt.Printf("Connected to Database : '%v'\n", addr)
	return db
}
