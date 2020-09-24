package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

func main() {
	anakin := pessoa{"Anakin", 20}
	
	fmt.Println("Lado da força ", anakin)
	mudeMe(&anakin)
	fmt.Println("Lado nego da força ", anakin)

}

func mudeMe(p *pessoa) {
	p.nome = "Darth Vader"
	p.idade = 40
}
