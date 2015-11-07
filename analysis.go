package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/zaynetro/tweets-sentiment-analysis/Godeps/_workspace/src/github.com/ChimeraCoder/anaconda"
)

type AnalysisResult struct {
	AnalysedTweet AnalysedTweet
	Failed        bool
}

type AnalysedTweet struct {
	Text      string    `json:"text"`
	Lang      string    `json:"lang"`
	User      TweetUser `json:"user"`
	Sentiment Sentiment `json:"sentiment"`
}

type TweetUser struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	ScreenName     string `json:"screenName"`
	FollowersCount int    `json:"followersCount"`
	FriendsCount   int    `json:"friendsCount"`
}

type AlchemySentimentScore struct {
	Type  string `json:"type"`
	Score string `json:"score"`
}

type Sentiment struct {
	Type  string  `json:"type"`
	Score float64 `json:"score"`
}

type AlchemySentiment struct {
	Status    string                `json:"status"`
	Sentiment AlchemySentimentScore `json:"docSentiment"`
}

func analyseTweets(tweets []anaconda.Tweet) {
	var processed []AnalysedTweet
	for _, tweet := range tweets {
		//log.Printf("Analyse lang: %s and tweet: %s\n", tweet.Lang, tweet.Text)
		analysed := analyseTweet(tweet)
		if !analysed.Failed {
			processed = append(processed, analysed.AnalysedTweet)
		}
	}

	log.Println("Processed")
	for _, tweet := range processed {
		log.Printf("%v\n\n", tweet)
		sendToDb(tweet)
	}
	//log.Printf("Processed: %v\n", processed)
}

func analyseTweet(tweet anaconda.Tweet) AnalysisResult {
	data := url.Values{}
	data.Set("apikey", config.Alchemy.ApiKey)
	data.Add("text", tweet.Text)
	data.Add("outputMode", "json")

	req, err := http.NewRequest("POST", config.Alchemy.Url, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Printf("Cannot form sentiment request: %s\n", err)
	}
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to process request: %s\n", err)
		return AnalysisResult{
			Failed: true,
		}
	}
	defer res.Body.Close()

	//log.Printf("Res status: %s\n", res.Status)
	body, _ := ioutil.ReadAll(res.Body)
	//log.Printf("Res: %s\n", body)

	alchemySentiment := AlchemySentiment{}
	err = json.Unmarshal(body, &alchemySentiment)
	if err != nil {
		log.Printf("Failed to marshal alchemy response: %s\n%s\n", err, body)
	}

	if alchemySentiment.Status != "OK" {
		return AnalysisResult{
			Failed: true,
		}
	}

	sentiment := Sentiment{
		Type: alchemySentiment.Sentiment.Type,
	}

	score, err := strconv.ParseFloat(alchemySentiment.Sentiment.Score, 32)
	if err != nil {
		score = 0.0
	}
	sentiment.Score = score

	return AnalysisResult{
		AnalysedTweet: AnalysedTweet{
			Text: tweet.Text,
			Lang: tweet.Lang,
			User: TweetUser{
				Id:             tweet.User.Id,
				Name:           tweet.User.Name,
				ScreenName:     tweet.User.ScreenName,
				FollowersCount: tweet.User.FollowersCount,
				FriendsCount:   tweet.User.FriendsCount,
			},
			Sentiment: sentiment,
		},
		Failed: false,
	}

}
