package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func carregarDados() Banco {
	banco := Banco{
		Cadastros: make(map[string]Cadastro),
	}

	os.MkdirAll("db", 0755)

	dados, err := os.ReadFile("db/dados.json")
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Erro ao ler o arquivo dados.json:", err)
		}
		return banco
	}

	err = json.Unmarshal(dados, &banco)
	if err != nil {
		fmt.Println("Erro ao interpretar dados.json:", err)
		return banco
	}

	if banco.Cadastros == nil {
		banco.Cadastros = make(map[string]Cadastro)
	}

	return banco
}

func salvarDados(banco Banco) {
	dados, err := json.MarshalIndent(banco, "", "  ")
	if err != nil {
		fmt.Println("Erro ao formatar os dados:", err)
		return
	}

	os.MkdirAll("db", 0755)

	err = os.WriteFile("db/dados.json", dados, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo dados.json:", err)
	}
}
