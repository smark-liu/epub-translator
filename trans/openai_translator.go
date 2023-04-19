package trans

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type OpenAITranslator struct {
	apiKey     string
	httpClient *http.Client
}

func (o *OpenAITranslator) Translate(source, sourceLang, targetLang string) (string, error) {
	url := "https://api.openai.com/v1/engines/davinci-codex/completions"
	requestBody := fmt.Sprintf(`{
        "prompt": "%s",
        "max_tokens": 60,
        "temperature": 0.7,
        "n": 1,
        "stop": ["\n"]
    }`, source)

	req, err := http.NewRequest("POST", url, strings.NewReader(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", o.apiKey))

	resp, err := o.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var respBody struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return "", err
	}

	if len(respBody.Choices) == 0 {
		return "", errors.New("no response from OpenAI API")
	}

	return respBody.Choices[0].Text, nil
}
