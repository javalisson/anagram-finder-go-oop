package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sandbox/anagramfinder" // Adjust the import path based on your module setup
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
    fmt.Println("Endpoint Hit: helloWorldHandler")
}

var anagramFinder = anagramfinder.NewAnagramFinder()

func addWordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Word string `json:"word"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	anagramFinder.AddWord(data.Word)
	w.WriteHeader(http.StatusCreated)
}

func getAllAnagramsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	allAnagrams := anagramFinder.FindAnagrams()
	json.NewEncoder(w).Encode(allAnagrams)
}

func getAnagramsForWordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	word := r.URL.Query().Get("word")
	if word == "" {
		http.Error(w, "Word parameter is required", http.StatusBadRequest)
		return
	}

	anagrams := anagramFinder.GetAnagramsForWord(word)
	json.NewEncoder(w).Encode(anagrams)
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/addword", addWordHandler)
	http.HandleFunc("/allanagrams", getAllAnagramsHandler)
	http.HandleFunc("/anagrams", getAnagramsForWordHandler)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
