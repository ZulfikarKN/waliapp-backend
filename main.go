package main

import (
	"Go/WALIAPP-BACKEND/config"
	"Go/WALIAPP-BACKEND/router"
	"log"
)

func init() {
	config.ReadEnv()
}

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	app := router.SetupRouter()
	app.Run()
}