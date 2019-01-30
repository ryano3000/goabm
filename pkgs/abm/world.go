package abm

import "sync"

// World struct has buffer to store the channel buffer
// and guard to limit no of goroutines
type World struct {
	buffer int
	guard  int
}

// NewWorld instantiates the world with buffer for the message channel
func NewWorld(size, guard int) *World {
	return &World{buffer: size, guard: guard}
}

// Start method runs the simulation  one tick at a time
func (w *World) Start(ticks int, aa []Agent) chan interface{} {
	out := make(chan interface{}, w.buffer)
	guard := make(chan struct{}, w.guard)

	// wg to help identify when all agents complete their
	// action after a tick
	var wg sync.WaitGroup

	go func() {
		// outer loop for tick
		for i := 0; i < ticks; i++ {
			// inner loop to range over given slice of agents
			for _, a := range aa {
				wg.Add(1)
				// guard channels buffer size limits no of goroutines
				// started
				guard <- struct{}{}
				go func(a Agent) {
					// output of Run fed to the client through a channel
					out <- a.Run()
					wg.Done()
					<-guard
				}(a)
			}
			wg.Wait()
		}
		// close channel after completion of all ticks
		close(out)
	}()

	// out channel utilized by client to read message from agent runs
	return out
}
