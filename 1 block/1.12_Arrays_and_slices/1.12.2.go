package main

import "fmt"

/*Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4),
а затем NN целых чисел-элементов среза.
На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

Sample Input:
5
41 -231 24 49 6
Sample Output:
49
*/

func main() {
	var a int
	fmt.Scan(&a)
	arr := make([]int, a)

	for i := range arr {
		fmt.Scan(&arr[i])

	}
	fmt.Print(arr[3])
}
