package main

import (
	"log"
	"net/url"

	"github.com/zaynetro/tweets-sentiment-analysis/Godeps/_workspace/src/github.com/ChimeraCoder/anaconda"
)

var config = Config{}

func main() {
	parseConfig()

	tweets := getTweets()
	analyseTweets(tweets)
}

func getTweets() []anaconda.Tweet {
	anaconda.SetConsumerKey(config.Twitter.ConsumerKey)
	anaconda.SetConsumerSecret(config.Twitter.ConsumerSecret)
	api := anaconda.NewTwitterApi(config.Twitter.AccessToken, config.Twitter.AccessTokenSecret)

	v := url.Values{}
	v.Set("count", "200")
	tweets, err := api.GetHomeTimeline(v)
	if err != nil {
		log.Printf("Can't get home timeline: %s\n", err)
		return []anaconda.Tweet{}
	}

	return tweets
}
