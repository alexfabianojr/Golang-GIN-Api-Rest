package routes

import (
	"github.com/alexfabianojr/Golang-GIN-Api-Rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()

	router.GET("/alunos", controllers.ExibeTodosAlunos)
	router.POST("/alunos", controllers.CriarNovoAluno)
	router.GET("/aluno/:id", controllers.BuscaAlunoPorId)
	router.DELETE("/aluno/:id", controllers.DeletarAluno)
	router.PUT("/aluno/:id", controllers.EditaAluno)
	router.GET("/aluno/cpf/:cpf", controllers.BuscaAlunoPorCpf)

	router.Run()
}
