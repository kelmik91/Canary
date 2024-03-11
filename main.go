package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

var projectName string
var logPath string
var logFile string
var tgGroup string
var tgMy string

func init() {
	err := godotenv.Load()
	if err != nil {
		sendError(err)
	}
	token := os.Getenv("BOT_TOKEN")

	projectName = os.Getenv("PROJECT_NAME")

	logPath = os.Getenv("LOG_PATH")
	logFile = os.Getenv("LOG_ACCESS_FILE")

	myID := os.Getenv("MY_ID")
	groupID := os.Getenv("GROUP_ID")

	tgGroup = "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + groupID + "&text=" + url.QueryEscape(projectName+"\n")
	tgMy = "https://api.telegram.org/bot" + token + "/sendMessage?chat_id=" + myID + "&text=" + url.QueryEscape(projectName+"\n")
}

func main() {
	ch := make(chan string)

	go tail(logPath+projectName+logFile, ch)
	for {
		str := <-ch
		if strings.Contains(str, " 301 ") {
			continue
		}
		if strings.Contains(str, " 200 ") {
			continue
		}
		//strSlice := strings.Split(str, " ")
		//switch strSlice[3] {
		//case "200":
		//	continue
		//case "301":
		//	continue
		//}

		if errorRegexp, _ := regexp.MatchString(`\s5\d{2}\s`, str); errorRegexp {
			fmt.Print(str)
			str = url.QueryEscape(str)
			get, err := http.Get(tgGroup + str)
			if err != nil {
				sendError(err)
			}
			get.Body.Close()
		}
	}
}

func tail(filename string, out chan string) {
	f, err := os.Open(filename)
	if err != nil {
		sendError(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	info, err := f.Stat()
	if err != nil {
		sendError(err)
	}
	oldSize := info.Size()
	for {
		for line, prefix, err := r.ReadLine(); err != io.EOF; line, prefix, err = r.ReadLine() {
			if prefix {
				out <- fmt.Sprint(string(line))
			} else {
				out <- fmt.Sprintln(string(line))
			}
		}
		pos, err := f.Seek(0, io.SeekCurrent)
		if err != nil {
			sendError(err)
		}
		for {
			time.Sleep(time.Second)
			newInfo, err := f.Stat()
			if err != nil {
				sendError(err)
			}
			newSize := newInfo.Size()
			if newSize != oldSize {
				if newSize < oldSize {
					_, err := f.Seek(0, 0)
					if err != nil {
						return
					}
				} else {
					_, err := f.Seek(pos, io.SeekStart)
					if err != nil {
						return
					}
				}
				r = bufio.NewReader(f)
				oldSize = newSize
				break
			}
		}
		if time.Now().Day() != info.ModTime().Day() {
			f.Close()
			time.Sleep(time.Minute)
			f, err = os.Open(filename)
			if err != nil {
				sendError(err)
			}
			r = bufio.NewReader(f)
			info, err = f.Stat()
			if err != nil {
				sendError(err)
			}
			oldSize = info.Size()
		}
	}
}

func sendError(err error) {
	str := url.QueryEscape(err.Error())
	_, err = http.Get(tgMy + str)
	panic(err)
}
