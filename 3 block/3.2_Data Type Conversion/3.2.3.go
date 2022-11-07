package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Для решения данной задачи вам понадобится пакет strconv, возможно использовать пакеты strings или encoding/csv, или даже bufio - вы не ограничены в выборе способа решения задачи.
В решениях мы поделимся своими способами решения этой задачи, предлагаем вам сделать то же самое.
В привычных нам редакторах электронных таблиц присутствует удобное представление числа с разделителем разрядов в виде пробела,
кроме того в России целая часть от дробной отделяется запятой. Набор таких чисел был экспортирован в формат CSV, где в качестве разделителя используется символ ";".
На стандартный ввод вы получаете 2 таких вещественных числа, в качестве результата требуется вывести частное от деления первого числа на второе с точностью до четырех знаков после "запятой"
(на самом деле после точки,результат не требуется приводить к исходному формату).
P.S. небольшое отступление, связанное с чтением из стандартного ввода. Кто-то захочет использовать для этого пакет bufio.Reader.
Это вполне допустимый вариант, но если вы резонно обрабатываете ошибку метода ReadString('\n'), то получаете ошибку EOF, а точнее (io.EOF - End Of File).
На самом деле это не ошибка, а состояние, означающее, что файл (а os.Stdin является файлом) прочитан до конца. Чтобы ошибка была обработана правильно, вы можете поступить так:

	if err != nil && err != io.EOF {
		...

Sample Input:
1 149,6088607594936;1 179,0666666666666
Sample Output:
0.9750
*/
func main() {
	strData, _ := bufio.NewReader(os.Stdin).ReadString('\r')
	strData = strings.Trim(strings.Replace(strings.Replace(strData, " ", "", -1), ",", ".", -1), "\n \r")
	/*strData = strings.Trim(strData, "\n \r")
	strData = strings.Replace(strData, ",", ".", -1) // замена "," на "."
	strData = strings.Replace(strData, " ", "", -1)  // убираем пробелы
	*/

	var arrData = strings.Split(strData, ";") // разделяем строку

	// V1 без функций
	/*
		num1, err := strconv.ParseFloat(arrData[0], 64)
		if err != nil {
			panic(err)
		}
		num2, err := strconv.ParseFloat(arrData[1], 64)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%.4f", num1/num2)
	*/

	// V2 через вызов функции
	num1, num2 := arrData[0], arrData[1]
	calculation(num1, num2)

}

func transformation(text string) float64 {
	num, err := strconv.ParseFloat(text, 64) //
	if err != nil {
		panic(err)
	}
	//fmt.Println(num)
	//fmt.Printf("%T \n", num)
	return num

}
func calculation(num1, num2 string) {
	res := transformation(num1) / transformation(num2)

	//fmt.Printf("%T \n", res)
	fmt.Printf("%.4f", res)
}
