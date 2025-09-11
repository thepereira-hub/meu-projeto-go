//  Função Fibonacci em Go
//  Autor: Kleverton Pereira
//  Data: 2025-09-11
//  Descrição: Esta função calcula o n-ésimo número na sequência de Fibonacci usando recursão.
 package fibonacci

// Início da função fibonacci:
func Fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}

/* Como funciona? 


A função Fibonacci(n int) int calcula o n-ésimo número da sequência de Fibonacci, que é uma sequência onde cada número é a soma dos dois anteriores.

De forma simples:

Se você pedir o Fibonacci de 0 ou 1, ele retorna o próprio número.
Para qualquer outro número, ele soma o Fibonacci do número anterior com o do número antes dele.
Exemplo:

Fibonacci(4) = Fibonacci(3) + Fibonacci(2)
Fibonacci(3) = Fibonacci(2) + Fibonacci(1)
E assim por diante, até chegar nos casos de 0 ou 1.
Ou seja, é como subir uma escada onde cada degrau depende dos dois anteriores!*/
