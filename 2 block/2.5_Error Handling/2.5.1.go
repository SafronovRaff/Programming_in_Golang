package main

import (
	"fmt"
)

/*
Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат деления, но и информацию об ошибке.
В случае какой-либо ошибки вы должны вывести "ошибка", если ошибки нет - результат функции. Функция divide(a int, b int) (int, error)
получает на вход два числа которые нужно поделить и возвращает результат (int) и ошибку (error). Пакет main уже объявлен, а нужные пакеты уже импортированы!

Не забудьте считать два числа которые необходимо поделить (передать в функцию)!

Sample Input:

10 5
Sample Output:

2
*/

func main() {
	a, b := 0, 0
	_, err := fmt.Scan(&a, &b)

	if err != nil {
		fmt.Println("ошибка")
		return
	}
	c, err := divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
		return
	}
	fmt.Println(c)
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		panic("division by zero!")
	}
	return a / b, nil

}
