package gotelegram

import "sync"

//Router is handler for muxing chat contains bot commands
type Router struct {
	handler         map[string]Handler
	defaultHandler  Handler
	notFoundHandler Handler
	mutex           sync.RWMutex
}

//NewRouter creates new mux handler
func NewRouter() *Router {
	return &Router{
		handler: make(map[string]Handler),
	}
}

//DefaultHandler assigns default handler for handling normal incoming messages
func (r *Router) DefaultHandler(handler Handler) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.defaultHandler = handler
}

//NotFoundHandler assigns not found handler for handling unhandled bot command
func (r *Router) NotFoundHandler(handler Handler) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.notFoundHandler = handler
}

//HandleCommand assignes handler for command
//if command is duplicated last assigned one will be used
func (r *Router) HandleCommand(command string, handler Handler) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.handler[command] = handler
}

//Handle is Mux Handler
func (r *Router) Handle(res MessageSender, update *Update) {
	if len(update.Message.Entities) > 0 {
		if update.Message.Entities[0].Type == "bot_command" {
			command := update.Message.Text[:update.Message.Entities[0].Length]
			r.mutex.RLock()
			if h := r.handler[command]; h != nil {
				r.mutex.RUnlock()
				h.Handle(res, update)
			} else {
				r.mutex.RUnlock()
				if r.notFoundHandler != nil {
					r.notFoundHandler.Handle(res, update)
				} else if r.defaultHandler != nil {
					// handle message using defaultHandler
					// if notFoundHandler is not setted up
					r.defaultHandler.Handle(res, update)
				}
			}
			return
		}
	}
	if r.defaultHandler != nil {
		r.defaultHandler.Handle(res, update)
	}
}
