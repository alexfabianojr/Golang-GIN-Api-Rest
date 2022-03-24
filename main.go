package main

import (
	"github.com/alexfabianojr/Golang-GIN-Api-Rest/models"
	"github.com/alexfabianojr/Golang-GIN-Api-Rest/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Alex", CPF: "1111111", RG: "2222222"},
		{Nome: "Manu", CPF: "2222222", RG: "4444444"},
	}
	routes.HandleRequests()
}
