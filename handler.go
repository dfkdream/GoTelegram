package gotelegram

//Handler handles update
type Handler interface {
	Handle(update *Update)
}
