package main

import "fmt"

/*
Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений ее аргументов x, y, z встречается чаще.

# Входные данные
Вводится 3 числа - x, y и z (x, y и z равны 0 или 1).

# Выходные данные
Необходимо вернуть значение функции от x, y и z.

Sample Input:
0 0 1
Sample Output:
0
*/
func main() {
	fmt.Println(vote(1, 0, 0))
	fmt.Println(vote2(1, 0, 0))
}
func vote(x int, y int, z int) int {
	if x == y || x == z {
		return x
	}
	return y
}

func vote2(x int, y int, z int) int {
	//print your code here

	if (x + y + z) >= 2 {
		return 1
	}
	return 0
}
