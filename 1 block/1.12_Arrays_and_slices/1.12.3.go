/*
На ввод подаются пять целых чисел, которые записываются в массив.
Однако эта часть программы уже написана.
Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.

package main
import "fmt"

	func main()  {
		array := [5]int{}
		var a int
		for i:=0; i < 5; i++{
			fmt.Scan(&a)
			array[i] = a
		}
	    // здесь ваш код
	    // ...
	}

Sample Input:
2
3
56
45
21
Sample Output:
56
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		_, err := fmt.Scan(&a)
		if err != nil {
			fmt.Println("ошибка")
		}
		array[i] = a
	}
	// здесь ваш код
	slice := array[:] // перемудрил, надо циклом... т.к. изменил первоначальные условия, добавил библиотеку.
	sort.Ints(slice)
	fmt.Println(slice[len(slice)-1])

	vers2(array)
}

func vers2(array [5]int) {

	max := array[0]
	for _, el := range array {
		if el > max {
			max = el
		}

	}
	fmt.Printf("максимальное число %v", max)
}
