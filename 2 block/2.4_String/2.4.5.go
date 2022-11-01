package main

import (
	"bufio"
	"fmt"
	"os"
	strings "strings"
)

/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

Sample Input:

zaabcbd
Sample Output:

zcd
*/
func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var b = []rune{}
	for _, i := range text {
		if strings.Count(text, string(i)) == 1 { // Кол-во подстрок в строке strings.Count("test", "t"),  результат: 2
			b = append(b, i)
		}

	}
	fmt.Printf(string(b))
}
