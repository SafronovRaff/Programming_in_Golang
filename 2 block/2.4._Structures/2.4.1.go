package main

import "fmt"

type Circle struct {
	x, y, r float64
}

func main() {
	a := Circle{1, 2, 3}
	fmt.Println("Call linked method", a.testLinked())
	fmt.Println("after", a)
	b := Circle{1, 2, 3}
	fmt.Println("Call method", b.test())
	fmt.Println("after", b)
}

func (c *Circle) testLinked() *Circle {
	c.r = 123
	return c
}

func (c Circle) test() Circle {
	c.r = 123
	return c
}
