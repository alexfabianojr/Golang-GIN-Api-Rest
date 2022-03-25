package main

import (
	"github.com/alexfabianojr/Golang-GIN-Api-Rest/database"
	"github.com/alexfabianojr/Golang-GIN-Api-Rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
