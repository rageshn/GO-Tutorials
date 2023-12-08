package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	//Open file
	f, err := os.Open("tale-of-two-cities")
	if err != nil {
		fmt.Println("Error while opening file:", err)
		os.Exit(1)
	}
	defer f.Close() // Waits till the program end

	//Gets the word frequency from the file
	words, err := getWordsFrequency(f)
	if err != nil {
		fmt.Println("Error while getting words frequency:", err)
	}

	for word, freq := range words {
		fmt.Printf("Word = %s Frequency = %d \n", word, freq)
	}

}

func getWordsFrequency(r io.Reader) (map[string]int, error) {
	// Map to store the words and their frequency
	wordsFrequency := make(map[string]int)

	// Create a scanner to read the file
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	// REad every word and update the frequency
	for s.Scan() {
		word := strings.ToLower(s.Text())
		wordsFrequency[word]++
	}

	// check for errors
	if err := s.Err(); err != nil {
		return nil, err
	}

	return wordsFrequency, nil
}
