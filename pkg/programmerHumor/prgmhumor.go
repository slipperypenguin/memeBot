package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

type RedditPosts struct {
	Kind string `json:"kind,omitempty"`
	Data struct {
		Modhash  string `json:"id,omitempty"`
		Dist     int    `json:"dist,omitempty"`
		Children []struct {
			Kind string     `json:"kind,omitempty"`
			Data RedditPost `json:"data,omitempty"`
		} `json:"children,omitempty"` //only 25 items return by default
	} `json:"data,omitempty"`
}

type RedditPost struct {
	Title string `json:"title,omitempty"`
	URL   string `json:"url,omitempty"`
}

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found... using env_var values")
	}
}

// main will fetch a random Top Post (image/gif), then post to slack
func main() {
	client := &http.Client{
		Timeout: time.Second * 60,
	}

	url := "https://www.reddit.com/r/ProgrammerHumor.json"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-agent", "test test")
	resp, _ := client.Do(req)

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)
	raw := RedditPosts{}
	json.Unmarshal(respBody, &raw)
	fmt.Println("response Status: ", resp.Status)

	// random number from 25 response items
	postNum := randomPost()
	fmt.Println("selected Post: ", postNum)
	processPost := raw.Data.Children[postNum].Data
	postTitle := processPost.Title
	postURL := processPost.URL

Loop:
	for {
		str := strings.Split(postURL, ".")
		fileExt := str[len(str)-1]
		switch fileExt {
		case "gif":
			break Loop
		case "img":
			break Loop
		case "jpg":
			break Loop
		case "png":
			break Loop
		default:
			postNum = randomPost()
			fmt.Println("selected New Post: ", postNum)
			processPost = raw.Data.Children[postNum].Data
			postTitle = processPost.Title
			postURL = processPost.URL
		}
	}

	fmt.Println("reddit Title: ", string(postTitle))
	fmt.Println("reddit URL: ", string(postURL))

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
	headerText := slack.NewTextBlockObject("plain_text", string(postTitle), false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)
	// Divider + Image
	divSection := slack.NewDividerBlock()
	imgSection := slack.NewImageBlock(string(postURL), "test", "", headerText)

	blocks := make([]slack.Block, 0)

	blocks = append(blocks, headerSection)
	blocks = append(blocks, divSection)
	blocks = append(blocks, imgSection)

	slackurl := "https://hooks.slack.com/" + hpath
	payload := &slack.WebhookMessage{
		Blocks:  slack.Blocks{blocks},
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
