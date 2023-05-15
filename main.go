package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	// Get config parameters
	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	telegramApiUrl := "https://api.telegram.org/bot"
	offset := 0
	webAppUrl := "https://www.google.ru/"
	replyKeyboardMarkup := getReplyKeyboardMarkup(webAppUrl)

	for {
		updates, err := getUpdates(telegramBotToken, telegramApiUrl, offset)
		if err != nil {
			log.Println("Something went wrong: ", err.Error())
		}

		for _, update := range updates {
			err = respond(telegramBotToken, telegramApiUrl, update, replyKeyboardMarkup)
			if err != nil {
				log.Println("respond doesn't work: ", err.Error())
			}
			offset = update.UpdateId + 1
		}

		// Timeout
		time.Sleep(time.Second * 1)

		fmt.Println(updates)
	}
}

// Getting new updates from Telegram Bot API
func getUpdates(botToken string, apiUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(apiUrl + botToken + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}

	return restResponse.Result, nil
}

// Make and send responses
func respond(botToken string, apiUrl string, update Update, replyKeyboardMarkup ReplyKeyboardMarkup) error {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.Id
	botMessage.Text = "You asked: " + update.Message.Text
	botMessage.ReplyMarkup = replyKeyboardMarkup

	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}

	resp, err := http.Post(apiUrl+botToken+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

func getReplyKeyboardMarkup(webAppUrl string) ReplyKeyboardMarkup {
	webAppInfo := WebAppInfo{
		Url: webAppUrl,
	}

	keyboardButton := KeyboardButton{
		Text:   "Open app",
		WebApp: webAppInfo,
	}

	keyboardButtonRow := []KeyboardButton{keyboardButton}

	replyKeyboardMarkup := ReplyKeyboardMarkup{
		Keyboard: [][]KeyboardButton{keyboardButtonRow},
	}

	return replyKeyboardMarkup
}
