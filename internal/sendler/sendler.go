package sendler

import (
	"main/internal/config"
	"main/internal/logger"
	"net/http"
	"net/url"
	"strings"
)

func SendServer(message string) {
	if config.Server == "" {
		request, err := http.NewRequest(http.MethodPost, config.Server, strings.NewReader(message))
		if err != nil {
			return
		}

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			logger.SendError(err)
			SendTg(message)
		}
		defer response.Body.Close()
	} else {
		SendTg(message)
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
