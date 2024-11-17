package canary

import (
	"bufio"
	"errors"
	"io"
	"main/internal/logger"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Canary struct {
	Site       string
	Ip         string
	Identity   string
	User       string
	Date       string
	Time       string
	Timezone   string
	Method     string
	URI        string
	Protocol   string
	StatusCode int
	Bytes      string
	Referer    string
	Agent      string
	StrLogRaw  string
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

	status, _ := strconv.Atoi(match[10])
	return Canary{
		Ip:         match[1],
		Identity:   match[2],
		User:       match[3],
		Date:       match[4],
		Time:       match[5],
		Timezone:   match[6],
		Method:     match[7],
		URI:        match[8],
		Protocol:   match[9],
		StatusCode: status,
		Bytes:      match[11],
		Referer:    match[12],
		Agent:      match[13],
		StrLogRaw:  str,
	}
}
