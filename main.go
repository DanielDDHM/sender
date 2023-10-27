package main

import (
	svr "DanielDDHM/sender/src/server"
	utils "DanielDDHM/sender/src/utils"
	"log"
)

func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Println("[SENDER] Cannot read .env file")
		log.Fatal(err)
	}

	log.Println(config)

	server := svr.CreateNewServer()
	server.Run()

	log.Fatal("[SENDER] Application started successfully, enjoy!")
}
