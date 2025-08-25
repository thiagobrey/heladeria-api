package main

import (
	"clean_code/database"
	"clean_code/handlers/cantidad/validators"
	pedidosHandlers "clean_code/handlers/pedidos"
	tastesHandlers "clean_code/handlers/tastes"
	handlers "clean_code/handlers/users"
	pedidosRepositories "clean_code/internal/repositories/pedidos"
	"clean_code/middleware"

	tastesRepositories "clean_code/internal/repositories/tastes"
	repositories "clean_code/internal/repositories/users"
	pedidosServices "clean_code/internal/services/pedidos"
	tastesServices "clean_code/internal/services/tastes"
	services "clean_code/internal/services/users"

	cantHandlers "clean_code/handlers/cantidad"
	cantRepositories "clean_code/internal/repositories/cantidad"
	cantServices "clean_code/internal/services/cantidad"

	admHandlers "clean_code/handlers/admins"
	adminRepositories "clean_code/internal/repositories/admins"
	adminServices "clean_code/internal/services/admins"

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
	tasteRepo := tastesRepositories.TasteRepository{DB: db}
	cantRepo := cantRepositories.CantidadRepository{DB: db}
	pedRepo := pedidosRepositories.PedidosRepository{DB: db}
	admRepo := adminRepositories.AdminRepository{DB: db}

	userServices := services.Services{Repo: &userRepo, RepoTastes: &tasteRepo}
	tasteServices := tastesServices.TasteServices{Repo: &tasteRepo}
	cantServices := cantServices.CantidadServices{RepoCantidad: &cantRepo}
	pedServices := pedidosServices.PedidosServices{RepoPedidos: &pedRepo, RepoUser: &userRepo, RepoCantidad: &cantRepo, RepoTastes: &tasteRepo}
	admServices := adminServices.AdminServices{RepoAdmin: &admRepo}

	userHandler := handlers.UserHandler{Services: &userServices}
	tasteHandler := tastesHandlers.TasteHandler{Services: &tasteServices}
	cantHandler := cantHandlers.CantidadHandler{CantidadService: &cantServices}
	pedHandler := pedidosHandlers.PedidosHandler{Services: &pedServices}
	admHandler := admHandlers.AdminHandler{AdminService: &admServices}

	if err := adminRepositories.SeedAdmin(db); err != nil {
		panic("failed to seed admin user: " + err.Error())
	}

	router := gin.New()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("count_validator", func(fl validator.FieldLevel) bool {
			return validators.CountValidator(fl, db)
		})
	}

	router.Use(middleware.AuthMiddleware())

	

	router.Use(gin.Recovery())

	user := router.Group("user")

	user.GET(":id", userHandler.GetById)
	user.GET("/all", userHandler.List)
	user.POST("", userHandler.Create)
	user.PUT(":id", userHandler.Update)
	user.DELETE(":id", userHandler.Delete)

	pedido := router.Group("pedido")
	pedido.GET(":user_id", pedHandler.GetCommentByUserId) // cambiar el response
	pedido.GET("/all", pedHandler.List)                   // cambiar el response
	pedido.POST("", pedHandler.Create)
	pedido.PUT(":ID", pedHandler.Update) // no guarda lo updateado
	pedido.DELETE(":id", pedHandler.Delete)

	taste := router.Group("taste")
	taste.GET(":id", tasteHandler.GetById)
	taste.POST("", tasteHandler.Create)
	taste.DELETE(":id", tasteHandler.Delete)
	taste.PUT(":id", tasteHandler.Update)
	taste.GET("", tasteHandler.List)

	cantidad := router.Group("cantidad")
	cantidad.GET(":id", cantHandler.GetById)
	cantidad.POST("", cantHandler.Create)
	cantidad.PUT("/:id", cantHandler.Update)
	cantidad.GET("", cantHandler.List)
	cantidad.DELETE(":id", cantHandler.Delete)

	admin := router.Group("admin")
	admin.POST("", admHandler.Create)
	admin.GET(":id", admHandler.GetById)
	admin.PUT(":id", admHandler.Update)
	admin.DELETE(":id", admHandler.Delete)

	router.Run() // listen and serve on 0.0.0.0:8080

}
