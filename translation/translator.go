package translation

import "strings"

func sanitizeInput(w string) string {
	w = strings.ToLower(w)
	return strings.TrimSpace(w)
}

func Translate(word string, language string) string {

	word = sanitizeInput(word)
	//Sanitizes incoming word
	language = sanitizeInput(language)
	//Sanitizes incoming language

	if word != "hello" {
		return ""
	}
	switch language {
	case "english":
		return "hello"
	case "finnish":
		return "hei"
	case "german":
		return "hallo"
	default:
		return ""
	}
}



