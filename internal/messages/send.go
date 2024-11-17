package messages

import (
	"encoding/json"
	"errors"
)

type Message struct {
	Site string `json:"site"`
	Log  string `json:"log"`
}

func (m *Message) Encode() ([]byte, error) {
	if len(m.Site) == 0 {
		return nil, errors.New("empty site")
	}
	if len(m.Log) == 0 {
		return nil, errors.New("empty log")
	}

	marshal, err := json.Marshal(&m)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

func (m *Message) Decode(body []byte) (Message, error) {
	err := json.Unmarshal(body, m)
	if err != nil {
		return *m, err
	}
	return *m, nil
}
