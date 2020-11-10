package actions

import (
	"encoding/json"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Twitter int

// Create default http client for Twitter
func (t *Twitter) HttpClient() *http.Client {
	config := oauth1.NewConfig(os.Getenv("TWITTER_CONSUMER_KEY"), os.Getenv("TWITTER_CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_TOKEN_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	return httpClient
}

// Get user's Twitter profile data
func (t *Twitter) GetProfile() (*twitter.User, error) {
	client := twitter.NewClient(t.HttpClient())

	user, _, err := client.Users.Show(&twitter.UserShowParams{
		ScreenName: os.Getenv("TWITTER_USERNAME"),
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update user's Twitter profile data
func (t *Twitter) UpdateProfile(field string, val string) (bool, error) {
	// Define params that should update in user's Twitter profile
	params := url.Values{}
	params.Add(field, val)

	// Twitter's profile update endpoint url
	reqUrl := "https://api.twitter.com/1.1/account/update_profile.json?" + params.Encode()

	// Create new request
	req, err := http.NewRequest("POST", reqUrl, nil)
	if err != nil {
		return false, err
	}

	// Send req using http Client
	client := t.HttpClient()
	res, err := client.Do(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	// Read the response body
	body, _ := ioutil.ReadAll(res.Body)

	// Decode json response
	user := twitter.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return false, err
	}

	return user.Description == val, nil
}
