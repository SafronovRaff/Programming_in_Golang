package main

/*
Двоичная запись
Дано натуральное число N. Выведите его представление в двоичном виде.
Входные данные
Задано единственное число N
Выходные данные
Необходимо вывести требуемое представление числа N.
Sample Input:
12
Sample Output:
1100
*/
import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Printf("%b", a)
}

/*
package main
import (
	"fmt"
	"strconv"
)
func main() {
	var dec string
	fmt.Scanln(&dec)
	v, _ := ConvertInt(dec, 10, 2)
	fmt.Printf(v) //Десятичное значенияконвертировано двоичное значение:
	               //... по аналогии можно сделать любую конвертацию.
}
// ConvertInt конвертирует значение из одной системы счисления
// в другую, которая указана в toBase
func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}



*/
