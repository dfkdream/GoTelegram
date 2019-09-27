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

//ListenAndServePolling listen updates using long polling
func ListenAndServePolling(token string, timeout int, handler Handler) {
	pollingURL := "https://" + path.Join("api.telegram.org/", "bot"+token, "getUpdates")
	var offset = -1
	for {
		if resp, err := http.Get(pollingURL + "?offset=" + strconv.Itoa(offset) + "&timeout=" + strconv.Itoa(timeout)); err == nil {
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
				go handler.Handle(u)
			}
		} else {
			log.Println(err)
		}
	}
}
