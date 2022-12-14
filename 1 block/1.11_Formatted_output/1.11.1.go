/*
На вход подается число типа float64. Вам нужно вывести преобразованное число по правилу:
число возводится в квадрат, затем обрезается дробная часть так что остается 4 знака после запятой.
Но если число равно или меньше нуля - выводить:
"число R не подходит", где R - исходное число с 2 цифрами после запятой и с 2 по ширине.
А если число больше чем 10 000 - выводить исходное число в экспоненциальном представлении (знак экспоненты в нижнем регистре).

Sample input:
-000012.2123
Sample output:
число -12.21 не подходит

Sample input:
1000001
Sample output:
1.000001e+06

Sample Input:

12.12345678
Sample Output:

146.9782

1) Провести проверку вводимого числа:

1.1 оно меньше или равно нулю: вывод печати по условию

1.2 оно больше 10000: вывод печати по условию

2) Если предидущих два условия False, то третье условие else {возводим в квадтра, потом печать как задана по условию задачи}
*/
package main

import "fmt"

func main() {

	var input float64
	fmt.Scanln(&input)
	switch {
	case input <= 0:
		fmt.Printf("число %2.2f не подходит", input)
	case input > 10000:
		fmt.Printf("%e", input)
	default:
		fmt.Printf("%.4f", input*input)

	}
}
