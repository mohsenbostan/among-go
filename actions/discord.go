package actions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type Discord struct {
	Members []DiscordMember `json:"members"`
}

type DiscordMember struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Status   string `json:"status"`
}

// Check member's status
func (d *Discord) IsOnline() (bool, error) {
	isOnline := false

	// Discord endpoint url
	reqUrl := "https://discord.com/api/guilds/" + os.Getenv("DISCORD_SERVER_ID") + "/widget.json"

	// Create new request
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return false, err
	}

	// Send request using http Client
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	// Read the response body
	body, _ := ioutil.ReadAll(res.Body)

	// Decode json response
	err = json.Unmarshal(body, &d)
	if err != nil {
		return false, err
	}

	// Check if the member is online
	for _, member := range d.Members {
		if member.Username == os.Getenv("DISCORD_USERNAME") {
			isOnline = member.Status == "online"
		}
	}

	return isOnline, nil
}
