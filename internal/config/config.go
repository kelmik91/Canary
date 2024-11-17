package config

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
)

var TgGroup string
var TgMy string
var Server string

var LogPath = "/var/www/fastuser/data/logs/"
var LogFileName = "-backend.access.log"

func Config() {
	dev := flag.Bool("dev", false, "development mode")
	flag.Parse()
	if *dev {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	} else {
		err := godotenv.Load("/etc/canary/.config")
		if err != nil {
			panic(err)
		}
	}

	token := os.Getenv("BOT_TOKEN")
	myID := os.Getenv("MY_ID")
	groupID := os.Getenv("GROUP_ID")

	TgGroup = "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + groupID + "&text="
	TgMy = "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + myID + "&text="

	Server = os.Getenv("SERVER")
}
