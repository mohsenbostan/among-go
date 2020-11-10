package main

import (
	"github.com/joho/godotenv"
	"github.com/mohsenbostan/among-go/actions"
	"log"
	"os"
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
	// Loading environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	// Creating Discord and Twitter objects
	var discord actions.Discord
	var twitter actions.Twitter

	// Setup a ticker to send request each minute
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {

		// Call the handler
		Handle(discord, twitter)
	}
}
