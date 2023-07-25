package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init logger (bonus)
	//init db connection

	//create user store
	//create user service
	//create user handler

	//route init

	pass := os.Getenv("MYSQL_PASSWORD")
	fmt.Println(pass)
}
