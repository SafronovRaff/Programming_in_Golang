/*
Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10. Квадрат каждого числа должен выводится в новой строке.

Sample Input:

Sample Output:

1
4
9
16
25
36
49
64
81
100
*/
package main

import "fmt"

func main() {
	var i = 1
	for i <= 10 {
		fmt.Println(i * i)
		i++
	}
}
