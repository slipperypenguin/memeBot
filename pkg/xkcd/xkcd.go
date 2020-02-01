package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//"strings"

	"github.com/joho/godotenv"
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
	json.Unmarshal(respBody, &comic)

	// Display Results
	//fmt.Println("response Status : ", resp.Status)
	//fmt.Println("response Headers : ", resp.Header)
	//fmt.Println("response Body : ", string(respBody))
	//fmt.Printf("Title: %s, ImgURL: %s, AltText: %s", comic.Title, comic.Img, comic.Alt)

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

	
	//url := "https://hooks.slack.com/" + hpath
	// setup post to be title + url
	//post := comic.Title + " " + comic.Img

	// Use Block API
	//post := SlackBlock(comic)
	// /// //// ///
	// Header Section
	headerText := slack.NewTextBlockObject("mrkdwn", comic.Title, false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// Actual post
	/*
		comicImg := slack.NewImageBlock(comic.Img, "some profile", "image-block", headerText)
		//comicImg := slack.NewImageBlockElement(comic.Img, "test text")
		imgSection := slack.NewSectionBlock(comicImg, nil, nil)
	*/

	comicImg := slack.NewTextBlockObject("mrkdwn", comic.Img, false, false)
	imgSection := slack.NewSectionBlock(comicImg, nil, nil)

	// Divider
	divSection := slack.NewDividerBlock()

	// Alt Text
	aText := slack.NewTextBlockObject("mrkdwn", comic.Alt, false, false)
	altTextSection := slack.NewSectionBlock(aText, nil, nil)

	// Build Message with blocks created above
	msg := slack.NewBlockMessage(
		headerSection,
		imgSection,
		divSection,
		altTextSection,
	)

	b, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

	//post := string(b)

	/*
	p := []byte(post)
	var s struct{}
	json.Unmarshal(p, &s)
	fmt.Println("unmarshalled Body : ", string(p))

	fix := []string{post}
	i := 0
	for x := 0; x < 2; x++ {
		// remove the element at index i from fix
		fix[i] = fix[len(fix)-1] // copy last element to index i
		fix[len(fix)-1] = ""     // erase last element
		fix = fix[:len(fix)-1]   // truncate slice
	}

	jString := strings.Join(fix, " ")
	fmt.Println("unmarshalled 'fix' : ", string(jString))
	*/
	/// /// /// /// /// ///
/*
	payload := &slack.WebhookMessage{
		Text:    string(jString),
		Channel: "#testing-zone",
	}

	slack.PostWebhook(url, payload)
	*/
}

/*
// SlackBlock formats an interactive message
func SlackBlock(c XKCD_COMIC) string {
	// Header Section
	headerText := slack.NewTextBlockObject("mrkdwn", c.Title, false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)

	// Actual post
	comicImg := slack.NewImageBlockElement(c.Img, "stub text")
	imgSection := slack.NewSectionBlock(headerText, nil, nil)

	// Divider
	divSection := slack.NewDividerBlock()

	// Alt Text
	aText := slack.NewTextBlockObject("mrkdwn", c.Alt, false, false)
	altTextSection := slack.NewSectionBlock(aText, nil, nil)

	// Build Message with blocks created above
	msg := slack.NewBlockMessage(
		headerSection,
		imgSection,
		divSection,
		altTextSection,
	)

	b, err := json.MarshalIndent(msg, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))

	return string(b)
}
*/
