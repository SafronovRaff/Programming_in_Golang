package main

import "fmt"

/*
Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0.
Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).

Входные данные
Сначала задано число NN — количество элементов в массиве (1 \leq N \leq 1001≤N≤100).
Далее через пробел записаны NN чисел — элементы массива. Массив состоит из целых чисел.

Выходные данные
Необходимо вывести все элементы массива с чётными индексами.

Sample Input:
6
4 5 3 4 2 3
Sample Output:
4 3 2
*/

func main() {
	var a int
	fmt.Scan(&a)
	arr := make([]int, a)
	for i := range arr {
		fmt.Scan(&arr[i])
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	// 6
	// 4 5 3 4 2 3
}

/*
var a int
	fmt.Scan(&a)
	arr := make([]int, a)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	for in, i := range arr {
		if in%2 == 0 {
			fmt.Print(i, " ")
		}
	}
*/
