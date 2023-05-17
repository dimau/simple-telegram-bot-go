package main

// Update is a Telegram object that the handler receives every time user interacts with the bot.
type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message is a Telegram object that can be found in an update.
type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

// A Telegram Chat indicates the conversation to which the message belongs.
type Chat struct {
	Id int `json:"id"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatId int    `json:"chat_id"`
	Text   string `json:"text"`
	//ReplyMarkup ReplyKeyboardMarkup `json:"reply_markup"`
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup"`
}

type ReplyKeyboardMarkup struct {
	Keyboard [][]KeyboardButton `json:"keyboard"`
}

type KeyboardButton struct {
	Text   string     `json:"text"`
	WebApp WebAppInfo `json:"web_app"`
}

type WebAppInfo struct {
	Url string `json:"url"`
}

type InlineKeyboardMarkup struct {
	Keyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text   string     `json:"text"`
	WebApp WebAppInfo `json:"web_app"`
}
