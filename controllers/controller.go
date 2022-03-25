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

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func BuscaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(
			http.StatusNotFound, gin.H{
				"Not Found": "Aluno não encontrado",
			})
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Deletado"})
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)

	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")

	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"Not Found": "Aluno nao encontrado",
			})

		return
	}

	c.JSON(http.StatusOK, aluno)
}
