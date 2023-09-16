package arquivos

import (
	"fmt"
	"os"
	"strings"
)

func ContadorLetras() {
	// ContadorLetras usando map
	palavras := os.Args[1:]
	letras := contarLetras(palavras)
	imprimir(letras)
}

func contarLetras(palavras []string) map[string]int {
	letras := make(map[string]int)

	for _, palavra := range palavras {
		inicial := strings.ToUpper(string(palavra[0]))
		contador, encontrado := letras[inicial]
		if encontrado {
			letras[inicial] = contador + 1
		} else {
			letras[inicial] = 1
		}
	}
	return letras
}

func imprimir(palavras map[string]int) {
	fmt.Println("Contagem de palavras iniciadas em cada letra:")
	for inicial, contador := range palavras {
		fmt.Printf("%s = %d\n", inicial, contador)
	}
}
