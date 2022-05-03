package main

import (
	"discord2line/bot"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	loadEnv()
	goc, err := bot.NewBot(os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("D2L start.")
	err = goc.Start()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func loadEnv() {
	current, err := os.Getwd()
	path := current + "/.env"
	if err != nil {
		log.Println(path)
		log.Fatalln(err.Error())
	}
	godotenv.Load(current + "/.env")
}
