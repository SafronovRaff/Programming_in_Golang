package main

import "fmt"

/*
По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

Входные данные
Вводится натуральное число.

Выходные данные
Выведите ответ на задачу.

Sample Input:
50
Sample Output:
1 2 4 8 16 32
*/

func main() {
	var num int

	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("ошибка")
	}
	for i := 1; i <= num; i *= 2 {
		fmt.Print(i, " ")
	}

	bit_shift(uint(num)) // не проходит 6 тест, надо разбирать..

}

func bit_shift(num uint) {

	for i := uint(1); i < num; i <<= 1 {
		fmt.Printf("%d ", i)
	}
}
