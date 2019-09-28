package gotelegram

// Handler handles update
type Handler interface {
	Handle(res MessageSender, update *Update)
}

// HandlerFunc type is an adaptor to allow the use of
// ordinary functions as handlers.
type HandlerFunc func(res MessageSender, update *Update)

// Handle calls f(update)
func (f HandlerFunc) Handle(res MessageSender, update *Update) {
	f(res, update)
}
