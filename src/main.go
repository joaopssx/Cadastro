package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1 - listar todos")
		fmt.Println("2 - buscar por ID")
		fmt.Println("3 - adicionar")
		fmt.Println("4 - remover")
		fmt.Println("5 - alterar situacao")
		fmt.Println("6 - estatisticas")
		fmt.Println("0 - sair")
		fmt.Print("Opcao: ")

		opcaoStr, _ := reader.ReadString('\n')
		opcao := strings.TrimSpace(opcaoStr)

		switch opcao {
		case "1":
			lista := listarTodos()
			if len(lista) == 0 {
				fmt.Println("nenhum cadastro encontrado")
			} else {
				fmt.Println("--------------------------------------------------")
				fmt.Println("ID | Nome | Situacao")
				fmt.Println("--------------------------------------------------")
				for _, c := range lista {
					cor := ""
					if c.Situacao == "Pago" {
						cor = "\033[32m"
					} else if c.Situacao == "Atrasado" {
						cor = "\033[31m"
					}
					reset := "\033[0m"
					fmt.Printf("%s | %s | %s%s%s\n", c.ID, c.Nome, cor, c.Situacao, reset)
				}
			}

		case "2":
			fmt.Print("ID: ")
			idStr, _ := reader.ReadString('\n')
			id := strings.TrimSpace(idStr)
			
			c, encontrado := buscarPorID(id)
			if encontrado {
				fmt.Printf("Encontrado: ID=%s, Nome=%s, Situacao=%s\n", c.ID, c.Nome, c.Situacao)
			} else {
				fmt.Println("cadastro nao encontrado")
			}

		case "3":
			fmt.Print("Nome: ")
			nomeStr, _ := reader.ReadString('\n')
			nome := strings.TrimSpace(nomeStr)

			fmt.Print("Situacao: ")
			sitStr, _ := reader.ReadString('\n')
			situacao := strings.TrimSpace(sitStr)

			_, msg := adicionar(nome, situacao)
			fmt.Println(msg)

		case "4":
			fmt.Print("ID: ")
			idStr, _ := reader.ReadString('\n')
			id := strings.TrimSpace(idStr)

			_, msg := remover(id)
			fmt.Println(msg)

		case "5":
			fmt.Print("ID: ")
			idStr, _ := reader.ReadString('\n')
			id := strings.TrimSpace(idStr)

			fmt.Print("Nova situacao: ")
			sitStr, _ := reader.ReadString('\n')
			situacao := strings.TrimSpace(sitStr)

			_, msg := alterarSituacao(id, situacao)
			fmt.Println(msg)

		case "6":
			total, pagos, atrasados := estatisticas()
			fmt.Printf("Total: %d\n", total)
			fmt.Printf("Pagos: %d\n", pagos)
			fmt.Printf("Atrasados: %d\n", atrasados)

			if total > 0 {
				maxBar := 40
				barPagos := int((float64(pagos) / float64(total)) * float64(maxBar))
				barAtrasados := int((float64(atrasados) / float64(total)) * float64(maxBar))

				fmt.Printf("Pagos:     [%s]\n", strings.Repeat("#", barPagos))
				fmt.Printf("Atrasados: [%s]\n", strings.Repeat("#", barAtrasados))
			}

		case "0":
			fmt.Println("encerrando.")
			return

		default:
			fmt.Println("opcao invalida")
		}
	}
}
