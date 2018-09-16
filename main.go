package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"time"

	"gopkg.in/telegram-bot-api.v4"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tidwall/gjson"
)

func run() {
	/*
		Authorize Telegram Bot
	*/
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	/*
		Get Reddit data
	*/
	redditJSONString, err := getRedditJSONString(RedditEndPoint)

	if err != nil {
		fmt.Printf("%s", err)
	}

	result := gjson.Get(redditJSONString, "data.children")

	/*
		Each Post
	*/
	result.ForEach(func(key, value gjson.Result) bool {
		childResultFlair := gjson.Get(value.String(), "data.link_flair_text")
		childResultTitle := gjson.Get(value.String(), "data.title")
		childResultURL := gjson.Get(value.String(), "data.url")
		childResultTimestamp := gjson.Get(value.String(), "data.created_utc")

		println(childResultFlair.String())
		println(childResultTitle.String())
		println(childResultURL.String())
		println(time.Now().Unix())
		println(childResultTimestamp.Int())

		// Seconds
		Interval, parseIntErr := strconv.ParseInt(os.Getenv("INTERVAL"), 10, 64)
		if parseIntErr != nil {
			Interval = 60
		}

		GroupID, parseIntErr := strconv.ParseInt(os.Getenv("GROUP_ID"), 10, 64)
		if parseIntErr != nil {
			GroupID = -212534862
		}

		if time.Now().Unix()-childResultTimestamp.Int() < Interval && childResultFlair.String() == "Media" {
			println("Inside interval && Is Media - send msg: " + childResultTitle.String())

			title := childResultTitle.String()
			link := childResultURL.String()

			msg := tgbotapi.NewMessage(GroupID, makeTextMessage(title, link))
			msg.ParseMode = "Markdown"

			bot.Send(msg)
		}

		println("\n")

		return true // keep iterating
	})

	log.Printf("Job done")
}

func main() {
	// run()
	lambda.Start(run)
}
