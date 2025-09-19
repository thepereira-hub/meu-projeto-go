// Arquivo principal do programa (entrypoint)
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"
    
	"github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
    "github.com/seu-usuario/meu-projeto-go/internal/imc"
)

// Função principal do programa
func main() {
    fmt.Println("🚀 Meu projeto em Go, com estrutura de mercado!")

    // Menu de opções para o usuário:
    fmt.Println("\nEscolha uma opção:")
    fmt.Println("1 - Calcular Fibonacci")
    fmt.Println("2 - Calcular IMC (Índice de Massa Corporal)")
    fmt.Print("Digite o número da opção desejada: ")

    var opcao int
    fmt.Scanln(&opcao)

    switch opcao {
    // case 1: contém o cálculo do Fibonacci
    case 1:
        // variável para armazenar o número inserido pelo usuário;
        var n int

        // Conteúdo exibido ao usuário
        fmt.Print("Digite um número para calcular o Fibonacci: ")
        
        // solicita a entrada de dados do usuário
        fmt.Scanln(&n)

        fmt.Printf("Fibonacci de %d é: %d\n", n, fibonacci.Fibonacci(n))
    // case 2: contém o cálculo do IMC
    case 2:
        var peso, altura float64
        fmt.Print("Digite seu peso (kg): ")
        fmt.Scanln(&peso)
        fmt.Print("Digite sua altura (m): ")
        fmt.Scanln(&altura)
        resultado := imc.CalculaIMC(peso, altura)
        fmt.Printf("Seu IMC é: %.2f\n", resultado)
    default:
        fmt.Println("Opção inválida.")
    }
}

 

