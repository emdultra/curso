package arquivos

import "errors"

func Calculadora(valor1 float64, operacao string, valor2 float64) (float64, error) {
	if operacao == "/" && valor2 == 0 {
		return 0.0, errors.New("error: Divisão por zero")
	}
	resp := 0.0
	switch operacao {
	case "-":
		resp = valor1 - valor2
	case "+":
		resp = valor1 + valor2
	case "*":
		resp = valor1 * valor2
	case "/":
		resp = valor1 / valor2
	default:
		return 0.0, errors.New("error: Operação não reconhecida")
	}
	return resp, nil
}
