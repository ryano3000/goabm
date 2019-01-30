package main

import (
	"fmt"

	"github.com/Vasanthakumar.V/goab/internal/shopper"
	"github.com/Vasanthakumar.V/goab/pkgs/abm"
)

func main() {
	// Create a world with arg which specify
	// the buffer size for the channel and the max number of 
	// goroutines
	w := abm.NewWorld(1, 2)

	// Create slice of agents to send in
	ss := []abm.Agent{&shopper.Agent{}, &shopper.Agent{}}

	// Run the simulation for 5 ticks and 
	in := w.Start(5, ss)

	// Read in and print the changed status of agents
	// from the channel returned after start
	for s := range in {
		ss := s.(*shopper.Agent)
		fmt.Printf("%#v\n", *ss)
	}
}
