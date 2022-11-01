package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
На вход подается строка. Нужно определить, является ли она правильной или нет.
Правильная строка начинается с заглавной буквы и заканчивается точкой.
Если строка правильная - вывести Right иначе - вывести Wrong
Маленькая подсказка: fmt.Scan() считывает строку до первого пробела, вы можете считать полностью строку за раз с помощью bufio:
text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
Sample Input:
Быть или не быть.
Sample Output:
Right
*/

func main() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	rs := []rune(text)

	text = strings.Trim(text, "\n \r") // // Возвращаем строку с вырезанным пробелом и переносом строки

	if unicode.IsUpper(rs[0]) == true && strings.HasSuffix(text, ".") { // проверка нулевого индекса на верхний регистр и заканчивается ли строка суффиксом "."
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}

	v2(text)

}

func v2(tex string) {

	rs := []rune(tex)

	if unicode.IsUpper(rs[0]) == true && string(rs[len(rs)-1]) == `.` { // на винде -3 строка т.к  строка /r /n (строка + 2 стмвола)
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
