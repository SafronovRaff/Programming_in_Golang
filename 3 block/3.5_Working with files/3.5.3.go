package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	if info.IsDir() {
		return nil // Проигнорируем директории
	}

	fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)
	return nil
}

func main() {
	const root = "C:/Users/R/GolandProjects/Programming_in_Golang" // Файлы моей программы находятся в другой директории

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}

}
