package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Config struct {
	DefaultModel      string           `json:"defaultModel"`
	DefaultType       string           `json:"defaultType"`
	DefaultAuth       string           `json:"defaultAuth"`
	DefaultBase       string           `json:"defaultBase"`
	DefaultVersion    string           `json:"defaultVersion"`
	DefaultKey        string           `json:"defaultKey"`
	AvailableModels   []AvailableModel `json:"availableModels"`
	WritingCharacters []WritingCharacter `json:"writingCharacters"`
	WritingTones      []string         `json:"writingTones"`
	WritingStyles     []string         `json:"writingStyles"`
	WritingFormats    []WritingFormat  `json:"writingFormats"`
	ShowDownloadLink  bool             `json:"showDownloadLink"`
	AllowDarkModeToggle bool           `json:"allowDarkModeToggle"`
	AllowSettingsModal bool            `json:"allowSettingsModal"`
	AllowDatabaseModal bool            `json:"allowDatabaseModal"`
	ShowTwitterLink   bool             `json:"showTwitterLink"`
	ShowFeedbackLink  bool             `json:"showFeedbackLink"`
}

type AvailableModel struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type WritingCharacter struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type WritingFormat struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

var config Config

func loadConfig() (Config, error) {
	response, err := http.Get("config.json")
	if err != nil {
		return Config{}, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func main() {
	loadedConfig, err := loadConfig()
	if err != nil {
		panic(err)
	}
	config = loadedConfig
}