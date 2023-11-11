package fix

import (
	"math/rand"
	"time"
)

func search() []string {
	done := make(chan bool)
	result := make(chan []string)
	for i := 0; i < 3; i++ {
		go func() {
			select {
			case result <- get():
			case <-done:
			}
		}()
	}
	r := <-result
	close(done)
	return r
}

func get() []string {
	time.Sleep(time.Duration(rand.Intn(10) * int(time.Millisecond)))
	return []string{"hello", "how"}
}
