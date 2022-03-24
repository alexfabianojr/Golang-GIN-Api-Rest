package routes

import (
	"github.com/alexfabianojr/Golang-GIN-Api-Rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()

	router.GET("/alunos", controllers.ExibeTodosAlunos)
	router.GET("/:nome", controllers.Saudacao)

	router.Run()
}
