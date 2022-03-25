package database

import (
	"fmt"
	"log"

	"github.com/alexfabianojr/Golang-GIN-Api-Rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados!")
	}

	DB.AutoMigrate(&models.Aluno{}) //& endereço de memória | {} instância dele
	fmt.Println("Conexao database iniciada")
}
