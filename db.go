package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func sendToDb(tweet AnalysedTweet) {
	dbUrl := config.Db.Url + "tweets/"

	json, err := json.Marshal(tweet)
	if err != nil {
		log.Printf("Cannot marshal tweet: %s\n", err)
		return
	}

	req, err := http.NewRequest("POST", dbUrl, bytes.NewBuffer(json))
	if err != nil {
		log.Printf("Cannot form request: %s\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Cannot process response: %s\n", err)
		return
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	log.Printf("DB: %s\n", body)
}
