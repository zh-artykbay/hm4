package main

import (
	"sort"
	"strings"
)

type WordCount struct {
	word  string
	count int
}

func TopWords(s string, n int) []string {

	words := strings.Split(s, " ")
	m := make(map[string]int)
	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	wordCounts := make([]WordCount, 0, len(m))
	for key, val := range m {
		wordCounts = append(wordCounts, WordCount{word: key, count: val})
	}


	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].count > wordCounts[j].count
	})

	result := []string{}
	for i := 0; i < len(wordCounts) && i < n; i++ {
		result = append(result, wordCounts[i].word)
	}
	return result
}
