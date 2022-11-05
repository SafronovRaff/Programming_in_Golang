package main

import (
	"fmt"
	"strconv"
)

func main() {

	res := strconv.FormatBool(10.1 == float32(float64(100/10.1)))

	fmt.Print(res)
}
