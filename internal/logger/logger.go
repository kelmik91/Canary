package logger

import (
	"log"
	"main/internal/config"
	"net/http"
	"net/url"
)

func Error(err error) {
	SendError(err)
}

func SendError(err error) {
	str := url.QueryEscape(err.Error())
	get, err := http.Get(config.TgMy + str)
	if err != nil {
		log.Println(err)
	}
	defer get.Body.Close()
}
