package main

import "fmt"

type Ball struct {
	Radius int
}

func (b Ball) Bounce() {
	fmt.Println(b, b.Radius)
}

type BasketBall struct {
	Radius int
	*Ball
	Weight int
}

type Action interface {
	Bounce()
}

type VolleyBall struct {
	Action
	Weight int
}

func main() {
	// b := Ball{40}
	// bb := BasketBall{Weight: 100, Ball: &b, Radius: 60}
	// bb.Bounce()
	// fmt.Println(bb.Ball.Radius)

	// embed interface
	vb := VolleyBall{&Ball{50}, 100}

	vb.Bounce()
	// vb.Radius
}
