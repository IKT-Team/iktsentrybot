package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const ENDPOINT_TEMPLATE = "https://api.telegram.org/bot%s/sendMessage"
const MESSAGE_TEMPLATE = "*\\[Sentry\\]* %s %s: *%s*\n```\n%s\n```\n[View in Sentry](%s)"

type TelegramResponse struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description,omitempty"`
}

func SendTelegramMessage(s *SentryEvent) error {
	text := fmt.Sprintf(
		MESSAGE_TEMPLATE,
		escape(s.ProjectName),
		escape(s.Level),
		escape(strOrDefault(s.Event.Title, "no title")),
		escape(strOrDefault(s.Message, "no message")),
		escape(s.URL),
	)
	chat_id := os.Getenv("CHAT_ID")
	endpoint := fmt.Sprintf(ENDPOINT_TEMPLATE, os.Getenv("BOT_TOKEN"))
	data := url.Values{"chat_id": {chat_id}, "text": {text}, "parse_mode": {"MarkdownV2"}}
	httpResp, err := http.PostForm(endpoint, data)
	if err != nil {
		return fmt.Errorf("tg request failed with error %w", err)
	}
	defer httpResp.Body.Close()
	fmt.Printf("tg request status code %d\n", httpResp.StatusCode)
	tgResp := new(TelegramResponse)
	err = json.NewDecoder(httpResp.Body).Decode(tgResp)
	if err != nil {
		return fmt.Errorf("tg response decode failed with error %w", err)
	}
	if !tgResp.Ok {
		return fmt.Errorf("tg request is unsuccessful: %s", tgResp.Description)
	}
	return nil
}

func strOrDefault(str, def string) string {
	if len(str) == 0 {
		return def
	}
	return str
}

func escape(str string) string {
	result := make([]rune, 0, len(str)*2)
	for _, r := range str {
		if r >= 1 && r <= 126 {
			result = append(result, '\\')
		}
		result = append(result, rune(r))
	}
	return string(result)
}
