package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://service.1dogma.ru/api/layouts-filter/layouts")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response []Response
	json.Unmarshal(body, &response)

	if err != nil {
		log.Fatal(err)
	}

	var result [][]string

	for i := 0; i < len(response); i++ {
		if response[i].Type == 1 {
			var m []string
			m = append(m, fmt.Sprintf("%d", response[i].ComplexId), fmt.Sprintf("%d", response[i].LetterId),
				fmt.Sprintf("%d", response[i].Door), fmt.Sprintf("%d", response[i].SumPrice),
				fmt.Sprintf("%f", response[i].Area), response[i].LayoutsUrl)
			result = append(result, m)
		}
	}

	file, err := os.Create("layouts.csv")

	if err != nil {
		log.Println("Cannot create CSV file:", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	writer.WriteAll(result)
}

type Response struct {
	Area       float64 `json:"area"`        // общая площадь
	Door       int     `json:"door"`        // подъезд
	Type       int     `json:"type"`        // тип помещения, 1-квартиры, 2-коммерция, 3-подсобные помещения, 4-парковки
	LayoutsUrl string  `json:"layouts_url"` //ссылка на изображение
	SumPrice   int     `json:"sum_price"`   // цена
	LetterId   int     `json:"letter_id"`   //номер литера
	ComplexId  int     `json:"complex_id"`  // номер ЖК
}
