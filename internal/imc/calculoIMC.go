// autor: Kleverton Pereira
// data: 2025-09-11
// descrição: Esta função calcula o Índice de Massa Corporal (IMC) dado o peso e a altura.
package imc

func CalculaIMC(peso float64, altura float64) float64 {
	return peso / (altura * altura)
}
	
