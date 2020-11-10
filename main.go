package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mohsenbostan/among-go/actions"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
	Handler
	This function will check the Discord's member's status and if the member was online,
	it will update his Twitter profile by adding MESSAGE at the end of the profile description.
*/
func Handle(d actions.Discord, t actions.Twitter) {
	log.Println("Start checking discord status...")

	// Checking if member is online
	isOnline, err := d.IsOnline()
	if err != nil {
		log.Fatalln(err)
	}

	// Getting member's twitter profile
	profile, err := t.GetProfile()
	if err != nil {
		log.Fatalln(err)
	}

	// Preparing new description by adding message to end of it
	message := os.Getenv("MESSAGE")
	newDescription := profile.Description + message

	if isOnline {
		// Checking if the profile was not updated before and then update the Twitter description
		if profile.Description[(len(profile.Description)-len(message)):] != message {

			// Updating Twitter description
			ok, err := t.UpdateProfile("description", newDescription)
			if err != nil {
				log.Fatalln(err)
			} else if ok {
				log.Println("Twitter profile updated to online!")
			}
		} else {
			log.Println("Twitter profile is up-to-date!")
		}
	} else {
		// Checking if the profile was updated before and then remove the message from the end of the description
		if profile.Description[(len(profile.Description)-len(message)):] == message {

			// remove the message from the end of the description
			newDescription = profile.Description[:(len(profile.Description) - len(message))]
			ok, err := t.UpdateProfile("description", newDescription)
			if err != nil {
				log.Fatalln(err)
			} else if ok {
				log.Println("Twitter profile updated to offline!")
			}
		}
		log.Println("Discord member was offline!")
	}
}

func main() {
	fmt.Print("\n    ___                                            ______     \n   /   |   ____ ___   ____   ____   ____ _        / ____/____ \n  / /| |  / __ `__ \\ / __ \\ / __ \\ / __ `/______ / / __ / __ \\\n / ___ | / / / / / // /_/ // / / // /_/ //_____// /_/ // /_/ /\n/_/  |_|/_/ /_/ /_/ \\____//_/ /_/ \\__, /        \\____/ \\____/ \n                                 /____/                       \n")

	fmt.Println(strings.Repeat("=", 62))

	// Loading environment variables
	LoadEnvVariables()

	// Creating Discord and Twitter objects
	var discord actions.Discord
	var twitter actions.Twitter

	// Setup a ticker to send request each minute
	interval, err := strconv.Atoi(os.Getenv("INTERVAL"))
	if err != nil {
		log.Fatalln()
	}

	ticker := time.NewTicker(time.Duration(interval) * time.Minute)
	for range ticker.C {

		// Call the handler
		Handle(discord, twitter)
	}
}

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
