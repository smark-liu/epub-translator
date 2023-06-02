package parser

import (
	"github.com/smark-d/epub-translator/trans"
)

type Parser interface {
	// Parse returns the translated file Path.
	Parse() (string, error)
}

func GetParser(parserType, path, sourceLang, targetLong, translator string, keepOrigin bool) Parser {
	switch parserType {
	case "epub":
		return &EpubParser{
			Path:       path,
			From:       sourceLang,
			To:         targetLong,
			KeepOrigin: keepOrigin,
			Translator: trans.GetTranslator(translator),
		}
	default:
		return nil
	}
}
