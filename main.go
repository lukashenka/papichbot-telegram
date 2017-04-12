package main

import (
	"gopkg.in/telegram-bot-api.v4"
	//"github.com/lukashenka/papichizator"
	"log"
	"github.com/lukashenka/papichizator"
	"crypto/md5"
	"encoding/hex"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("333292218:AAGzHkaogC5eF0v1IMG3IuQDMSNgplkekFg")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	inlineConfig := tgbotapi.InlineConfig{}
	response, err := bot.AnswerInlineQuery(inlineConfig)
	log.Println(response)

	for update := range updates {
		if update.InlineQuery == nil {
			continue
		}

		papich := papichizator.Papichizator{}
		articles := make([]interface{}, 5)
		message:=update.InlineQuery.Query
		for i := 0; i < 5; i++ {

			papichtext := papich.Papichize(message)
			log.Println(message+ " - "+papichtext)
			id:=GetMD5Hash(string(i)+"id")
			article := tgbotapi.NewInlineQueryResultArticle(id, "По папански", papichtext)
			article.Description = papichtext
			articles[i] = article
		}
		inlineConf := tgbotapi.InlineConfig{
			InlineQueryID: update.InlineQuery.ID,
			IsPersonal:    false,
			CacheTime:     0,
			Results:       articles,

		}

		if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
			log.Println(err)
		}
	}
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}