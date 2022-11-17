package main

import (
	// "archive/zip"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

/*Поиск файла в заданном формате и его обработка
Данная задача поможет вам разобраться в пакете encoding/csv и path/filepath, хотя для решения может быть использован также пакет archive/zip
(поскольку файл с заданием предоставляется именно в этом формате).

В тестовом архиве, который вы можете скачать из нашего репозитория на github.com, содержится набор папок и файлов.
Один из этих файлов является файлом с данными в формате CSV, прочие же файлы структурированных данных не содержат.

Требуется найти и прочитать этот единственный файл со структурированными данными (это таблица 10х10, разделителем является запятая),
а в качестве ответа необходимо указать число, находящееся на 5 строке и 3 позиции (индексы 4 и 2 соответственно).
https://github.com/semyon-dev/stepik-go/tree/master/work_with_files/task_csv_1

*/func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() { // пропускаем папки
		return nil
	}

	csvFile, _ := os.Open(path)

	r := csv.NewReader(csvFile)

	rec, err := r.ReadAll()
	if err != nil {
		return err
	}

	if len(rec) == 10 {
		fmt.Print(rec[4][2])
	}

	return nil
}
func main() {
	const root = "/home/roman/GolandProjects/Programming_in_Golang/3 block/3.5_Working with files/task"
	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
	// zipFunc()
}

/*func zipFunc() {
	zipFile, _ := zip.OpenReader("task.zip")

	defer zipFile.Close()

	for _, file := range zipFile.File {
		f, _ := file.Open()

		data, _ := csv.NewReader(f).ReadAll()
		if len(data) == 10 {
			fmt.Println(data[4][2])
		}

		f.Close()
	}

}
*/
