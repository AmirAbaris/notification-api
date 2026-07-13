package renderer

import (
	"errors"
	"strings"
)

func Render(template string, data map[string]string) (string, error) {
	var result string
	// var target string
	var end int
	var start int

	for i, ch := range template {
		if ch == '{' {
			start = i + 2
			break
		}
	}

	for j, ch := range template {
		if ch == '}' {
			end = j - 1
			break
		}
	}

	for i := start; i <= end; i++ {
		result += string(template[i])
	}

	if _, ok := data[result]; ok {
		result = strings.ReplaceAll(template, "{{"+result+"}}", data[result])
		return result, nil
	}

	return "", errors.New("not compatable data with template")
}
