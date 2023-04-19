package trans

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

type Translator interface {
	// Translate returns the translation of the given string.
	// source, sourceLang, targetLang
	Translate(string, string, string) (string, error)
}

func GetTranslator(translatorType string) Translator {
	switch translatorType {
	case "google":
		return NewGoogleTranslator()
	default:
		return nil
	}
}

func getHttpClient() *http.Client {
	httpProxy := os.Getenv("http_proxy")
	proxyUrl, err := url.Parse(httpProxy)
	if err != nil || httpProxy == "" {
		log.Printf("no proxy detected, using default http client")
		return &http.Client{}
	}
	log.Printf("proxy detected: %v", proxyUrl)
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	return httpClient
}
