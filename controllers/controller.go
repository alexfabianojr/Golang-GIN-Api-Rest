package controllers

import (
	"net/http"

	"github.com/alexfabianojr/Golang-GIN-Api-Rest/database"
	"github.com/alexfabianojr/Golang-GIN-Api-Rest/models"
	"github.com/gin-gonic/gin"
)

func CriarNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"Error: ": err.Error(),
			})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(
		http.StatusCreated,
		aluno,
	)
}

func ExibeAlex(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   1,
		"nome": "Alex",
	})
}

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(
		200,
		gin.H{
			"API diz: ": "E ai " + nome + ", tudo beleza?",
		},
	)
}
