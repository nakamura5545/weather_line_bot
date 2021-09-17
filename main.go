package main

import (
	"log"

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

}
