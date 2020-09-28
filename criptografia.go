package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Digite uma mensagem qualquer:")
	// Scanner usado para conseguir pegar mensagens com espaços entre palavras.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto := scanner.Text()
	fmt.Scanln(&texto)
	fmt.Println("\nDigite a cifra desejada:")
	var cifra rune = 0
	fmt.Scanln(&cifra)

	cripto := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+cifra)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+cifra)%26
		}
		return r
	}

	descripto := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'-cifra)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'-cifra)%26
		}
		return r
	}

	fmt.Println("\nCifra utilizada: \n", cifra)
	fmt.Println("\nTexto criptografado: \n", strings.Map(cripto, texto))
	fmt.Println("\nTexto descriptografado: \n", strings.Map(descripto, strings.Map(cripto, texto)))
}

// Não ficou bem como desejado. Tentarei em outras linguagens depois.
