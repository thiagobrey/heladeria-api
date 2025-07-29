package main

import (
	"clean_code/database"
	"clean_code/handlers/cantidad/validators"
	pedidosHandlers "clean_code/handlers/pedidos"
	tastesHandlers "clean_code/handlers/tastes"
	handlers "clean_code/handlers/users"
	pedidosRepositories "clean_code/internal/repositories/pedidos"

	tastesRepositories "clean_code/internal/repositories/tastes"
	repositories "clean_code/internal/repositories/users"
	pedidosServices "clean_code/internal/services/pedidos"
	tastesServices "clean_code/internal/services/tastes"
	services "clean_code/internal/services/users"

	cantHandlers "clean_code/handlers/cantidad"
	cantRepositories "clean_code/internal/repositories/cantidad"
	cantServices "clean_code/internal/services/cantidad"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	db, err := database.InitConection()
	if err != nil {
		panic("failed to connect database")
	}

	userRepo := repositories.Repository{DB: db}
	userServices := services.Services{Repo: &userRepo}
	userHandler := handlers.UserHandler{Services: &userServices}

	tasteRepo := tastesRepositories.TasteRepository{DB: db}
	tasteServices := tastesServices.TasteServices{Repo: &tasteRepo}
	tasteHandler := tastesHandlers.TasteHandler{Services: &tasteServices}

	cantRepo := cantRepositories.CantidadRepository{DB: db}
	cantServices := cantServices.CantidadServices{RepoCantidad: &cantRepo}
	cantHandler := cantHandlers.CantidadHandler{CantidadService: &cantServices}

	pedRepo := pedidosRepositories.PedidosRepository{DB: db}
	pedServices := pedidosServices.PedidosServices{RepoPedidos: &pedRepo, RepoUser: &userRepo, RepoCantidad: &cantRepo, RepoTastes: &tasteRepo}
	pedHandler := pedidosHandlers.PedidosHandler{Services: &pedServices}

	router := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("count_validator", func(fl validator.FieldLevel) bool {
			return validators.CountValidator(fl, db)
		})
	}
	router.Use(gin.Recovery())

	user := router.Group("user")

	user.GET(":id", userHandler.GetById)
	user.GET("/all", userHandler.List)
	user.POST("", userHandler.Create)
	user.PUT(":id", userHandler.Update)
	user.DELETE(":id", userHandler.Delete)

	pedido := router.Group("pedido")
	//pedido.GET(":user_id", commHandler.GetpedidoByUserId)
	//pedido.GET("/all", commHandler.Listpedidos)
	pedido.POST("", pedHandler.Create)
	//edido.PUT(":ID", commHandler.Update)
	//pedido.DELETE(":id", commHandler.Deletepedido)

	taste := router.Group("taste")

	taste.POST("", tasteHandler.Create)

	cantidad := router.Group("cantidad")

	cantidad.POST("", cantHandler.Create)
	cantidad.PUT("/:id", cantHandler.Update)
	cantidad.GET("", cantHandler.List)

	router.Run() // listen and serve on 0.0.0.0:8080

}
