package main

import "fmt"

/*
Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1.
Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
Напишите функцию, которая по указанному натуральному n возвращает φn.

Входные данные
Вводится одно число n.

Выходные данные
Необходимо вывести  значение φn.

Sample Input:
4
Sample Output:
3
*/
func main() {
	fmt.Println(fibonacci(4))
}
func fibonacci(n int) int {
	//print your code here

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
