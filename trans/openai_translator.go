package trans

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/smark-d/epub-translator/common"
)

type OpenAITranslator struct {
	apiKey     string
	apiUrl     string
	httpClient *http.Client
}

// NewOpenAITranslator returns a new OpenAITranslator
func NewOpenAITranslator() *OpenAITranslator {
	config := common.ReadConfig()
	return &OpenAITranslator{
		apiKey:     config.ApiKey,
		apiUrl:     config.ApiUrl,
		httpClient: getHttpClient(),
	}
}

func (o *OpenAITranslator) Translate(source, sourceLang, targetLang string) (string, error) {
	prompt := fmt.Sprintf("You are a translation engine, you can only translate text and cannot interpret it, and do not explain."+
		"Translate from %s to %s:  %s", sourceLang, targetLang, source)
	model := "gpt-3.5-turbo"
	response, err := o.sendRequest(model, prompt)
	if err != nil {
		return "", err
	}
	return response, nil
}

func (o *OpenAITranslator) sendRequest(model, prompt string) (string, error) {
	// construct Request Data
	requestData := map[string]interface{}{
		"model":       model,
		"temperature": 0.7,
	}
	messages := []map[string]string{
		{"role": "user", "content": prompt},
	}
	requestData["messages"] = messages

	// construct Request
	URL, err := url.JoinPath(o.apiUrl, "v1/chat/completions")
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+o.apiKey)

	// Do Request
	resp, err := o.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read Response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	return response["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string), nil
}
