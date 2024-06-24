package main

import (
	"log"
	"main/internal/canary"
	"main/internal/config"
	"main/internal/sendler"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {
	config.Config()

	//glob, err := filepath.Glob(config.LogPath + "*" + config.LogFileName)
	glob, err := filepath.Glob("bin/" + "*" + config.LogFileName)
	if err != nil {
		log.Fatalln(err)
	}

	ch := make(chan canary.Canary)

	allowedSubstrings := []string{
		"test.",
		"tw1.su",
		"tw1.ru",
		"webtm.ru",
		"twc1.net",
	}

	for _, s := range glob {

		found := true
		for _, substr := range allowedSubstrings {
			if strings.Contains(s, substr) {
				found = false
				break
			}
		}

		if found {
			site := strings.TrimSuffix(s, config.LogFileName)
			site = strings.TrimPrefix(site, config.LogPath)
			go canary.Tail(site, s, ch)
		}
	}

	for {
		str := <-ch
		if str.StatusCode == http.StatusOK || str.StatusCode == http.StatusMovedPermanently {
			continue
		}

		sendler.SendServer(str.Site, str.StrLogRaw)
	}

}
