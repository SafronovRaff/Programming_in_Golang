package main

import (
	"fmt"
	"strings"
)

type bat uint8

func (f *bat) newBat(value string) {
	*f = bat(strings.Count(value, `1`))
}

func (f bat) Strings() string {
	return fmt.Sprintf("[%10s]", strings.Repeat("X", int(f)))

}

func main() {
	var ForTest bat
	var value string
	fmt.Scan(&value)
	ForTest.newBat(value)
	fmt.Println(ForTest)

}
