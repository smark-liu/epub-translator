package parser

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/smark-d/epub-translator/trans"
)

type TextParser struct {
	Path       string // epub file Path
	From       string // source language
	To         string // target language
	KeepOrigin bool   // keep the original file
	Translator trans.Translator
}

func (e *TextParser) Parse() (string, error) {

	// Read the text file
	text, err := os.ReadFile(e.Path)
	if err != nil {
		return "", err
	}

	// Split the text into segments
	segments := strings.Split(string(text), "\n")

	// Join the segments back into a single string
	translatedText := strings.Join(e.translateSegments(segments), "\n")

	// Write the translated text to a new file
	newName := strings.TrimSuffix(e.Path, filepath.Ext(e.Path)) + ".trans.txt"
	err = os.WriteFile(newName, []byte(translatedText), 0644)
	if err != nil {
		return "", err
	}

	return translatedText, nil
}

func (e *TextParser) translateSegments(segments []string) []string {
	var translatedSegments []string

	// Translate every 5 segments
	for i := 0; i < len(segments); i += 5 {
		end := i + 5
		if end > len(segments) {
			end = len(segments)
		}
		subSegments := segments[i:end]
		subTranslated, err := e.Translator.Translate(strings.Join(subSegments, "\n"), e.From, e.To)
		if err != nil {
			return []string{}
		}
		log.Printf("Translated from %s to %s: %s -> %s\n", e.From, e.To, strings.Join(subSegments, "\n"), subTranslated)
		if e.KeepOrigin {
			translatedSegments = append(translatedSegments, subSegments...)
		}
		translatedSegments = append(translatedSegments, subTranslated)
	}
	return translatedSegments
}
