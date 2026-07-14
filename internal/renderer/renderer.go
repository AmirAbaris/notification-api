package renderer

import (
	"fmt"
	"strings"
)

func Render(template string, data map[string]string) string {
	extractedKeys := extractKeys(template)
	result := template

	for _, word := range extractedKeys {
		if _, ok := data[word]; ok {
			fmt.Println(result)
			result = strings.ReplaceAll(result, "{{"+word+"}}", data[word])
		}
	}

	return result
}

func extractKeys(text string) []string {
	var result []string
	words := strings.SplitSeq(text, " ")
	for word := range words {
		if word[0] == '{' && word[len(word)-1] == '}' {
			result = append(result, word[2:len(word)-2])
		}
	}

	return result
}
