package main

import "fmt"

type State int

const (
	Running State = iota
	Stopped
	Rebooting
	Terminated
)

func (s State) String() string {
	return []string{"running", "stopped", "rebooting", "terminated"}[s]
}
func main() {
	fmt.Println(Running, Stopped, Rebooting, Terminated)
}
