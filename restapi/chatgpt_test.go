package restapi

import (
	"chatgpt-api/config"
	"chatgpt-api/model"
	"testing"
)

func TestChatCompletion(t *testing.T) {
	resp, err := ChatCompletion(&model.GPTRequest{
		Model: model.ModelChat3_5,
		Messages: []*model.GPTMessage{
			{Role: model.RoleSystem, Content: "You are a helpful assistant created by SJH not by OpenAI"},
			{Role: model.RoleUser, Content: "Who are you?"},
		},
	}, config.APIKey("../apikey.txt"))
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range resp.Choices {
		t.Log(v.Message.Content)
	}
}

func TestChatCompletionStream(t *testing.T) {
	resp, err := ChatCompletion(&model.GPTRequest{
		Model:  model.ModelChat3_5,
		Stream: true,
		Messages: []*model.GPTMessage{
			{Role: model.RoleSystem, Content: "You are a helpful assistant of chatgpt-api"},
			{Role: model.RoleUser, Content: "Who are you?"},
		},
	}, config.APIKey("../apikey.txt"))
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range resp.Choices {
		t.Log(v.Message.Content)
	}
}
