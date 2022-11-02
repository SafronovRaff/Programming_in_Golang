package main

import (
	"fmt"
	"math"
)

/*
На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы

Sample Input:

6 8
Sample Output:

10
*/
func main() {
	var a, b float64
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("ошибка ввода")
		return
	}
	fmt.Println(math.Hypot(a, b))

}
