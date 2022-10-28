package main

import "fmt"

/*

Напишите функцию, которая умножает значения на которые ссылаются два указателя и выводит получившееся произведение в консоль.
Ввод данных уже написан за вас.

Sample Input:
2 2
Sample Output:
4
*/

func main() {
	a := 2
	b := 3
	test(&a, &b)
}

func test(x1 *int, x2 *int) {
	// здесь пишите ваш код
	i := *x1 * *x2

	fmt.Println(i)
}
