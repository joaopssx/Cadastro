package main

import (
	"strings"
)

func normalizarSituacao(situacao string) (string, bool) {
	s := strings.ToLower(strings.TrimSpace(situacao))
	if s == "pago" {
		return "Pago", true
	}
	if s == "atrasado" {
		return "Atrasado", true
	}
	return "", false
}

func listarTodos() []Cadastro {
	banco := carregarDados()
	lista := make([]Cadastro, 0, len(banco.Cadastros))
	for _, c := range banco.Cadastros {
		lista = append(lista, c)
	}
	return lista
}

func buscarPorID(id string) (Cadastro, bool) {
	banco := carregarDados()
	c, ok := banco.Cadastros[id]
	return c, ok
}

func adicionar(nome, situacao string) (bool, string) {
	banco := carregarDados()

	sitNormalizada, valida := normalizarSituacao(situacao)
	if !valida {
		return false, "Situação inválida. Use apenas 'pago' ou 'atrasado'"
	}

	id := gerarID()
	for {
		if _, existe := banco.Cadastros[id]; !existe {
			break
		}
		id = gerarID()
	}

	banco.Cadastros[id] = Cadastro{
		ID:       id,
		Nome:     nome,
		Situacao: sitNormalizada,
	}
	salvarDados(banco)

	return true, "Cadastro adicionado com sucesso. ID gerado: " + id
}

func remover(id string) (bool, string) {
	banco := carregarDados()

	if _, existe := banco.Cadastros[id]; !existe {
		return false, "Cadastro não encontrado"
	}

	delete(banco.Cadastros, id)
	salvarDados(banco)

	return true, "Cadastro removido com sucesso"
}

func alterarSituacao(id, novaSituacao string) (bool, string) {
	banco := carregarDados()

	cadastro, existe := banco.Cadastros[id]
	if !existe {
		return false, "Cadastro não encontrado"
	}

	sitNormalizada, valida := normalizarSituacao(novaSituacao)
	if !valida {
		return false, "Situação inválida. Use apenas 'pago' ou 'atrasado'"
	}

	cadastro.Situacao = sitNormalizada
	banco.Cadastros[id] = cadastro
	salvarDados(banco)

	return true, "Situação alterada com sucesso"
}

func estatisticas() (total, pagos, atrasados int) {
	banco := carregarDados()

	for _, c := range banco.Cadastros {
		total++
		if c.Situacao == "Pago" {
			pagos++
		} else if c.Situacao == "Atrasado" {
			atrasados++
		}
	}

	return total, pagos, atrasados
}
