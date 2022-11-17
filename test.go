package main

import "fmt"

func main() {
	var num, sample int
	fmt.Scan(&num, &sample)
	var arr []int
	counter := 0

	for ; num > 0; num /= 10 {
		if num%10 != sample {
			arr = append(arr, num%10)
			counter++
		}
	}

	for i := counter - 1; i >= 0; i-- {
		fmt.Print(arr[i])
	}

}
