package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	buf, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	{
		fmt.Println("не могу считать данные")
	}
	buf = strings.TrimSpace(buf)

}
