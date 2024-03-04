package anagramfinder

import (
	"reflect"
	"sort"
	"testing"
)

// Helper function to sort a slice of slices for consistent comparison
func sortSliceOfSlices(sos [][]string) {
	for _, s := range sos {
		sort.Strings(s)
	}
	sort.Slice(sos, func(i, j int) bool {
		if len(sos[i]) == 0 {
			return true
		}
		if len(sos[j]) == 0 {
			return false
		}
		return sos[i][0] < sos[j][0]
	})
}

func TestAnagramFinder(t *testing.T) {
	af := NewAnagramFinder()
	words := []string{"listen", "silent", "enlist", "inlets", "google", "gooegl"}
	for _, word := range words {
		af.AddWord(word)
	}

	anagrams := af.FindAnagrams()
	expected := [][]string{
		{"listen", "silent", "enlist", "inlets"}, // These words are anagrams of each other
		{"google", "gooegl"},                     // These words are anagrams of each other
	}

	sortSliceOfSlices(anagrams)
	sortSliceOfSlices(expected)

	if !reflect.DeepEqual(anagrams, expected) {
		t.Errorf("Expected anagram groups to be %v, got %v", expected, anagrams)
	}

	anagramsForListen := af.GetAnagramsForWord("listen")
	expectedAnagramsForListen := []string{"enlist", "inlets", "listen", "silent"}

	sort.Strings(anagramsForListen)
	sort.Strings(expectedAnagramsForListen)

	if !reflect.DeepEqual(anagramsForListen, expectedAnagramsForListen) {
		t.Errorf("Expected anagrams for 'listen' to be %v, got %v", expectedAnagramsForListen, anagramsForListen)
	}

	// Testing for a word with no anagrams
	anagramsForWord := af.GetAnagramsForWord("word")
	if len(anagramsForWord) != 0 {
		t.Errorf("Expected no anagrams for 'word', got %v", anagramsForWord)
	}
}
