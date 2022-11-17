package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil { // норм ли путь
		return err
	}

	if info.IsDir() { // пропускаем папки
		return nil
	}

	csvFile, _ := os.Open(path) // открываем файл, но не закрываем?

	r := csv.NewReader(csvFile) // какая-то магия

	record, err := r.ReadAll() // record типа буфер на каждой итерации
	if err != nil {
		return err
	}

	if len(record) == 10 && len(record[len(record)-1]) == 10 {
		fmt.Print(record[4][2])
	}

	return nil
}
func main() {
	const root = "C:/Users/R/GolandProjects/Programming_in_Golang/task"
	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}

}
