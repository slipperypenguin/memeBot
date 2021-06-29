package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
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

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found... using env_var values")
	}
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
	err = json.Unmarshal(respBody, &comic)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	//fmt.Println("response Headers : ", resp.Header)
	//fmt.Println("response Body : ", string(respBody))
	//fmt.Printf("Title: %s, ImgURL: %s, AltText: %s", comic.Title, comic.Img, comic.Alt)
	fmt.Println("xkcd Title: ", comic.Title)
	fmt.Println("image URL: ", comic.Img)
	fmt.Println("xkcd ID: ", comic.Num)

	// print to #testing-zone
	var hpath string
	testPath, exists := os.LookupEnv("TESTPATH")
	if exists {
		fmt.Println("Test .env found 🚧")
		hpath = testPath
	}

	// print to #spicy_memes
	prodPath, exists := os.LookupEnv("SLACK_HOOK_PATH_MEMES")
	if exists {
		fmt.Println("Prod env_var found 🛳")
		hpath = prodPath
	}

	// Slack Block API
	// Heading
	headerText := slack.NewTextBlockObject("plain_text", string(comic.Alt), false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)
	// Divider
	divSection := slack.NewDividerBlock()
	// Img Heading + Image
	imgHeaderText := slack.NewTextBlockObject("plain_text", string(comic.Title), false, false)
	imgSection := slack.NewImageBlock(string(comic.Img), string(comic.Alt), "", imgHeaderText)

	blocks := make([]slack.Block, 0)

	blocks = append(blocks, headerSection)
	blocks = append(blocks, divSection)
	blocks = append(blocks, imgSection)

	slackurl := "https://hooks.slack.com/" + hpath
	payload := slack.WebhookMessage{
		Channel: "#testing-zone",
		Blocks:  &slack.Blocks{BlockSet: blocks},
	}

	err = slack.PostWebhook(slackurl, &payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	client.CloseIdleConnections()
}
