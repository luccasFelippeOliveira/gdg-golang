package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Cidade struct {
	Nome string `json:nome`
}

// Interfaces, vamos implementar o to String.
// A interface Stringer define o metodo String(), vamos fazer Cidade implemntar essa interface
func (c Cidade) String() string {
	return "Cidade = " + c.Nome
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

	fmt.Println(obj)
}
