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

		optionStr, _ := reader.ReadString('\n')
		option := strings.TrimSpace(optionStr)

		switch option {
		case "1":
			list := listAll()
			if len(list) == 0 {
				fmt.Println("nenhum cadastro encontrado")
			} else {
				fmt.Println("--------------------------------------------------")
				fmt.Println("ID | Nome | Situacao")
				fmt.Println("--------------------------------------------------")
				for _, r := range list {
					color := ""
					if r.Status == "Pago" {
						color = "\033[32m"
					} else if r.Status == "Atrasado" {
						color = "\033[31m"
					}
					reset := "\033[0m"
					fmt.Printf("%s | %s | %s%s%s\n", r.ID, r.Name, color, r.Status, reset)
				}
			}

		case "2":
			fmt.Print("ID: ")
			idStr, _ := reader.ReadString('\n')
			id := strings.TrimSpace(idStr)
			
			r, found := findByID(id)
			if found {
				fmt.Printf("Encontrado: ID=%s, Nome=%s, Situacao=%s\n", r.ID, r.Name, r.Status)
			} else {
				fmt.Println("cadastro nao encontrado")
			}

		case "3":
			fmt.Print("Nome: ")
			nameStr, _ := reader.ReadString('\n')
			name := strings.TrimSpace(nameStr)

			fmt.Print("Situacao: ")
			statusStr, _ := reader.ReadString('\n')
			status := strings.TrimSpace(statusStr)

			_, msg := add(name, status)
			fmt.Println(msg)

		case "4":
			fmt.Print("ID: ")
			idStr, _ := reader.ReadString('\n')
			id := strings.TrimSpace(idStr)

			_, msg := remove(id)
			fmt.Println(msg)

		case "5":
			fmt.Print("ID: ")
			idStr, _ := reader.ReadString('\n')
			id := strings.TrimSpace(idStr)

			fmt.Print("Nova situacao: ")
			statusStr, _ := reader.ReadString('\n')
			status := strings.TrimSpace(statusStr)

			_, msg := updateStatus(id, status)
			fmt.Println(msg)

		case "6":
			total, paid, overdue := statistics()
			fmt.Printf("Total: %d\n", total)
			fmt.Printf("Pagos: %d\n", paid)
			fmt.Printf("Atrasados: %d\n", overdue)

			if total > 0 {
				maxBar := 40
				barPaid := int((float64(paid) / float64(total)) * float64(maxBar))
				barOverdue := int((float64(overdue) / float64(total)) * float64(maxBar))

				fmt.Printf("Pagos:     [%s]\n", strings.Repeat("#", barPaid))
				fmt.Printf("Atrasados: [%s]\n", strings.Repeat("#", barOverdue))
			}

		case "0":
			fmt.Println("encerrando.")
			return

		default:
			fmt.Println("opcao invalida")
		}
	}
}
