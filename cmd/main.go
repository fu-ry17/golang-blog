package main

import (
	"log"
	"simple-blog/cmd/api"
	"simple-blog/config"
	db2 "simple-blog/db"
)

func main() {
	db, err := db2.NewStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(config.Env.PORT, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
