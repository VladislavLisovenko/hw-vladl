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
	}
}

func countWords(text string) map[string]int {
	m := map[string]int{}

	for _, s := range extraSymbols() {
		text = strings.ReplaceAll(text, s, " ")
	}

	words := strings.Split(text, " ")
	for _, word := range words {
		if word == "" {
			continue
		}
		m[word]++
	}

	return m
}
