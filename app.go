package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/zaynetro/tweets-sentiment-analysis/Godeps/_workspace/src/github.com/ChimeraCoder/anaconda"
)

const (
	DEFAULT_PORT = "8080"
)

var config = Config{}

var index = template.Must(template.ParseFiles(
	"templates/_base.html",
	"templates/index.html",
))

func helloworld(w http.ResponseWriter, req *http.Request) {
	index.Execute(w, nil)
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	tweets := getTweets()
	analyseTweets(tweets)
	return

	//http.HandleFunc("/", helloworld)
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//log.Printf("Starting app on port %+v\n", port)
	//http.ListenAndServe(":"+port, nil)
}

func getTweets() []anaconda.Tweet {
	anaconda.SetConsumerKey(config.Twitter.ConsumerKey)
	anaconda.SetConsumerSecret(config.Twitter.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.Twitter.AccessToken, config.Twitter.AccessTokenSecret)

	v := url.Values{}
	v.Set("count", "5")
	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
		log.Printf("Can't get home timeline: %s\n", err)
		return []anaconda.Tweet{}
	}

	return tweets
}
