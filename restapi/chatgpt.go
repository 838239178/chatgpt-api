package restapi

import (
	"bytes"
	"chatgpt-api/model"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	bastUrl        = "https://api.openai.com/v1"
	chatCompletion = "/chat/completions"
)

var client = &http.Client{
	Transport: &http.Transport{
		Proxy: func(*http.Request) (*url.URL, error) {
			proxyUrl, ok := os.LookupEnv("PROXY_URL")
			if !ok {
				proxyUrl = "http://127.0.0.1:7890"
			}
			return url.Parse(proxyUrl)
		},
	},
}

func ChatCompletion(req *model.GPTRequest, apiKey string) (*model.GPTResponse, error) {
	bt, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	hreq, err := http.NewRequest(http.MethodPost, bastUrl+chatCompletion, bytes.NewBuffer(bt))
	if err != nil {
		return nil, err
	}
	hreq.Header.Set("Content-Type", "application/json")
	hreq.Header.Set("Authorization", fmt.Sprint("Bearer ", apiKey))
	resp, err := client.Do(hreq)
	if err != nil {
		return nil, err
	}
	bt, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(string(bt))
	}
	var res model.GPTResponse
	err = json.Unmarshal(bt, &res)
	return &res, err
}
