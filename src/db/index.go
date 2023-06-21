package main

import (
	"time"

	"github.com/tevino/abool"
	"github.com/tevino/dexie"
)

type Chat struct {
	ID          string
	Description string
	TotalTokens int
	CreatedAt   time.Time
}

type Message struct {
	ID        string
	ChatID    string
	Role      string
	Content   string
	CreatedAt time.Time
}

type Prompt struct {
	ID        string
	Title     string
	Content   string
	CreatedAt time.Time
}

type Settings struct {
	ID             string
	OpenAiApiKey   *string
	OpenAiModel    *string
	OpenAiApiType  *string
	OpenAiApiAuth  *string
	OpenAiApiBase  *string
	OpenAiApiVersion *string
}

type Database struct {
	Chats    *dexie.Table
	Messages *dexie.Table
	Prompts  *dexie.Table
	Settings *dexie.Table
}

func NewDatabase() *Database {
	db := dexie.New("chatpad")
	chats := db.Table("chats", "id, createdAt")
	messages := db.Table("messages", "id, chatId, createdAt")
	prompts := db.Table("prompts", "id, createdAt")
	settings := db.Table("settings", "id")

	db.On("populate", func() {
		settings.Add(Settings{
			ID:             "general",
			OpenAiModel:    abool.String(config.defaultModel),
			OpenAiApiType:  abool.String(config.defaultType),
			OpenAiApiAuth:  abool.String(config.defaultAuth),
			OpenAiApiKey:   abool.String(config.defaultKey),
			OpenAiApiBase:  abool.String(config.defaultBase),
			OpenAiApiVersion: abool.String(config.defaultVersion),
		})
	})

	return &Database{
		Chats:    chats,
		Messages: messages,
		Prompts:  prompts,
		Settings: settings,
	}
}

var db = NewDatabase()