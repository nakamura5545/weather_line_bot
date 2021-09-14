package line

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

func New() *linebot.Client {
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
	return bot
}
func TextMessage(message string) *linebot.TextMessage {
	return linebot.NewTextMessage(message)
}
