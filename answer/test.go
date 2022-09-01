package main

import (
	"fmt"
	"log"
	"net/http"
	// "github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	// ハンドラの登録
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/callback", lineHandler)

	fmt.Println("http://localhost:8080 で起動中...")
	// HTTPサーバーを起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World!!"
	fmt.Fprintf(w, msg)
}

func lineHandler(w http.ResponseWriter, r *http.Request) {
	bot, err := linebot.New(
		"b805bb7481fb026c1d42bec6aaec6f32",
		"2yQmWYV0iN8CGIP494SIrceebLoREAjIYSchlFTsjw1PuwDvRBd3nQV45powtMQmlm6Hg1uPlFO9UpIMEN2zkjepZklWyozKI74+555M+/4HCs0J9nj+cbrCxyaK4ny5L9HU4PvNzyrhwkfZH7KHawdB04t89/1O/w1cDnyilFU=",
	)

	if err != nil {
		log.Fatal(err)
	}

	// リクエストからBOTのイベントを取得
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidsignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		// イベントがメッセージの受信だった場合
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			// メッセージがテキスト形式の場合
			case *linebot.TextMessage:
				replyMessage := message.Text
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do()
				if err != nil {
					log.Print(err)
				}
			}

		}
	}
}
