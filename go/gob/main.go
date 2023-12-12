package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// https://go.dev/blog/gob

type P struct {
	X, Y, Z int
	Name    string
}

func main() {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}
	// Decode (receive) the value.
	var p2 P
	err = dec.Decode(&p2)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: {%d,%d,%d}\n", p2.Name, p2.X, p2.Y, p2.Z)
}
