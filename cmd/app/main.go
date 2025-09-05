// Arquivo principal do programa (entrypoint)
// Convenção de mercado: colocar em cmd/<nome-app>/main.go
package main

// Importa os pacotes necessários
import (
	"fmt"

	fibonnaci "github.com/seu-usuario/meu-projeto-go/internal/fibonacci"
	"github.com/seu-usuario/meu-projeto-go/internal/hello"
)

// Função principal do programa
func main() {
	fmt.Println("🚀 Meu primeiro projeto em Go com estrutura de mercado!")
	hello.SayHello()

	var num int

	fmt.Print("Digita um número para calcular fibonnaci..: ")
	_, err := fmt.Scan(&num)

	if err != nil {
		fmt.Println("Error ao ler número.")
		return
	}

	fibResult, _ := fibonnaci.Fibonnaci(num)

	fmt.Println(fmt.Sprintf("O resultado Fibonnaci do número %d é %d.", num, fibResult))
}
