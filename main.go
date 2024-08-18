package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
}

func main() {
	p := Pessoa{
		Nome: "Teste",
	}

	fmt.Println(p)
}
