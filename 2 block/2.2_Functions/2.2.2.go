package main

import "fmt"

/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные
Вводится четыре числа.

Выходные данные
Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:
4 5 6 7

Sample Output:
4
*/
func main() {

	fmt.Println(minimumFromFour())
	//fmt.Println(minimumFromFourV2())
}

func minimumFromFour() int {
	var num int
	_, _ = fmt.Scan(&num)

	for i := 0; i < 3; i++ {
		var min int
		_, _ = fmt.Scan(&min)

		if min < num {
			num = min
		}
	}
	return num
	/*	//print your code here
		var a, b, c, d int
		fmt.Scan(&a, &b, &c, &d)

		m := 0
		if a >= b {
			m = b
		} else {
			m = a
		}
		if c >= d {
			if m <= d {
				return (m)
			} else {
				return (d)
			}
		} else if m <= d {
			return (m)
		} else {
			return (c)
		}
	*/

}

func minimumFromFourV2() int {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	return min(min(a, b), min(c, d))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
