package main

import "fmt"

/*
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.

Входные данные

Вводится натуральное число.

Выходные данные

Выведите ответ на задачу
*/

func fibbonachi(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}
func main() {
	var n, a int
	fmt.Scanln(&a)
	if a == 1 {
		fmt.Println(2)
	} else if a == 0 {
		fmt.Println(1)
	} else {
		arr := make([]int, 1)
		for n = 1; n <= 35; n++ {
			arr = append(arr, fibbonachi(n))
		}

		for key, value := range arr {
			if a == value {
				fmt.Println(key)
				return
			}
		}
		fmt.Print(-1)
	}
}
