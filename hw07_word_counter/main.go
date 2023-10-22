package hw07wordcounter

import "strings"

func extraSymbols() []string {
	return []string{
		",",
		".",
		"!",
		"?",
		":",
		";",
		"-",
		"\t",
		"\n",
		"\r",
	}
}

func countWords(text string) map[string]int {
	m := map[string]int{}
	text = strings.ToLower(text)

	for _, s := range extraSymbols() {
		text = strings.ReplaceAll(text, s, " ")
	}

	words := strings.Fields(text)
	for _, word := range words {
		m[word]++
	}

	return m
}
