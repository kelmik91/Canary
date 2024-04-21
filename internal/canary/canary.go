package canary

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"main/internal/logger"
	"os"
	"regexp"
	"time"
)

type Canary struct {
	Site       string
	Method     string
	URI        string
	StatusCode string
	StrLog     string
}

func Tail(site, filename string, out chan Canary) {
	f, err := os.Open(filename)
	if err != nil {
		logger.Error(err)
	}
	defer f.Close()
	r := bufio.NewReader(f)
	info, err := f.Stat()
	if err != nil {
		logger.Error(err)
	}
	oldSize := info.Size()
	for {
		for line, prefix, err := r.ReadLine(); err != io.EOF; line, prefix, err = r.ReadLine() {
			if prefix {
				//out <- fmt.Sprint(string(line))
				logger.Error(errors.New("пришла слишком длинная строка"))
			} else {
				//out <- fmt.Sprintln(string(line))
				parsed := Parse(string(line))
				parsed.Site = site
				out <- parsed
			}
		}
		pos, err := f.Seek(0, io.SeekCurrent)
		if err != nil {
			logger.Error(err)
		}
		for {
			time.Sleep(time.Second)
			newInfo, err := f.Stat()
			if err != nil {
				logger.Error(err)
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
				logger.Error(err)
			}
			r = bufio.NewReader(f)
			info, err = f.Stat()
			if err != nil {
				logger.Error(err)
			}
			oldSize = info.Size()
		}
	}
}

func Parse(str string) Canary {

	//7 method
	//8 URI
	//10 status code
	//['ip'] =  [1];
	//['identity'] =  [2];
	//['user'] =  [3];
	//['date'] =  [4];
	//['time'] =  [5];
	//['timezone'] = [6];
	//['method'] =  [7];
	//['path'] = [8];
	//['protocol'] = [9];
	//['status'] = [10];
	//['bytes'] = [11];
	//['referer'] = [12];
	//['agent'] = [13];

	re := regexp.MustCompile(`(\S+) (\S+) (\S+) \[([^:]+):(\d+:\d+:\d+) ([^\]]+)\] \"(\S+) (.*?) (\S+)\" (\S+) (\S+) (\".*?\") (\".*?\")`)
	match := re.FindStringSubmatch(str)
	for i, s := range match {
		fmt.Println(i, s)
	}

	return Canary{
		Method:     match[7],
		URI:        match[8],
		StatusCode: match[10],
		StrLog:     str,
	}
}
