package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	chatID := useChatID(r.URL.Path)
	if chatID != "" {
		fmt.Fprintf(w, "Chat ID: %s", chatID)
	} else {
		fmt.Fprint(w, "No chat ID found")
	}
}

func useChatID(path string) string {
	regex := regexp.MustCompile(`/chats/([^/]+)`)
	matches := regex.FindStringSubmatch(path)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}