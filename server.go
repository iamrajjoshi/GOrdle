package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type GAME_STATE struct {
	target_word string
	game_over   bool
}

func make_game_state() *GAME_STATE {
	return &GAME_STATE{
		target_word: target_words[rand.Intn(TARGET_WORDS_SIZE-1)],
		game_over:   false,
	}
}

func load_words() {
	read_file("wordle-La.txt", target_words[:])
	read_file("wordle-Ta.txt", guessable_words[:])
}

func process_guess(state *GAME_STATE, guess string) [WORD_LENGTH]Color {
	ret := [WORD_LENGTH]Color{}
	for index, letter := range guess {
		if letter == rune(state.target_word[index]) {
			ret[index] = GREEN
		} else if strings.ContainsRune(state.target_word, letter) {
			ret[index] = YELLOW
		} else {
			ret[index] = GRAY
		}
	}
	return ret
}

type Color int

const (
	GRAY Color = iota
	YELLOW
	GREEN
)

const TARGET_WORDS_SIZE int = 2315
const GUESSABLE_WORDS_SIZE int = 10657
const WORD_LENGTH int = 5

var target_words [TARGET_WORDS_SIZE]string
var guessable_words [GUESSABLE_WORDS_SIZE]string

func read_file(filename string, data []string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var i = 0
	for scanner.Scan() {
		// do something with a line
		// fmt.Printf("line: %s\n", scanner.Text())
		data[i] = scanner.Text()
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	load_words()

	state := make_game_state()

	var c Color = GRAY
	fmt.Println(c)

	// fmt.Println(target_words[rand.Intn(TARGET_WORDS_SIZE-1)])
	// fmt.Println(guessable_words[rand.Intn(GUESSABLE_WORDS_SIZE-1)])
	fmt.Println(state.target_word)
	t := process_guess(state, "hello")
	fmt.Println(t)

}
