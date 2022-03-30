package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nexsabre/mikropoker-go/config"
	"github.com/nexsabre/mikropoker-go/db"
	"github.com/nexsabre/mikropoker-go/server"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()
	config.Init(*environment)
	db.Init()
	conn := db.GetDB()
	server.Start(conn)
}
