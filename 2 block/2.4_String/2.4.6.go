package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		if strings.Count(text, string(i)) == 1 {
			b = append(b, i)
		}

	}
	fmt.Printf(string(b))
}
