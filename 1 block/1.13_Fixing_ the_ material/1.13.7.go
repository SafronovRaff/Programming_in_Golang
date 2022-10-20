package main

import "fmt"

/*
Количество нулей
По данным числам, определите количество чисел, которые равны нулю.
Входные данные
Вводится натуральное число N, а затем N чисел.
Выходные данные
Выведите количество чисел, которые равны нулю.
Sample Input:
5
1
8
100
0
12
Sample Output:
1
*/
func main() {
	var a, b int
	fmt.Scanln(&a)
	n := 0
	for i := 0; i < a; i++ {
		fmt.Scanln(&b)
		if b == 0 {
			n++
		}
	}
	fmt.Println(n)
}
