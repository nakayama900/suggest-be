package main

import (
	"chatgpt-230308/src/web"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	godotenv.Load(".env.local")
	godotenv.Load()
	if os.Getenv("GPT_KEY") == "" {
		panic("GPT_KEY is not set")
	}
	router := web.Router(web.Server{})
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}
