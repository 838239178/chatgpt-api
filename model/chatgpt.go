package model

import "encoding/gob"

func init() {
	gob.Register([]*GPTMessage{})
}

const (
	ModelChat3_5      = "gpt-3.5-turbo"
	ModelChat3_5_0301 = "gpt-3.5-turbo-0301"
	ModelCodeCushman  = "code-cushman-001"
)

const (
	RoleSystem    = "system"
	RoleAssistant = "assistant"
	RoleUser      = "user"
)

const (
	FinishByStop   = "stop"
	FinishByLen    = "length"
	FinishByFilter = "content-filter"
	FinishByNull   = "null"
)

type GPTMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GPTRequest struct {
	Model            string        `json:"model"`
	Messages         []*GPTMessage `json:"messages"`
	Temperature      float32       `json:"temperature"`
	FrequencyPenalty float32       `json:"frequency_penalty"`
	PresencePenalty  float32       `json:"presence_penalty"`
	Stream           bool          `json:"stream"`
}

type GPTUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type GPTChoice struct {
	Message      *GPTMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
	Index        int         `json:"index"`
}

type GPTResponse struct {
	Id      string       `json:"id"`
	Object  string       `json:"object"`
	Created int64        `json:"created"`
	Model   string       `json:"model"`
	Usage   *GPTUsage    `json:"usage"`
	Choices []*GPTChoice `json:"choices"`
}
