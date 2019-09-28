package gotelegram

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"
)

type getResult struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

//PollingListener is listener for long polling update
type PollingListener struct {
	Client  *http.Client
	Token   string
	Timeout int
	s       chan interface{} //Channel s will be closed on shutdown
}

//NewPollingListener returns initialized polling listener
func NewPollingListener() PollingListener {
	return PollingListener{
		s: make(chan interface{}),
	}
}

//ListenAndServe listen updates using long polling
func (p PollingListener) ListenAndServe(handler Handler) {
	pollingURL := "https://" + path.Join("api.telegram.org/", "bot"+p.Token, "getUpdates")
	pollingURL += "?timeout=" + strconv.Itoa(p.Timeout)
	var offset = -1
mainLoop:
	for {
		select {
		case <-p.s:
			break mainLoop
		default:
			if resp, err := p.Client.Get(pollingURL + "&offset=" + strconv.Itoa(offset)); err == nil {
				var result getResult
				err = json.NewDecoder(resp.Body).Decode(&result)
				if err != nil {
					log.Println(err)
					continue
				}
				if !result.OK {
					log.Println("Result not OK")
					continue
				}
				for _, u := range result.Result {
					offset = u.UpdateID + 1
					go handler.Handle(&u)
				}
			} else {
				log.Println(err)
			}
		}
	}
}

//Shutdown shuts down listener after current polling is done
func (p *PollingListener) Shutdown() {
	close(p.s)
}
