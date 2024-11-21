package cliq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CLIQ_CONFIG struct {
	WebhookEndpoint string `json:"webhookendpoint"`
	ZAPIKey         string `json:"zapikey"`
}

func (c CLIQ_CONFIG) CliqEndpoint() string {
	return fmt.Sprintf("%s?zapikey=%s", c.WebhookEndpoint, c.ZAPIKey)
}

func (c CLIQ_CONFIG) SendMessageToCliq(message string) error {
	payload, _ := json.Marshal(map[string]string{
		"text": message,
	})
	resp, err := http.Post(c.CliqEndpoint(), "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
