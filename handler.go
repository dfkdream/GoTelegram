package gotelegram

// Handler handles update
type Handler interface {
	Handle(update *Update)
}

// HandlerFunc type is an adaptor to allow the use of
// ordinary functions as handlers.
type HandlerFunc func(update *Update)

// Handle calls f(update)
func (f HandlerFunc) Handle(update *Update) {
	f(update)
}
