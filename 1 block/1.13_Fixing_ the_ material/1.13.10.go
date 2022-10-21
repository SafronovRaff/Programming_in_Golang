package main

/*
Самое большое число, кратное 7
Найдите самое большее число на отрезке от a до b, кратное 7 .

Входные данные
Вводится два целых числа a и b (a≤b).

Выходные данные
Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b),
кратное 7 , или выведите "NO" - если таковых нет.

Sample Input:
100
500
Sample Output:

497
*/
import "fmt"

func main() {
	var a, b, highest int
	fmt.Scan(&a)
	fmt.Scan(&b)
	for i := b; i >= a; i-- {
		if i%7 == 0 {
			highest = i
			fmt.Println(highest)
			return
		}

	}
	fmt.Println("NO")
}

/*
func main() {

	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	for ; a <= b; b-- {
		if (b % 7) == 0 {
			fmt.Println(b)
			break
		} else if a == b {
			fmt.Println("NO")
		}
	}

}
*/
