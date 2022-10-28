package main

import "fmt"

/*

Поменяйте местами значения переменных на которые ссылаются указатели. После этого переменные нужно вывести.
ВАЖНО: Считайте что пакет main уже объявлен, а также функция main() уже есть.

Sample Input:

2 4
Sample Output:

4 2
*/

func main() {
	a := 2
	b := 4
	tes(&a, &b)
}
func tes(x1 *int, x2 *int) {
	// здесь ваш код
	*x1, *x2 = *x2, *x1

	fmt.Println(*x1, *x2)
}
