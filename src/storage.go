package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func loadData() Database {
	db := Database{
		Records: make(map[string]Record),
	}

	os.MkdirAll("db", 0755)

	data, err := os.ReadFile("db/dados.json")
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Erro ao ler o arquivo dados.json:", err)
		}
		return db
	}

	err = json.Unmarshal(data, &db)
	if err != nil {
		fmt.Println("Erro ao interpretar dados.json:", err)
		return db
	}

	if db.Records == nil {
		db.Records = make(map[string]Record)
	}

	return db
}

func saveData(db Database) {
	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		fmt.Println("Erro ao formatar os dados:", err)
		return
	}

	os.MkdirAll("db", 0755)

	err = os.WriteFile("db/dados.json", data, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo dados.json:", err)
	}
}
