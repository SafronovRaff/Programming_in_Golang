package main

import "fmt"

/*
На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.

Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181

Sample Input:

9119
Sample Output:

811181
*/
func main() {
	var text string
	_, err := fmt.Scanln(&text)
	if err != nil {
		fmt.Println("ошибка")
	}
	t := []rune(text)
	for _, i := range t {
		fmt.Print((i - 48) * (i - 48)) //Цифры начинаются с 48 символа в ASCII https://www.rapidtables.com/code/text/ascii-table.html
	}

}
