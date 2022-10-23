package main

import "fmt"

/*
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n,
что φn=A. Если А не является числом Фибоначчи, выведите число -1.

Входные данные

Вводится натуральное число.

Выходные данные

Выведите ответ на задачу
*/
/*
// Не проходит тест, время выполнения превышает отведенное платформой время.
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
*/
/*
func fibbonachi(n int) int {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		a, b = b, b+a
	}
	return a
}
func main() {
	var n int
	fmt.Scanln(&n)

	if n == 1 {
		fmt.Println(2)
	} else if n == 2 {
		fmt.Println(3)
	} else if n == 3 {
		fmt.Println(4)
	} else if n == 5 {
		fmt.Println(5)
	} else {
		for i := 0; i < n; i++ {
			if n == fibbonachi(i) {
				fmt.Print(i)
				return
			}
		}
		fmt.Println(-1)
	}

}
*/

func main() {
	var a int
	fmt.Scan(&a)
	f1 := 0
	f2 := 1
	cnt := 0
	for f1 < a {
		f1, f2 = f2, f1+f2
		cnt++
	}
	if f1 == a {
		fmt.Println(cnt)
	} else {
		fmt.Println(-1)
	}
}
