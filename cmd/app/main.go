// Arquivo principal do programa (entrypoint)
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"

    "github.com/seu-usuario/meu-projeto-go/internal/hello"
	"github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
)

// Função principal do programa
func main() {
    fmt.Println("🚀 Meu primeiro projeto em Go, com estrutura de mercado!")
    hello.SayHello()

    // variável para armazenar o número inserido pelo usuário;
    var n int

    // Conteúdo exibido ao usuário
    fmt.Print("Digite um número para calcular o Fibonacci: ")
    
    // solicita a entrada de dados do usuário
    fmt.Scanln(&n)

    fmt.Printf("Fibonacci de %d é: %d\n", n, fibonacci.Fibonacci(n))
}
 

