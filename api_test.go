package main

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"sync"
	"testing"
)

var basicAuth = "Basic " + base64.StdEncoding.EncodeToString([]byte("windows:password"))

func getHistoryAPI() {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:9080/v1/chatHistory", nil)
	req.Header.Set("Authorization", basicAuth)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err)
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		println(err)
	}
	println(string(content))
}

func postChatAPI() {
	resp, err := http.Post("http://localhost:9080/v1/chat", "application/json", bytes.NewBufferString(`{"msg":"1+1=?"}`))
	if err != nil {
		println(err)
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		println(err)
	}
	println(string(content))
}

func TestChatHistory(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getHistoryAPI()
		}()
	}
	wg.Wait()
}

func TestChat(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			postChatAPI()
		}()
	}
	wg.Wait()
}
