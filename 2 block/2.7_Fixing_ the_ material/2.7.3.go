package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

/*
Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.

# Входные данные

Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.

# Выходные данные

Выведите максимальную цифру, которая встречается во введенной строке.

Sample Input:

1112221112
Sample Output:

2
*/
func main() {
	var a string
	fmt.Scanln(&a)
	b := []byte(a)

	for _, i := range a {
		if utf8.RuneCountInString(a) > 1000 || !unicode.Is(unicode.Digit, i) {
			fmt.Println("в строке более 1000 символов или НЕ только арабские цифры")
			return
		}

	}
	var max byte
	for _, i := range b {
		if i > max {
			max = i
		}
	}
	fmt.Print(string(max))

	//arr := make([]string, a)
	//slice := b[:]
	//sort.Ints(slice)
	//fmt.Println(slice[len(slice)-1])

}
