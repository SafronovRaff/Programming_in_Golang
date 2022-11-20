package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
Данная задача ориентирована на реальную работу с данными в формате JSON.
Для решения мы будем использовать справочник ОКВЭД (Общероссийский классификатор видов экономической деятельности),
который был размещен на web-портале data.gov.ru.

Необходимая вам информация о структуре данных содержится в файле structure-20190514T0000.json,
а сами данные, которые вам потребуется декодировать, содержатся в файле data-20190514T0100.json. Файлы размещены в нашем репозитории на github.com.

Для того, чтобы показать, что вы действительно смогли декодировать документ вам необходимо в качестве ответа записать сумму полей global_id всех элементов,
закодированных в файле.
*/
type dataStruct struct {
	GlobalID int `json:"global_id"`
}

func main() {

	var dataStruct []dataStruct

	const URL = "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"

	rs, err := http.Get(URL)
	if err != nil {
		panic("не могу прочитать файл с git ")
	}
	defer rs.Body.Close()

	byteValue, _ := ioutil.ReadAll(rs.Body)
	json.Unmarshal(byteValue, &dataStruct)

	sum := 0.0
	for i := 0; i < len(dataStruct); i++ {
		sum += float64(dataStruct[i].GlobalID)
	}
	fmt.Println(sum) // 7.63804981288e+11

}
