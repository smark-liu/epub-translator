package parser

import (
	"github.com/smark-d/epub-translator/common"
	"github.com/smark-d/epub-translator/trans"
)

type Parser interface {
	// Parse returns the translated file Path.
	Parse() (string, error)
}

func GetParser(parserType, path, from, to string) Parser {
	switch parserType {
	case "epub":
		return &EpubParser{
			Path:       path,
			From:       common.EN,
			To:         common.ZH,
			Translator: &trans.GoogleTranslator{},
		}
	default:
		return nil
	}
}
