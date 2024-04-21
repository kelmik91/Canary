package logger

import (
	"main/internal/config"
	"net/http"
	"net/url"
)

func Error(err error) {
	SendError(err)
}

func SendError(err error) {
	str := url.QueryEscape(err.Error())
	_, err = http.Get(config.TgMy + str)
	panic(err)
}
