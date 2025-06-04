package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

type Phrase struct {
	English  string
	Japanese string
}

var phrases []Phrase

func loadPhrases(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records {
		if len(record) < 2 {
			continue // 行のフォーマットが不正ならスキップ
		}
		p := Phrase{
			English:  record[0],
			Japanese: record[1],
		}
		phrases = append(phrases, p)
	}
	return nil
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	if len(phrases) == 0 {
		http.Error(w, "No phrases available", http.StatusInternalServerError)
		return
	}
	p := phrases[rnd.Intn(len(phrases))]
	fmt.Fprintf(w, "English: %s\n日本語: %s\n", p.English, p.Japanese)
}

func handler(w http.ResponseWriter, r *http.Request) {
	date := time.Date(2025, 6, 4, 9, 30, 0, 0, time.Local).Format(time.DateTime)
	fmt.Fprintln(w, "Welcome back. You're doing great studying today, too! 👏  This wep app updated at: "+date)
}

func main() {
	err := loadPhrases("phrases.csv")
	if err != nil {
		fmt.Printf("Failed to load phrases: %v\n", err)
		os.Exit(1)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/random", randomHandler)

	port := "8080"
	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, nil)
}
