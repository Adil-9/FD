package dictionary

import (
	"bufio"
	"dimploma/variables"
	"log"
	"os"
	"strings"
	"sync"
)

var count int = 0
var printChannel = make(chan string, 1024)
var done = make(chan struct{}) // Channel to signal completion
var SpecialCharAll = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

var wg sync.WaitGroup

func CreateDict(AllVar variables.AllVariables) {
	if AllVar.SpecialChar && len(AllVar.StringVar) > 4 {
		log.Println("Using 5 strings and including special charachters will result in 1 141 188 059 880 lines of possible passwords which is too many for the computer!!! Denied")
		return
	}
	// Generate and write combinations to the file

	log.Println(AllVar.StringVar)
	AllVar.FormattedStrings = formatAndCreateWordSamples(AllVar.StringVar)
	go printFunc(AllVar.Filename)
	generateCombinations(AllVar.FormattedStrings, 0, AllVar.NumberOfWords, AllVar.SpecialChar)

	wg.Wait()
	close(printChannel) // Close the printChannel to signal completion
	<-done              // Wait for the printFunc goroutine to finish
	log.Println(count, "Permutaions of passwords have been written to Passwords.txt")
}

// permute generates all permutations of the given slice of strings
func generateCombinations(words [][]string, start int, numOfWords int, includeSpechialChars bool) {
	if start == len(words)-1 {
		if includeSpechialChars {
			permuteAllPairsWithSpecialChars(words, 0, []string{}, numOfWords)
		} else {
			permuteAllPairs(words, 0, []string{}, numOfWords)
		}
		return
	}

	for i := start; i < len(words); i++ {
		// Swap words[start] and words[i]
		words[start], words[i] = words[i], words[start]
		// Recursively generate permutations for the rest of the words
		generateCombinations(words, start+1, numOfWords, includeSpechialChars)
		// Restore the original order for backtracking
		words[start], words[i] = words[i], words[start]
	}
}

// permuteAllPairs generates all possible pairings of words from different arrays
func permuteAllPairs(arr [][]string, index int, current []string, numOfWords int) {
	if numOfWords > len(arr) || numOfWords <= 0 {
		if index == len(arr) { //here can change len(arr) to int of words we need
			wg.Add(1)
			printChannel <- strings.Join(current, "")
			count++
			return
		}
	} else {
		if index == numOfWords { //here can change len(arr) to int of words we need
			wg.Add(1)
			printChannel <- strings.Join(current, "")
			count++
			return
		}
	}

	for _, word := range arr[index] {
		// Generate permutations by appending the word to the current combination
		permuteAllPairs(arr, index+1, append(current, word), numOfWords)
	}
}

func permuteAllPairsWithSpecialChars(arr [][]string, index int, current []string, numOfWords int) {
	if numOfWords > len(arr) || numOfWords <= 0 {
		if index == len(arr) { //here can change len(arr) to int of words we need
			wg.Add(1)
			printChannel <- strings.Join(current, "")
			count++
			return
		}
	} else {
		if index == numOfWords { //here can change len(arr) to int of words we need
			wg.Add(1)
			printChannel <- strings.Join(current, "")
			count++
			return
		}
	}

	for _, word := range arr[index] {
		// Generate permutations by appending the word to the current combination
		for _, v := range SpecialCharAll {
			permuteAllPairsWithSpecialChars(arr, index+1, append(current, word+string(v)), numOfWords)
		}
		permuteAllPairsWithSpecialChars(arr, index+1, append(current, word), numOfWords)
	}
}

func formatAndCreateWordSamples(strs []string) [][]string {
	temp := [][]string{}
	for i := range strs {
		temp = append(temp, []string{strings.ToUpper(strs[i]), strings.ToLower(strs[i]), strings.Title(strs[i])})
	}

	return temp
}

func printFunc(fileName string) {
	defer close(done) // Signal completion of printFunc when it's done

	if len(fileName) < 4 {
		fileName += ".txt"
	} else if fileName[len(fileName)-4:] != ".txt" {
		fileName += ".txt"
	}

	if _, err := os.Stat(fileName); err == nil {
		// Delete the file if it already exists
		err := os.Remove(fileName)
		if err != nil {
			log.Fatal("Error deleting file:", err)
			return
		}
	}

	// Create the file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for msg := range printChannel {
		writer.Write([]byte(msg + "\n"))
		wg.Done()
	}
	writer.Flush()
}
