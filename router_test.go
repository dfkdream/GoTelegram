package gotelegram

import (
	"testing"
)

func TestRouter(t *testing.T) {
	r := NewRouter()

	r.DefaultHandler(HandlerFunc(func(res MessageSender, u *Update) {
		if u.Message.Text != "testing" {
			t.Error("Expected 'testing' but " + u.Message.Text + " received")
		}
	}))

	r.NotFoundHandler(HandlerFunc(func(res MessageSender, u *Update) {
		if u.Message.Text != "/notavailable" {
			t.Error("Expected '/notavailable' but " + u.Message.Text + " received")
		}
	}))

	r.HandleCommand("/test", HandlerFunc(func(res MessageSender, u *Update) {
		if u.Message.Text != "/test hello" {
			t.Error("Expecte '/test hello' but " + u.Message.Text + " received")
		}
	}))

	r.Handle(MessageSender{}, &Update{
		Message: Message{
			Text: "testing",
		},
	})

	r.Handle(MessageSender{}, &Update{
		Message: Message{
			Text: "/test hello",
			Entities: []MessageEntity{
				MessageEntity{
					Type:   "bot_command",
					Offset: 0,
					Length: 5,
				},
			},
		},
	})

	r.Handle(MessageSender{}, &Update{
		Message: Message{
			Text: "/notavailable",
			Entities: []MessageEntity{
				MessageEntity{
					Type:   "bot_command",
					Offset: 0,
					Length: 13,
				},
			},
		},
	})
}
