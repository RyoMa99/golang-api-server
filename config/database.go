package config

import (
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	db  *gorm.DB
	err error
)

// DB接続
func Connect() *gorm.DB {
	// 実行環境取得
	env := os.Getenv("ENV")

	if "production" == env {
		env = "production"
	} else {
		env = "development"
	}

	// 環境変数取得
	godotenv.Load(".env." + env)
	godotenv.Load()

	// DB接続
	connect := mysql.Config{
		DBName:    os.Getenv("DB_NAME"),
		User:      os.Getenv("DB_USER"),
		Passwd:    os.Getenv("DB_PASSWORD"),
		Addr:      os.Getenv("DB_ADDR"),
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       time.UTC,
	}

	db, err = gorm.Open("mysql", connect.FormatDSN())

	if err != nil {
		panic(err)
	}

	return db
}

// DB終了
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
