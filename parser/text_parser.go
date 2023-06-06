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
	var translatedSegments []string

	// Translate each segment
	for _, segment := range segments {
		translated, err := e.Translator.Translate(segment, e.From, e.To)
		if err != nil {
			return "", err
		}
		log.Printf("Translated from %s to %s: %s -> %s\n", e.From, e.To, segment, translated)
		if e.KeepOrigin {
			translatedSegments = append(translatedSegments, segment)
		}
		translatedSegments = append(translatedSegments, translated)
	}

	// Join the segments back into a single string
	translatedText := strings.Join(translatedSegments, "\n")

	// Write the translated text to a new file
	newName := strings.TrimSuffix(e.Path, filepath.Ext(e.Path)) + ".trans.txt"
	err = os.WriteFile(newName, []byte(translatedText), 0644)
	if err != nil {
		return "", err
	}

	return translatedText, nil
}
