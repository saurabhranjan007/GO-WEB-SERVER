package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// 📝 1. Use capitalised function names, if you wanna use the same in other files
// 📝 2. LOAD ENV: if PORT is defined in env, server will auto run on that port

func LoadEnvvariables() {

	log.Println("LOADING ENV..")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
