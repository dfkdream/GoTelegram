package gotelegram

import (
	"bytes"
	"encoding/json"
	"net/http"
	"path"
)

//MessageSender sends message
type MessageSender struct {
	token string
}

//SendMessageParams contains parameters for SendMessage
type SendMessageParams struct {
	ChatID              int    `json:"chat_id"`
	Text                string `json:"text"`
	ParseMode           string `json:"parse_mode"`
	DisableWebPreview   bool   `json:"disable_web_page_preview"`
	DisableNotification bool   `json:"disable_notification"`
	ReplyToMessageID    int    `json:"reply_to_message_id"`
	//TODO: Implement reply_markup
}

//SendMessage is method for sending message
func (r MessageSender) SendMessage(params SendMessageParams) (*Message, error) {
	data, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post("https://"+path.Join("api.telegram.org", "bot"+r.token, "sendMessage"), "application/json", bytes.NewBuffer(data))
	var res Message
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
