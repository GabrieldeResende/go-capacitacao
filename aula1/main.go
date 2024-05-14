package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// import "net/http"

type Carro struct {
	Nome string
	Modelo string
	Ano int
}

func (c Carro) Andar() {
	fmt.Println("O carro, " + c.Nome + " Esta andando")
}

func (c Carro) Parar() {
	fmt.Println("O carro, " + c.Nome + " Esta parando")
}

func main() {

	carro1 := Carro{Nome: "Ford", Modelo: "Mustang", Ano: 1960}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		json.NewEncoder(w).Encode(carro1)
	})

	http.ListenAndServe(":8080", nil)

	// http.ListenAndServe(":3333", nil)
}