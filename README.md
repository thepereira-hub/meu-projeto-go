# 📘 Aula – Meu Projeto em Go

Este repositório é o **projeto base da disciplina**.  
Aqui vamos aprender a criar, rodar, buildar e publicar um projeto em Go, já usando uma **estrutura de pastas padrão de mercado**.

---

## 📂 Estrutura do Projeto
```
meu-projeto-go/
├── cmd/app/          -> ponto de entrada da aplicação (main.go)
├── internal/hello/   -> código interno, não exportável
├── go.mod            -> arquivo de módulo Go
└── README.md         -> instruções do projeto
```

- **cmd/app/** → onde fica o `main.go`, ponto de entrada do programa.  
- **internal/** → pacotes internos, só podem ser usados dentro do projeto.  
- **go.mod** → define que esse diretório é um módulo Go.  
- **README.md** → instruções passo a passo da aula.  

---

## 🚀 Passo a Passo da Aula

### 1. Clonar ou baixar este repositório /*OK*/
```bash
git clone URL_PROJETO
cd meu-projeto-go
```

Se estiver usando o ZIP entregue, basta descompactar e entrar na pasta.

### 2. Rodar o projeto
```bash
go run ./cmd/app
```

➡️ Saída esperada:
```
🚀 Meu primeiro projeto em Go com estrutura de mercado!
Hello Word in FacINpro! 👋
```

### 3. Gerar um executável (build)
```bash
go build -o meuapp ./cmd/app
./meuapp
```

### 4. Publicar no GitHub
```bash
git init
git add .
git commit -m "primeiro commit: projeto base em Go"
git branch -M main
git remote add origin https://github.com/<seu-usuario>/meu-projeto-go.git
git push -u origin main
```

---

## 🎯 Desafios realizados:
- Parte 1 - Fibonacci:  
  - Foi criada uma nova função `Fibonacci(n int) int` no arquivo `internal/fibonacci/fibonacci.go` que retorna o n-ésimo número da sequência de Fibonacci.
  - Ela stá sendo chamada na função `main.go` e imprimindo o resultado.
  - Ao final, foi: executada, testada(aprovada) e realizado commit do código no GitHub.

- Parte 2 - Cálculo IMC (índice de massa corporal):
  - Foi criada uma nova função `CalculaIMC(peso float64, altura float64)float64` no arquivo `internal/imc/calculoIMC.go` que calcula o índice de massa corporal.
  -Ela está sendo chamada na função `main.go` e imprimindo o resultado.

  - OBS: foi adaptado a função main, criando um menu para que sejam selecionadas quais funções executar.
