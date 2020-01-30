package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/nlopes/slack"
)

type XKCD_COMIC struct {
	Month      string `json:"month,omitempty"`
	Num        int    `json:"num,omitempty"`
	Link       string `json:"link,omitempty"`
	Year       string `json:"year,omitempty"`
	News       string `json:"news,omitempty"`
	SafeTitle  string `json:"safe_title,omitempty"`
	Transcript string `json:"transcript,omitempty"`
	Alt        string `json:"alt,omitempty"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day,omitempty"`
}

// getRandomComic returns the url of a random xkcd comic. https://xkcd.com/<id>/
func main() {
	// Create client; do not redirect
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	// GET random xkcd comic
	// TODO: try client.Head()
	resp, err := client.Get("https://c.xkcd.com/random/comic/")
	if err != nil {
		fmt.Println("Failure : ", err)
	}

	rLoc := resp.Header.Get("Location")
	sendRandomComic(string(rLoc))
}

// sendRandomComic will process the fetched comic and post to slack
func sendRandomComic(rLoc string) {
	client := &http.Client{}

	// Create request
	jsonEndP := rLoc + "info.0.json"

	resp, err := client.Get(jsonEndP)
	if err != nil {
		fmt.Println("Failure : ", err)
	}
	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	comic := XKCD_COMIC{}
	json.Unmarshal(respBody, &comic)

	// Display Results
	//fmt.Println("response Status : ", resp.Status)
	//fmt.Println("response Headers : ", resp.Header)
	//fmt.Println("response Body : ", string(respBody))
	//fmt.Printf("Title: %s, ImgURL: %s, AltText: %s", comic.Title, comic.Img, comic.Alt)

	// we should use Slack Webhook to post directly, and have a comment
	// print to #testing-zone
	/*
	hpath, exists := os.LookupEnv("SLACK_HOOK_PATH_MEMES")
	if !exists {
		fmt.Println("Failure : env_var not found")
	}
	*/
	hpath := "services/TQZHGFGMS/BT6J4F7FA/sXyVRP0mn03xg96mO874F5ZQ"
	url := "https://hooks.slack.com/" + hpath
	// setup post to be title + url
	post := comic.Title + " " + comic.Img
	payload := &slack.WebhookMessage{
		Text:    string(post),
		Channel: "#testing-zone",
	}

	slack.PostWebhook(url, payload)
}
