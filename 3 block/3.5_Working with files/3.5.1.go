package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Какие файлы останутся после выполнения данной программы?

	func main() {
		for i := 1; i <= 3; i++ {
	file, err := os.Create(strconv.Itoa(i) + "text.txt")
			if err != nil {
				panic(err)
			}
			defer file.Close()
		}
		os.Rename("2text.txt", "4text.txt")
	    for i := 4; i >= 2; i-- {
	os.Remove(strconv.Itoa(i) + "text.txt")
		}
	}
*/
func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println("Создаем файл", (strconv.Itoa(i) + "text.txt"))
		file, err := os.Create(strconv.Itoa(i) + "text.txt")
		if err != nil {
			panic(err)
		}
		//		defer file.Close()
		defer checkDefer(file)
	}

	fmt.Println("переименовываем файлы")
	err := os.Rename("2tex0t.txt", "4text.txt")
	if err != nil {
		fmt.Println("ОШИБКА!", err)
	}

	fmt.Println("Удаляем файлы")
	for i := 4; i >= 2; i-- {
		fmt.Println("Удаляем файл с префиксм", i)
		err := os.Remove(strconv.Itoa(i) + "text.txt")
		if err != nil {
			fmt.Println("ОШИБКА!", err)
		}
	}
}

// нужно чтобы протестировать когда же выполнится наш defer использовать функцию с выводом на консоль
func checkDefer(f *os.File) error {
	fmt.Printf("it's a DEFER TIME!!! (with file name %v)\n", f.Name())
	return f.Close()
}
