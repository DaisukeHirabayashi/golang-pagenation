package main

import (
	"log"

	"github.com/DaisukeHirabayashi/golang-pagenation/db"
	"github.com/DaisukeHirabayashi/golang-pagenation/server"
)

func main() {
	log.Print("Build starts...")
	db.Init()
	server.Init()
	db.Close()
}
