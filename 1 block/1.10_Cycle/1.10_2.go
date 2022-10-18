/*
Напишите программу, которая последовательно делает следующие операции с введённым числом:

Число умножается на 2;
Затем к числу прибавляется 100.

После этого должен быть вывод получившегося числа на экран.

Sample Input:

1
Sample Output:

102
*/
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	sum := 0
	for i := a; i <= b && b < 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}
