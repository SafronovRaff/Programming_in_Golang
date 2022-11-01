package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
Длина пароля должна быть не менее 5 символов,
он должен содержать только арабские цифры и буквы латинского алфавита.
На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
Sample Input:
fdsghdfgjsdDD1
Sample Output:
Ok
*/
func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	if log(text) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}

	version2(text)
}

func log(tex string) bool {
	in := []rune(tex)
	if len(in) < 5 {
		return false
	}
	for _, i := range tex {
		if !unicode.Is(unicode.Latin, i) && !unicode.Is(unicode.Digit, i) {
			return false
		}
	}
	return true
}

func version2(text string) {
	for _, i := range text {
		if !unicode.In(i, unicode.Latin, unicode.ASCII_Hex_Digit) || len(text) < 5 {
			fmt.Println("Wrong password")
			return
		}

	}
	fmt.Println("Ok")

}
