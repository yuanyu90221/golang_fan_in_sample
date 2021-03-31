package main

import (
	"fmt"

	fan "github.com/yuanyu90221/golang_fan_in_sample/fan"
)

func main() {
	in := fan.Producer(1, 2, 3, 4)
	// FAN-OUT
	c1 := fan.Square(in)
	c2 := fan.Square(in)
	c3 := fan.Square(in)
	// consumer
	for ret := range fan.Merge(c1, c2, c3) {
		fmt.Printf("%3d", ret)
	}
	fmt.Println()
}
