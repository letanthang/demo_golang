package main

func main() {
	ch := make(chan int)

	// read channel cause current routine sleep
	<-ch
}
