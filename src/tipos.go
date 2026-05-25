package main

type Cadastro struct {
	ID       string
	Nome     string
	Situacao string
}

type Banco struct {
	Cadastros map[string]Cadastro
}
