package main

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/nakamura5545/weather_line_bot/weather"
)

func main() {
	// LINE Botクライアント生成する
	// BOT にはチャネルシークレットとチャネルトークンを環境変数から読み込み引数に渡す
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	// エラーに値があればログに出力し終了する
	if err != nil {
		log.Fatal(err)
	}
	// 連絡をしたかどうかのリマインダーを送る
	remind := linebot.NewTextMessage(os.Getenv("REMIND_MESSAGE"))
	if _, err := bot.BroadcastMessage(remind).Do(); err != nil {
		log.Fatal(err)
	}
	// weatherパッケージパッケージから天気情報の文字列を取得する
	result := weather.GetWeather()
	// テキストメッセージを生成する
	message := linebot.NewTextMessage(result)
	// テキストメッセージを友達登録しているユーザー全員に配信する
	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}

}
