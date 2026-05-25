package main

import (
	"strings"
)

func normalizeStatus(status string) (string, bool) {
	s := strings.ToLower(strings.TrimSpace(status))
	if s == "pago" {
		return "Pago", true
	}
	if s == "atrasado" {
		return "Atrasado", true
	}
	return "", false
}

func listAll() []Record {
	db := loadData()
	list := make([]Record, 0, len(db.Records))
	for _, r := range db.Records {
		list = append(list, r)
	}
	return list
}

func findByID(id string) (Record, bool) {
	db := loadData()
	r, ok := db.Records[id]
	return r, ok
}

func add(name, status string) (bool, string) {
	db := loadData()

	normalizedStatus, isValid := normalizeStatus(status)
	if !isValid {
		return false, "Situação inválida. Use apenas 'pago' ou 'atrasado'"
	}

	id := generateID()
	for {
		if _, exists := db.Records[id]; !exists {
			break
		}
		id = generateID()
	}

	db.Records[id] = Record{
		ID:       id,
		Name:     name,
		Status:   normalizedStatus,
	}
	saveData(db)

	return true, "Cadastro adicionado com sucesso. ID gerado: " + id
}

func remove(id string) (bool, string) {
	db := loadData()

	if _, exists := db.Records[id]; !exists {
		return false, "Cadastro não encontrado"
	}

	delete(db.Records, id)
	saveData(db)

	return true, "Cadastro removido com sucesso"
}

func updateStatus(id, newStatus string) (bool, string) {
	db := loadData()

	record, exists := db.Records[id]
	if !exists {
		return false, "Cadastro não encontrado"
	}

	normalizedStatus, isValid := normalizeStatus(newStatus)
	if !isValid {
		return false, "Situação inválida. Use apenas 'pago' ou 'atrasado'"
	}

	record.Status = normalizedStatus
	db.Records[id] = record
	saveData(db)

	return true, "Situação alterada com sucesso"
}

func statistics() (total, paid, overdue int) {
	db := loadData()

	for _, r := range db.Records {
		total++
		if r.Status == "Pago" {
			paid++
		} else if r.Status == "Atrasado" {
			overdue++
		}
	}

	return total, paid, overdue
}
