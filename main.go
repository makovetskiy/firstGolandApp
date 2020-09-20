package main

import (
	"fmt"
	"log"
	"net/http"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	botgolang "github.com/mail-ru-im/bot-golang"
)

//weatherApiKey := "83f46135a557b22219f927bb45ef3ad7"

func main() {
	println("Start...")
	//sendMessage("TEST TEXT")
	sendICQ("icq bot send")
}

func tgBotInit() {
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI("1143584598:AAG0xwp1M3hKhexHfEfbMSLaJhQMbLhuZ-8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	upd, _ := bot.GetUpdatesChan(ucfg)
	// читаем обновления из канала
	for {
		select {
		case update := <-upd:
			// Пользователь, который написал боту
			UserName := update.Message.From.UserName

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID

			// Текст сообщения
			Text := update.Message.Text

			if Text == "/weather_now" {
				weather := "⛅20C. Розовка, Запорожская область, Украина.  пятница 18:00"
				msg := tgbotapi.NewMessage(ChatID, weather)
				bot.Send(msg)
			} else {
				log.Printf("[%s] %d %s", UserName, ChatID, Text)
			}

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

		}

	}
}
func sendMessage(text string) {
	OllieWilliamsBot := "1143584598:AAG0xwp1M3hKhexHfEfbMSLaJhQMbLhuZ-8"
	weather := "⛅20C. Розовка, Запорожская область, Украина.  пятница 18:00"
	println(text)
	println(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=347034164&text=%s", OllieWilliamsBot, weather))
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=347034164&text=%s", OllieWilliamsBot, weather))
	if err != nil {
		println(err)
	}
	println(resp)
	defer resp.Body.Close()
}

func sendICQ(text string) {
	bot, err := botgolang.NewBot("001.0165450212.4268689202:756422445")
	if err != nil {
		log.Println("wrong token")
	}

	message := bot.NewTextMessage("671757227@agent.chat", "text")
	message.Send()
}
