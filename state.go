package main

import (
	"math/rand"
)

type game_state struct {
	target_word string
	game_over   bool
}

func make_game_state() *game_state {
	return &game_state{
		target_word: target_words[rand.Intn(TARGET_WORDS_SIZE-1)],
		game_over:   false,
	}
}
