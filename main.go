package main

import (
	"clean_code/database"
	pedidosHandlers "clean_code/handlers/pedidos"
	tastesHandlers "clean_code/handlers/tastes"
	handlers "clean_code/handlers/users"
	pedidosRepositories "clean_code/internal/repositories/pedidos"

	tastesRepositories "clean_code/internal/repositories/tastes"
	repositories "clean_code/internal/repositories/users"
	pedidosServices "clean_code/internal/services/pedidos"
	tastesServices "clean_code/internal/services/tastes"
	services "clean_code/internal/services/users"
	"fmt"

	cantHandlers "clean_code/handlers/cantidad"
	cantRepositories "clean_code/internal/repositories/cantidad"
	cantServices "clean_code/internal/services/cantidad"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.InitConection()
	if err != nil {
		panic("failed to connect database")
	}

	userRepo := repositories.Repository{DB: db}
	userServices := services.Services{Repo: &userRepo}
	userHandler := handlers.UserHandler{Services: &userServices}

	pedRepo := pedidosRepositories.PedidosRepository{DB: db}
	pedServices := pedidosServices.PedidosServices{RepoPedidos: &pedRepo, RepoUser: &userRepo}
	pedHandler := pedidosHandlers.PedidosHandler{Services: &pedServices}
	fmt.Println(pedHandler)

	tasteRepo := tastesRepositories.TasteRepository{DB: db}
	tasteServices := tastesServices.TasteServices{Repo: &tasteRepo}
	tasteHandler := tastesHandlers.TasteHandler{Services: &tasteServices}
	fmt.Println(tasteHandler)

	cantRepo := cantRepositories.CantidadRepository{DB: db}
	cantServices := cantServices.CantidadServices{RepoCantidad: &cantRepo}
	cantHandler := cantHandlers.CantidadHandler{CantidadService: &cantServices}

	router := gin.Default()

	user := router.Group("user")

	user.GET(":id", userHandler.GetUserByID)
	user.GET("/all", userHandler.List)
	user.POST("", userHandler.Create)
	user.PUT(":id", userHandler.Update)
	user.DELETE(":id", userHandler.Delete)

	/*
	   	comment := router.Group("comment")
	   /*
	   	comment.GET(":user_id", commHandler.GetCommentByUserId)
	   	comment.GET("/all", commHandler.ListComments)
	   	comment.POST("", commHandler.NewComment)
	   	comment.PUT(":ID", commHandler.Update)
	   	comment.DELETE(":id", commHandler.DeleteComment)
	*/
	taste := router.Group("taste")

	taste.POST("", tasteHandler.Create)

	cantidad := router.Group("cantidad")

	cantidad.POST("", cantHandler.Create)
	cantidad.PUT("/:id", cantHandler.Update)

	router.Run() // listen and serve on 0.0.0.0:8080

}
