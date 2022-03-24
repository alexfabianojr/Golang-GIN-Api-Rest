package controllers

import (
	"github.com/alexfabianojr/Golang-GIN-Api-Rest/models"
	"github.com/gin-gonic/gin"
)

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
