package main

import (
	"fmt"
)

func main() {
	var num int
	m := map[int]int{}

	for i := 0; i < 10; i++ {
		fmt.Scan(&num)

		if _, value := m[num]; !value {
			m[num] = work1(num)
		}
		num = m[num]
		fmt.Print(num, " ")

	}

}

func work1(x int) int {
	return x - 1
}

// 3 1 5 2 3 5 3 0 3 4
