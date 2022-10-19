package main

import "fmt"

/*
Дано трехзначное число. Переверните его, а затем выведите.
Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.
Формат выходных данных
Выведите перевернутое число.
Sample Input:
843
Sample Output:
348
*/
func main() {

	var num, a, b, c int
	fmt.Scan(&num)
	a = num / 100 % 10
	b = num / 10 % 10
	c = num % 10

	fmt.Printf("%d%d%d\n\n\n", c, b, a)
	reverse_digit(num)
	fmt.Print(reverse_digit(num))
}

func reverse_digit(num int) int {
	result := 0
	for num > 0 {
		tmp := num % 10
		result = result*10 + tmp
		num /= 10
	}
	return result
}
