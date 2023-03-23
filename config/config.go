package config

import (
	"io"
	"os"
	"sync"
)

var readOnce = &sync.Once{}
var apiKey string

func APIKey(path ...string) string {
	defPath := "apikey.txt"
	if len(path) > 0 {
		defPath = path[0]
	}
	readOnce.Do(func() {
		file, err := os.Open(defPath)
		if err != nil {
			panic("read config fail: " + err.Error())
		}
		bt, err := io.ReadAll(file)
		if err != nil {
			panic("read config fail: " + err.Error())
		}
		apiKey = string(bt)
	})
	return apiKey
}
