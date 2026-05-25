package main

type Record struct {
	ID       string
	Name     string
	Status   string
}

type Database struct {
	Records map[string]Record
}
