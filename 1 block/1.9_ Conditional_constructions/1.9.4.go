/*
Определите является ли билет счастливым. Счастливым считается билет,
в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.

Формат входных данных
На вход подается номер билета - одно шестизначное  число.

Формат выходных данных
Выведите "YES", если билет счастливый, в противном случае - "NO".

Sample Input:
613244
Sample Output:
YES

package main
import "fmt"

	func main() {
		var d, a, b, c, e, f, g int
		fmt.Scanln(&d)
		a = d / 100000
		b = d / 10000 % 10
		c = d / 1000 % 10
		e = d / 100 % 10
		f = d / 10 % 10
		g = d % 10
		if a+b+c == e+f+g {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
*/
package main

import "fmt"

func main() {
	var d string
	if fmt.Scan(&d); d[0]+d[1]+d[2] == d[3]+d[4]+d[5] {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
