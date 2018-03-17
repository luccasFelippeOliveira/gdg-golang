package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

type Cidade struct {
	Nome string `json:"nome"`
}

type Temperatura struct {
	Temp       float64 `json:"temp"`
	Pressao    float64 `json:"pressure"`
	Humidade   int     `json:"humidity"`
	TempMinima float64 `json:"temp_min"`
	TempMaxima float64 `json:"temp_max"`
	AlturaMar  float64 `json:"sea_level"`
	AlturaChao float64 `json:"grnd_level"`
}

// Interfaces, vamos implementar o to String.
// A interface Stringer define o metodo String(), vamos fazer Cidade implemntar essa interface
func (c Cidade) String() string {
	return "Cidade = " + c.Nome
}

func callWeatherAPI(city string) *Temperatura {
	baseURL := "http://api.openweathermap.org/data/2.5/weather?"

	// Fazer o encode para url, go é nativamente utf-8, requisições GET não!
	values := url.Values{}
	values.Set("q", city)
	values.Set("appid", "38bd80b6337d7010e7ab270280ee784d") // Nossa api key

	// Call api.

	// Criar uma requisicao
	req, err := http.NewRequest("GET", baseURL+values.Encode(), nil)
	if err != nil {
		fmt.Println(err)
		return nil // Do nothing
	}

	// Criar um cliente
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error 43")
		fmt.Println(err)
		return nil // Do nothin
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error 52")
		fmt.Println(err)
		return nil
	}

	var obj map[string]json.RawMessage
	if err := json.Unmarshal(body, &obj); err != nil {
		fmt.Println(err)
		return nil
	}

	var temperatura Temperatura
	if err := json.Unmarshal(obj["main"], &temperatura); err != nil {
		fmt.Println(err)
		return nil
	}

	return &temperatura
}

func main() {
	// Le o arquivo de cidades.
	data, err := ioutil.ReadFile("cidades.json")
	if err != nil {
		panic(err)
	}

	// Slices sao como listas dinamicas.
	var obj = make([]Cidade, 1)

	// Deserializando o json que lemos no arquivo.
	if err := json.Unmarshal(data, &obj); err != nil {
		panic(err)
	}

	// Criamos o grupo de espera
	var waitGroup sync.WaitGroup

	// Criamos o canal
	apiResponse := make(chan string)

	for _, cidade := range obj {
		waitGroup.Add(1)
		// Chamamos a funcao callWeatherApi de forma concorrente
		go func(cidade string) {
			defer waitGroup.Done()
			temperatura := callWeatherAPI(cidade)
			if temperatura == nil {
				apiResponse <- cidade + ": Nao encontrado!"
			} else {
				apiResponse <- cidade + ": " + strconv.FormatFloat(temperatura.Temp, 'f', -1, 64)
			}
		}(cidade.Nome)
	}

	// Bloqueia main de terminar, aguardando as respostas da api.
	go func() {
		waitGroup.Wait()
		close(apiResponse)
	}()

	for r := range apiResponse {
		fmt.Println(r)
	}
}
