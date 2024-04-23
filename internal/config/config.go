package config

import (
	"github.com/joho/godotenv"
	"os"
)

var TgGroup string
var TgMy string
var Server string

var Path = "/var/www/fastuser/data/logs/"
var LogFileName = "-backend.access.log"

func Config() {
	err := godotenv.Load("/etc/canary/.config")
	if err != nil {
		panic(err)
	}
	token := os.Getenv("BOT_TOKEN")

	myID := os.Getenv("MY_ID")
	groupID := os.Getenv("GROUP_ID")

	TgGroup = "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + groupID + "&text="
	TgMy = "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + myID + "&text="
}
