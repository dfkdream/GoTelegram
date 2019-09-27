package gotelegram

//Update represents an incoming update
type Update struct {
	UpdateID          int     `json:"update_id"`
	Message           Message `json:"message"`
	EditedMessage     Message `json:"edited_message"`
	ChannelPost       Message `json:"channel_post"`
	EditedChannelPost Message `json:"edited_channel_post"`
	//TODO: Implement following fields
	/*
		InlineQuery        InlineQuery        `json:"inline_query"`
		ChosenInlineResult ChosenInlineResult `json:"chosen_inline_result"`
		CallbackQuery      CallbackQuery      `json:"callback_query"`
		ShippingQuery      ShippingQuery      `json:"shipping_query"`
		PreCheckoutQuery   PreCheckoutQuery   `json:"pre_checkout_query"`
		Poll               Poll               `json:"poll"`
	*/
}

//Message represents a message
type Message struct {
	MessageID int  `json:"message_id"`
	From      User `json:"from"`
	Date      int  `json:"date"`
	Chat      Chat `json:"chat"`
	//TODO: more fields to be implemented...
}

//User represents a Telegram user or bot
type User struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

//Chat represents a chat
type Chat struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	//TODO: more fields to be implemented...
}
