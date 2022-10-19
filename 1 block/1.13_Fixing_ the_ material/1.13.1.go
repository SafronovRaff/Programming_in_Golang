package main

import "fmt"

/*
Дано трехзначное число. Найдите сумму его цифр.
Формат входных данных
На вход дается трехзначное число.
Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.
Sample Input:
745
Sample Output:
16
*/
func main() {
	var num, a, b, c, sum int
	fmt.Scan(&num)
	a = num / 100 % 10
	b = num / 10 % 10
	c = num % 10
	sum = a + b + c
	fmt.Print(sum)
}
