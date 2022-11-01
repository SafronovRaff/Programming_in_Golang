package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

Sample Input:

ihgewlqlkot
Sample Output:

hello

*/

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	t := []byte(text)
	var a = []byte{}
	for i, _ := range t {
		if i%2 != 0 {
			a = append(a, t[i])
		}
	}
	fmt.Printf("%s", a)

}
