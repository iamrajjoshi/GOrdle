package main

const TARGET_WORDS_SIZE int = 2315
const GUESSABLE_WORDS_SIZE int = 10657
const WORD_LENGTH int = 5

var target_words [TARGET_WORDS_SIZE]string
var guessable_words [GUESSABLE_WORDS_SIZE]string

// Defined in wordle.pb.go
// type Color int

const (
	GRAY Color = iota
	YELLOW
	GREEN
)
