package main

import (
	"log"
	"main/internal/canary"
	"main/internal/config"
	"main/internal/sendler"
	"path/filepath"
	"strings"
)

func main() {
	config.Config()

	files, err := filepath.Glob(config.LogPath + "*" + config.LogFileName)
	//glob, err := filepath.Glob("bin/" + "*" + config.LogFileName)
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

	for _, s := range files {

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
		strParse := <-ch
		if strParse.StatusCode >= 500 && strParse.StatusCode < 600 {
			if len(strParse.Agent) > 3 && !strings.Contains(strParse.Agent, "bot") {
				sendler.Send(strParse.Site, strParse.StrLogRaw)
			}
		}
	}

}
