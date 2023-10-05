package primitive_var

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"
)

type Watchdog struct{ last int64 }

func (w *Watchdog) KeepAlive() {
	w.last = time.Now().UnixNano() // First conflicting access.
}

func (w *Watchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			// Second conflicting access.
			if w.last < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			} else {
				fmt.Println("Still Alive!")
			}
		}
	}()
}

type SafeWatchdog struct{ last int64 }

func (w *SafeWatchdog) KeepAlive() {
	atomic.StoreInt64(&w.last, time.Now().UnixNano()) // First conflicting access.
}

func (w *SafeWatchdog) Start() {
	go func() {
		for {
			time.Sleep(time.Second)
			// Second conflicting access.
			if atomic.LoadInt64(&w.last) < time.Now().Add(-10*time.Second).UnixNano() {
				fmt.Println("No keepalives for 10 seconds. Dying.")
				os.Exit(1)
			} else {
				fmt.Println("Still Alive!")
			}
		}
	}()
}
