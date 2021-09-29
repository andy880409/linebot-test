package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var (
	bot *linebot.Client
	err error
)

func init() {
	//先在Heroku上設置CHANNEL_SECRET 以及 CHANNEL_TOKEN
	bot, err = linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Println(err)
	}
}
func main() {
	http.HandleFunc("/callback", handler)
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal(err)
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	if err != nil {
		log.Fatal(err)
	}
	for _, event := range events {
		//訊息事件
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("reply token:"+event.ReplyToken+"ID:"+message.ID+":"+message.Text)).Do()
			}
		}
		fmt.Println(123)
		fmt.Println("port:", os.Getenv("PORT"))
		fmt.Println("event:", event)
	}
}
