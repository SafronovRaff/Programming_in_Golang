package main

import (
	"strconv"
)

func main() {
	del(727178)

}

func del(num uint) uint {

	numSTR := strconv.FormatUint(uint64(num), 10)

	arr := make([]rune, 0, len(numSTR))

	for _, i := range numSTR {
		if i%2 == 0 && i != 48 {
			arr = append(arr, i)
		}
	}

	numSTR = string(arr)

	conclusion, _ := strconv.ParseUint(numSTR, 10, 64)

	if conclusion == 0 {
		conclusion = 100
	}

	return uint(conclusion)
}
