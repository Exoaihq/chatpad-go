package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type Configuration struct {
	APIKey    string
	BasePath  string
}

type ChatCompletionRequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequest struct {
	Model    string                      `json:"model"`
	Stream   bool                        `json:"stream"`
	Messages []ChatCompletionRequestMessage `json:"messages"`
}

type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func NewConfiguration(apiKey string, basePath string) *Configuration {
	return &Configuration{
		APIKey:    apiKey,
		BasePath:  basePath,
	}
}

func CreateChatCompletion(config *Configuration, request ChatCompletionRequest) (*ChatCompletionResponse, error) {
	client := &http.Client{}
	data, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/engines/davinci-codex/completions", config.BasePath), strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.APIKey))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to create chat completion")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response ChatCompletionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func main() {
	apiKey := "your_api_key"
	basePath := "https://api.openai.com"
	config := NewConfiguration(apiKey, basePath)

	request := ChatCompletionRequest{
		Model:  "text-davinci-002",
		Stream: false,
		Messages: []ChatCompletionRequestMessage{
			{
				Role:    "user",
				Content: "hello",
			},
		},
	}

	response, err := CreateChatCompletion(config, request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:", response)
}