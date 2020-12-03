package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	defaultEnvs := []string{
		"TWITTER_USERNAME",
		"TWITTER_CONSUMER_KEY",
		"TWITTER_CONSUMER_SECRET",
		"TWITTER_ACCESS_TOKEN",
		"TWITTER_TOKEN_SECRET",
		"DISCORD_SERVER_ID",
		"DISCORD_USERNAME",
		"INTERVAL",
		"MESSAGE",
	}

	for _, env := range defaultEnvs {
		val, found := os.LookupEnv(env)
		if len(val) <= 0 || !found {
			log.Fatalln("All environment variables should be defined and they must have valid values. you can copy defaults from: '.env.example' .")
		}
		if env == "INTERVAL" {
			interval, err := strconv.Atoi(val)
			if err != nil || interval <= 0 {
				log.Fatalln("INTERVAL must be an unsigned number and the minimum is 1.")
			}
		}
	}
}
