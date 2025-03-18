package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	"github.com/ivan-ca97/rush/backend/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Environment variables file error")
	}

	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		fmt.Print("Port invalid")
	}

	logDirectory := os.Getenv("LOGS_DIRECTORY")
	jwtKey := os.Getenv("JWT_KEY")

	api.RushServer(port, logDirectory, jwtKey, time.Hour*2400)
}
