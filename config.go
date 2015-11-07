package main

import (
	"log"
	"os"
)

type Config struct {
	Db      DbConf
	Twitter TwitterConf
	Alchemy AlchemyConf
}

type DbConf struct {
	Url string
}

type TwitterConf struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

type AlchemyConf struct {
	ApiKey string
	Url    string
}

func parseConfig() {
	config.Db = DbConf{
		Url: os.Getenv("DB_URL"),
	}

	config.Alchemy = AlchemyConf{
		ApiKey: os.Getenv("ALCHEMY_API_KEY"),
		Url:    "http://gateway-a.watsonplatform.net/calls/text/TextGetTextSentiment",
	}

	config.Twitter = TwitterConf{
		ConsumerKey:       os.Getenv("TWITTER_CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("TWITTER_CONSUMER_SECRET"),
		AccessToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("TWITTER_ACCESS_TOKEN_SECRET"),
	}

	if len(config.Db.Url) == 0 {
		log.Fatalln("Specify DB_URL")
	}

	if len(config.Alchemy.ApiKey) == 0 {
		log.Fatalln("Specify ALCHEMY_API_KEY")
	}

	if len(config.Twitter.ConsumerKey) == 0 {
		log.Fatalln("Specify TWITTER_CONSUMER_KEY")
	}

	if len(config.Twitter.ConsumerSecret) == 0 {
		log.Fatalln("Specify TWITTER_CONSUMER_SECRET")
	}

	if len(config.Twitter.AccessToken) == 0 {
		log.Fatalln("Specify TWITTER_ACCESS_TOKEN")
	}

	if len(config.Twitter.AccessTokenSecret) == 0 {
		log.Fatalln("Specify TWITTER_ACCESS_TOKEN_SECRET")
	}
}
