package sendler

import (
	"main/internal/config"
	"main/internal/logger"
	"main/internal/messages"
	"net/http"
	"net/url"
	"strings"
)

func SendServer(site, message string) {
	if config.Server != "" {
		m := messages.Message{
			Site: site,
			Log:  message,
		}
		reqBody, err := m.Encode()
		if err != nil {
			logger.SendError(err)
		}
		request, err := http.NewRequest(http.MethodPost, "http://"+config.Server+"/canary", strings.NewReader(string(reqBody)))
		if err != nil {
			logger.SendError(err)
		}
		client := http.DefaultClient
		response, err := client.Do(request)
		if err != nil {
			logger.SendError(err)
			SendTg(site + " : " + message)
		}
		defer response.Body.Close()
	} else {
		SendTg(site + " : " + message)
	}
}

func SendTg(message string) {
	message = url.QueryEscape(message)
	get, err := http.Get(config.TgGroup + message)
	if err != nil {
		logger.SendError(err)
	}
	get.Body.Close()
}
