# Tweets sentiment analysis

Get user twitter home timeline, process them through
sentiment analysis using Alchemy API on Bluemix and
send to the Cloudfound DB.

## What?

Junction 2015 hackathon

## Deploy

1. Register to IBM Bluemix
2. Clone repo: `git clone git@github.com:zaynetro/tweets-sentiment-analysis.git`
3. Compile: `make`
4. Update `manifest.yml` file to meet your needs
5. Connect to [bluemix cli](https://github.com/cloudfoundry/cli): `cf login`
6. Deploy to Bluemix: `cf deploy`

**NOTE:** You also need to set up additional services:

* Cloudfound DB
* Alchemy API
* Twitter API application

and set environment variables as defined in `config.go`

**LICENSE:** MIT
