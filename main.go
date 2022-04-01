package main

import (
	"github.com/nexsabre/mikropoker-go/db"
	"github.com/nexsabre/mikropoker-go/server"
)

func main() {
	db.Init()
	conn := db.GetDB()
	server.Start(conn)
}
