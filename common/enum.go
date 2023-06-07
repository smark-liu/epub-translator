package common

import "strings"

// language enums
const (
	EN = "en"
	ZH = "zh-CN"
)

// parser enum
const (
	EPUB = "epub"
	TEXT = "text"
)

// GetParserByName returns the parser name by the file name.
func GetParserByName(filename string) string {
	var parser string
	switch {
	case strings.HasSuffix(filename, ".epub"):
		parser = EPUB
	case strings.HasSuffix(filename, ".txt"):
		parser = TEXT
	default:
		parser = ""
	}
	return parser

}
