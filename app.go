package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

const (
	UPDATE_MIN   = 1
	DEFAULT_PORT = "8080"
)

var config = Config{}

func main() {
	parseConfig()

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	http.HandleFunc("/", helloworld)

	log.Printf("Starting app on port %+v\n", port)
	http.ListenAndServe(":"+port, nil)

	ticker := time.NewTicker(UPDATE_MIN * time.Minute)

	lastId := int64(663074549127356416)
	for _ = range ticker.C {
		log.Printf("Start processing tweets since: %d\n", lastId)
		lastId = processTweets(lastId)
	}

}

func processTweets(sinceId int64) int64 {
	var lastId int64
	tweets := getTweets(sinceId)
	if l := len(tweets); l > 0 {
		lastId = tweets[0].Id
	}
	analysed := analyseTweets(tweets)
	for _, tweet := range analysed {
		//log.Printf("Analysed: %v\n", tweet)
		sendToDb(tweet)
	}

	if l := len(analysed); l > 0 {
		log.Printf("Last processed tweet: %s\n", analysed[0].Text)
	}

	log.Printf("Processed: %d tweets\n", len(analysed))

	return lastId
}

func getTweets(sinceId int64) []anaconda.Tweet {
	anaconda.SetConsumerKey(config.Twitter.ConsumerKey)
	anaconda.SetConsumerSecret(config.Twitter.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.Twitter.AccessToken, config.Twitter.AccessTokenSecret)

	v := url.Values{}
	v.Set("count", "2")
	if sinceId != 0 {
		v.Set("since_id", strconv.FormatInt(sinceId, 10))
	}
	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
		log.Printf("Can't get home timeline: %s\n", err)
		return []anaconda.Tweet{}
	}

	log.Printf("Fetched %d tweets\n", len(tweets))

	return tweets
}

func helloworld(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}
