package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	//"os"
	"time"

	"github.com/nlopes/slack"
	//"github.com/turnage/redditproto"
)

type RedditPosts struct {
	Kind string `json:"kind,omitempty"`
	Data struct {
		Modhash  string `json:"id,omitempty"`
		Dist     int    `json:"dist,omitempty"`
		Children []struct {
			Kind string     `json:"kind,omitempty"`
			Data RedditPost `json:"data,omitempty"`
		} `json:"children,omitempty"` //only 25 items return
	} `json:"data,omitempty"`
}

type RedditPost struct {
	Title string `json:"title,omitempty"`
	URL   string `json:"url,omitempty"`
}

// main will fetch a random Top Post (image), then post to slack
func main() {
	client := &http.Client{}

	url := "https://www.reddit.com/r/ProgrammerHumor.json"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-agent", "test test")
	resp, _ := client.Do(req)

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	raw := RedditPosts{}
	json.Unmarshal(respBody, &raw)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	//fmt.Println("response Headers : ", resp.Header)
	//fmt.Println("response Body : ", string(respBody))

	// random number from 25 response items
	postNum := randomPost()
	fmt.Println("selected Post: ", postNum)

	processPost := raw.Data.Children[postNum].Data
	postTitle := processPost.Title
	postURL := processPost.URL // TODO: check url formats (png, img, jpg, gif)
	fmt.Println("reddit Title: ", string(postTitle))
	fmt.Println("reddit URL: ", string(postURL))

	// print to #testing-zone
	hpath, exists := os.LookupEnv("SLACK_HOOK_PATH_MEMES")
	if !exists {
		fmt.Println("Failure : env_var not found")
	}

	slackurl := "https://hooks.slack.com/" + hpath
	// setup post to be title + url
	post := postTitle + " " + postURL
	payload := &slack.WebhookMessage{
		Text:    string(post),
		Channel: "#testing-zone",
	}

	slack.PostWebhook(slackurl, payload)

	client.CloseIdleConnections()
}

// randomPost returns a random int between 0-24
func randomPost() int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 24
	postNum := rand.Intn(max-min+1) + min
	return postNum
}
