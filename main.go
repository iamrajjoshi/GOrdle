package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var GAME_STATE *game_state

func run() error {
	if len(os.Args) < 2 {
		return errors.New("please provide an address to listen on as the first argument")
	}

	l, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		return err
	}
	log.Printf("listening on http://%v", l.Addr())

	cs := newChatServer()
	s := &http.Server{
		Handler:      cs,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	errc := make(chan error, 1)
	go func() {
		errc <- s.Serve(l)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	select {
	case err := <-errc:
		log.Printf("failed to serve: %v", err)
	case sig := <-sigs:
		log.Printf("terminating: %v", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return s.Shutdown(ctx)
}

func main() {

	load_words()

	log.SetFlags(0)

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	GAME_STATE = make_game_state()

	var c Color = GRAY
	fmt.Println(c)

	// fmt.Println(target_words[rand.Intn(TARGET_WORDS_SIZE-1)])
	// fmt.Println(guessable_words[rand.Intn(GUESSABLE_WORDS_SIZE-1)])
	fmt.Println(GAME_STATE.target_word)
	t := process_guess(GAME_STATE, "hello")
	fmt.Println(t)

}
