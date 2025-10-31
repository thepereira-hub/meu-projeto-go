package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)
const nomeArquivo = "aula.txt"

func main() {
	fmt.Println("== Aula: Manipulação de Arquivos e E/S em Go ==")

	// 1) Criação de arquivo: cria um novo arquivo chamado aula.txt
	if err := criarArquivo(nomeArquivo); err != nil {
		// Se o arquivo já existir, aviso e sigo adiante sem sobrescrever.
		if errors.Is(err, os.ErrExist) {
			fmt.Printf("Aviso: %s já existia (não sobrescrevi).\n", nomeArquivo)
		} else {
			check(err)
		}
	} else {
		fmt.Printf("OK: criei %s.\n", nomeArquivo)
	}

	// 2) Escrever no arquivo: adiciona uma string ao arquivo criado (append)
	texto := "Olá! Esta é a nossa aula de manipulação de arquivos em Go.\n" +
		"Linha 2: vamos ler isso já já.\n"
	check(appendTexto(nomeArquivo, texto))
	fmt.Println("OK: escrevi (append) no arquivo.")

	// 3) Leitura de arquivo: lê o conteúdo completo do arquivo
	dados, err := os.ReadFile(nomeArquivo)
	check(err)
	fmt.Println("== Conteúdo lido com os.ReadFile (completo) ==")
	fmt.Print(string(dados))

	// 4) Leitura de arquivo com buffer: lê o conteúdo em partes de tamanho definido
	fmt.Println("\n== Leitura com buffer (tamanho 8 bytes) ==")
	check(lerComBuffer(nomeArquivo, 8))

	// 5) Remoção de arquivo: descomente para remover ao final
	// fmt.Println("Removendo arquivo...")
	// check(os.Remove(nomeArquivo))
	// fmt.Println("OK: arquivo removido.")
}

// criarArquivo cria o arquivo apenas se NÃO existir (O_EXCL evita sobrescrever).
func criarArquivo(caminho string) error {
	// Flags:
	// - O_CREATE: cria se não existir
	// - O_EXCL: falha se já existir (protege contra sobrescrita acidental)
	// - O_WRONLY: abre para escrita
	f, err := os.OpenFile(caminho, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	return f.Close()
}

// appendTexto abre o arquivo para escrita em modo "append" (adicionar ao fim).
func appendTexto(caminho, texto string) error {
	// O_APPEND: sempre escreve no final
	// O_WRONLY: abre apenas para escrita
	f, err := os.OpenFile(caminho, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(texto)
	return err
}

// lerComBuffer lê o arquivo em pedaços de tamanho 'tamanhoBuffer', mostrando cada parte.
func lerComBuffer(caminho string, tamanhoBuffer int) error {
	arq, err := os.Open(caminho)
	if err != nil {
		return err
	}
	defer arq.Close()

	// bufio.Reader com buffer interno do tamanho desejado
	r := bufio.NewReaderSize(arq, tamanhoBuffer)

	// buffer de leitura (tamanho controlado por nós)
	buf := make([]byte, tamanhoBuffer)
	parte := 1

	for {
		n, err := r.Read(buf)
		if n > 0 {
			fmt.Printf("Parte %02d (%d bytes): %q\n", parte, n, string(buf[:n]))
			parte++
		}
		if err == io.EOF {
			break // chegou ao fim do arquivo
		}
		if err != nil {
			return err // qualquer outro erro interrompe
		}
	}
	return nil
}

// check simplifica o tratamento de erro neste exemplo didático.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

