/*
Найдите количество минимальных элементов в последовательности.
Входные данные
Вводится натуральное число N, а затем N целых чисел последовательности.
Выходные данные
Выведите количество минимальных элементов последовательности.
Sample Input:
3
21
11
4
Sample Output:
1

4
21
2
11
2
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, min int
	fmt.Scan(&a, &min)
	num := 1
	for i := 0; i < a-1; i++ {
		fmt.Scan(&b)
		if b < min {
			min = b
			num = 1
		} else if b == min {
			num++
		}
	}
	fmt.Println(num)
	v2()
}
func v2() {
	var inpurNum, counter, min int

	counter = 1
	fmt.Scan(&inpurNum)
	arr := make([]int, inpurNum)
	for i := 0; i < inpurNum; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Ints(arr)
	min = arr[0]
	for i := 1; i < inpurNum; i++ {
		if min == arr[i] {
			counter += 1

		}

	}
	fmt.Println(counter)
}
