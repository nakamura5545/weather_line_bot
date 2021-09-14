package main

import (
	"log"
	"os"

	"github.com/nakamura5545/weather_line_bot/eventhandler"
	"github.com/nakamura5545/weather_line_bot/infrastracture/line"
	"github.com/nakamura5545/weather_line_bot/weather"
)

func main() {
	client := line.New()
	result := weather.GetWeather()
	message := line.TextMessage(result)
	// テキストメッセージを友達登録しているユーザー全員に配信する
	if _, err := client.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}
	// 連絡をしたかどうかのリマインダーを送る(水曜日に出力する)
	if eventhandler.GetWeek() == 3 {
		remind := line.TextMessage(os.Getenv("REMIND_MESSAGE"))
		if _, err := client.BroadcastMessage(remind).Do(); err != nil {
			log.Fatal(err)
		}
	}

}
