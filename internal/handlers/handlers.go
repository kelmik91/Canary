package handlers

import (
	"fmt"
	"io"
	"main/internal/canary"
	"main/internal/logger"
	"main/internal/messages"
	"net/http"
	"os"
)

func GetClient(w http.ResponseWriter, r *http.Request) {
	fileBytes, err := os.ReadFile("bin/canary.deb")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}

func PostHandlerCanary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		logger.Error(fmt.Errorf("server request method not allowed: %s", r.Method))
		return
	}

	body, _ := io.ReadAll(r.Body)
	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		logger.Error(fmt.Errorf("server request method not allowed: %s", r.Method))
		return
	}

	m := messages.Message{}
	messageJSON, _ := m.Decode(body)
	message := canary.Parse(messageJSON.Log)
	fmt.Println(messageJSON.Site)
	return
	if message.StatusCode != http.StatusOK && message.Method == http.MethodGet {
		//TODO дописать проверку запроса
		response, err := http.Get(messageJSON.Site)
		if err != nil {
			logger.Error(fmt.Errorf("request check status code failed: %s", err))
			return
		}
		defer response.Body.Close()
		if response.StatusCode != message.StatusCode {
			//TODO
			// тут будет запись в БД
			// передача в анализ количества ошибок
		}
	}
}
