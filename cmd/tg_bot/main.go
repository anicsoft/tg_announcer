package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	token     = "TELEGRAM_TOKEN"
	webAppURL = "WEB_APP_URL"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv(token))
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Структура для обработки сообщений от пользователя
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // игнорировать любые неподдерживаемые типы обновлений
			continue
		}

		msg := update.Message
		chatID := msg.Chat.ID
		text := msg.Text

		switch text {
		case "/start":
			startCommand(chatID, bot)

			//formButton := tgbotapi.NewInlineKeyboardButtonURL("Fill in form", os.Getenv(webAppURL)+"/form")
			//formKeyboard := tgbotapi.NewInlineKeyboardMarkup(
			//	tgbotapi.NewInlineKeyboardRow(formButton),
			//)
			//
			//msg = tgbotapi.NewMessage(chatID, "Button will appear below")
			//msg.ReplyMarkup = formKeyboard
			//
			//_, err = bot.Send(msg)
			//if err != nil {
			//	log.Println(err)
			//}
		case "Hey":
			greet(chatID, update, bot)
		case "Share location":
			shareLocation(update, bot)
		}
	}
}

func startCommand(chatID int64, bot *tgbotapi.BotAPI) {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Hey"),
			tgbotapi.NewKeyboardButton("Share location"),
		),
	)
	keyboard.OneTimeKeyboard = true

	msg := tgbotapi.NewMessage(chatID, "Welcome")
	msg.ReplyMarkup = keyboard

	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func greet(chatID int64, update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Hey, %s!", update.Message.From))
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func shareLocation(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please share your location")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButtonLocation("Share Location"),
		),
	)

	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
