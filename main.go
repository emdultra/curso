package main

import (
	"fmt"

	"github.com/emdultra/curso/arquivos"
)

func main() {

}

func calculadora() {
	valor, erro := arquivos.Calculadora(125, "/", 0)
	if erro != nil {
		fmt.Println(erro.Error())
	}
	fmt.Println(valor)
}

func conversorArray() {
	arquivos.Conversor()
}

func ordernarSlice() {
	arquivos.Ordenar()
}

func contadorMap() {
	arquivos.ContadorLetras()
}
