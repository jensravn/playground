package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envvar struct {
	projectID string
}

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	projectID := os.Getenv("PROJECT_ID")
	fmt.Printf("PROJECT_ID=%s \n", projectID)
}
