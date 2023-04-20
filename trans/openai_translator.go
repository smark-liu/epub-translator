package trans

import (
	"encoding/json"
	"fmt"
	"github.com/smark-d/epub-translator/common"
	"io"
	"net/http"
	"net/url"
	"strings"
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
	// 提取翻译结果
	translations := response["choices"].([]interface{})[0].(map[string]interface{})["text"].(string)
	return translations, nil
}

func (o *OpenAITranslator) sendRequest(model, prompt string) (map[string]interface{}, error) {
	// construct Request Data
	data := url.Values{}
	data.Set("model", model)
	data.Set("prompt", prompt)
	data.Set("temperature", "0.7")
	data.Set("max_tokens", "60")

	// construct Request
	url, err := url.JoinPath(o.apiUrl, "v1/completions")
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+o.apiKey)

	// Do Request
	resp, err := o.httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read Response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
