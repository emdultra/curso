package arquivos

import (
	"fmt"
	"os"
	"strconv"
)

func Ordenar() {
	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for i, n := range entrada {
		numero, erro := strconv.Atoi(n)
		if erro != nil {
			fmt.Printf("%s não e um numero válido!", n)
			os.Exit(1)
		}
		numeros[i] = numero
	}

	fmt.Println(ordenacao(numeros))
}

func ordenacao(numeros []int) []int {
	if len(numeros) <= 1 {
		return numeros
	}

	n := make([]int, len(numeros))
	copy(n, numeros)

	indice := len(n) / 2
	pivo := n[indice]

	n = append(n[:indice], n[indice+1:]...)

	menores, maiores := particionar(n, pivo)

	return append(append(ordenacao(menores), pivo), ordenacao(maiores)...)
}

func particionar(numeros []int, pivo int) (menores []int, maiores []int) {

	for _, n := range numeros {
		if n <= pivo {
			menores = append(menores, n)
		} else {
			maiores = append(maiores, n)
		}
	}

	return menores, maiores
}
