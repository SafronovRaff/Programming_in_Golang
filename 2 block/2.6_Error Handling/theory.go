package main

import "fmt"

func main() {
	fmt.Println(div(15, 5))
	defer fmt.Println("Program has been finished")
	/*
		Если defer находится выше panic (функции, которая вызывает панику), то сначала
		выполнится дефер, а потом паника. Однако, если defer находится ниже panic,
		то defer не выполнится вовсе!
	*/
	fmt.Println(div(4, 0)) // момент паники

}
func div(x, y float64) float64 {
	if y == 0 {
		panic("division by zero!")
	}
	return x / y
}

// 3
// Program has been finished
// panic: division by zero!
