/*
По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:

237
Sample Output 1:

YES
Sample Input 2:

117
Sample Output 2:

NO
*/
package main

import "fmt"

func main() {
	var d, a, b, c int
	fmt.Scanln(&d)
	a = d / 100
	b = d / 10 % 10
	c = d % 10
	if a == b {
		fmt.Println("NO")
	} else if a == c {
		fmt.Println("NO")
	} else if b == c {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
