package trans

type Translator interface {
	// Translate returns the translation of the given string.
	// source, sourceLang, targetLang
	Translate(string, string, string) (string, error)
}

func GetTranslator(translatorType string) Translator {
	switch translatorType {
	case "google":
		return &GoogleTranslator{}
	default:
		return nil
	}
}
