package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	log.Print("DB init...")
	// TODO : os.Getenv("DATABASE_URL") does not work...
	db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=root password=password dbname=go-pagenation sslmode=disable")
	if err != nil {
		panic(err)
	}
	log.Print("DB connection success!")
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
