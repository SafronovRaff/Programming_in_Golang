package main

import "fmt"

/*
Что будет выведено в результате исполнения этого кода?

5 3

3 5

3 3

5 5
*/
func main() {
	v := 5
	p := &v
	fmt.Print(*p, " ")
	changePointer(p)
	fmt.Print(*p)
}

func changePointer(p *int) {
	v := 3
	p = &v

}
