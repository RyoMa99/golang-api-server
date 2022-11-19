package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	connect := fmt.Sprintf(
		"%s:%s@%s/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_Protocol"),
		os.Getenv("DB_NAME"),
	)

	db, err = gorm.Open("mysql", connect)

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
