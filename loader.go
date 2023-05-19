package main

import (
	"bufio"
	"log"
	"os"
)

// // Scope
// // client only handles itself and guesses it made
// // dictionary handled by server
// // --> validate guess

// // decide protocol
// // decide how to handle dictionary
// // dictionary class

// //https://gist.github.com/scholtes/94f3c0303ba6a7768b47583aff36654d

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
