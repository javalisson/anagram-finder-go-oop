package anagramfinder

import (
	"sort"
	"strings"
)

type AnagramFinder struct {
	words []string // Private field
}

func NewAnagramFinder() *AnagramFinder {
	return &AnagramFinder{}
}

func (af *AnagramFinder) AddWord(word string) {
	af.words = append(af.words, word)
}

func (af *AnagramFinder) NormalizeWord(word string) string {
	chars := strings.Split(word, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func (af *AnagramFinder) FindAnagrams() [][]string {
	anagramsMap := make(map[string][]string)
	for _, word := range af.words {
		normalized := af.NormalizeWord(word)
		anagramsMap[normalized] = append(anagramsMap[normalized], word)
	}

	result := make([][]string, 0)
	for _, group := range anagramsMap {
		if len(group) > 1 {
			result = append(result, group)
		}
	}
	return result
}

func (af *AnagramFinder) GetAnagramsForWord(word string) []string {
	normalized := af.NormalizeWord(word)
	for _, group := range af.FindAnagrams() {
		if af.NormalizeWord(group[0]) == normalized {
			return group
		}
	}
	return []string{}
}
