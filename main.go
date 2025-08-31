package main

import (
	"clean_code/database"
	"clean_code/handlers/cantidad/validators"
	pedidosHandlers "clean_code/handlers/pedidos"
	tastesHandlers "clean_code/handlers/tastes"
	handlers "clean_code/handlers/users"
	pedidosRepositories "clean_code/internal/repositories/pedidos"
	"clean_code/middleware"
	"fmt"

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

	sesHandlers "clean_code/handlers/sesions"
	sesRepositories "clean_code/internal/repositories/sesions"
	sesServices "clean_code/internal/services/sesions"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Error cargando .env")
	}

	db, err := database.InitConection()
	if err != nil {
		panic("failed to connect database")
	}

	userRepo := repositories.Repository{DB: db}
	tasteRepo := tastesRepositories.TasteRepository{DB: db}
	cantRepo := cantRepositories.CantidadRepository{DB: db}
	pedRepo := pedidosRepositories.PedidosRepository{DB: db}
	admRepo := adminRepositories.AdminRepository{DB: db}
	sesRepo := sesRepositories.SesionsRepository{DB: db}

	userServices := services.Services{Repo: &userRepo, RepoTastes: &tasteRepo}
	tasteServices := tastesServices.TasteServices{Repo: &tasteRepo}
	cantServices := cantServices.CantidadServices{RepoCantidad: &cantRepo}
	pedServices := pedidosServices.PedidosServices{RepoPedidos: &pedRepo, RepoUser: &userRepo, RepoCantidad: &cantRepo, RepoTastes: &tasteRepo}
	admServices := adminServices.AdminServices{RepoAdmin: &admRepo, RepoSesion: &sesRepo}
	sesServices := sesServices.SesionsServices{Repo: &sesRepo}

	userHandler := handlers.UserHandler{Services: &userServices}
	tasteHandler := tastesHandlers.TasteHandler{Services: &tasteServices}
	cantHandler := cantHandlers.CantidadHandler{CantidadService: &cantServices}
	pedHandler := pedidosHandlers.PedidosHandler{Services: &pedServices}
	admHandler := admHandlers.AdminHandler{AdminService: &admServices, SesionService: &sesServices}
	sesHandler := sesHandlers.SesionHandler{SesionService: &sesServices}

	fmt.Println(sesHandler)

	if err := adminRepositories.SeedAdmin(db); err != nil {
		panic("failed to seed admin user: " + err.Error())
	}

	router := gin.New()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("count_validator", func(fl validator.FieldLevel) bool {
			return validators.CountValidator(fl, db)
		})
	}

	router.Use(gin.Recovery())

	user := router.Group("user")

	user.GET(":id", middleware.AuthMiddleware(db), userHandler.GetById)
	user.GET("/all", middleware.AuthMiddleware(db), userHandler.List)
	user.POST("", middleware.AuthMiddleware(db), userHandler.Create)
	user.PUT(":id", middleware.AuthMiddleware(db), userHandler.Update)
	user.DELETE(":id", middleware.AuthMiddleware(db), userHandler.Delete)

	pedido := router.Group("pedido")
	pedido.GET(":user_id", middleware.AuthMiddleware(db), pedHandler.GetCommentByUserId) 
	pedido.GET("/all", middleware.AuthMiddleware(db), pedHandler.List)                  
	pedido.POST("", middleware.AuthMiddleware(db), pedHandler.Create)
	pedido.PUT(":ID", middleware.AuthMiddleware(db), pedHandler.Update) 
	pedido.DELETE(":id", middleware.AuthMiddleware(db), pedHandler.Delete)

	taste := router.Group("taste")
	taste.GET(":id", middleware.AuthMiddleware(db), tasteHandler.GetById)
	taste.POST("", middleware.AuthMiddleware(db), tasteHandler.Create)
	taste.DELETE(":id", middleware.AuthMiddleware(db), tasteHandler.Delete)
	taste.PUT(":id", middleware.AuthMiddleware(db), tasteHandler.Update)
	taste.GET("", middleware.AuthMiddleware(db), tasteHandler.List)

	cantidad := router.Group("cantidad")
	cantidad.GET(":id", middleware.AuthMiddleware(db), cantHandler.GetById)
	cantidad.POST("", middleware.AuthMiddleware(db), cantHandler.Create)
	cantidad.PUT("/:id", middleware.AuthMiddleware(db), cantHandler.Update)
	cantidad.GET("", middleware.AuthMiddleware(db), cantHandler.List)
	cantidad.DELETE(":id", middleware.AuthMiddleware(db), cantHandler.Delete)

	admin := router.Group("admin")
	admin.POST("", middleware.AuthMiddleware(db), admHandler.Create)
	admin.GET(":id", middleware.AuthMiddleware(db), admHandler.GetById)
	admin.PUT(":id", middleware.AuthMiddleware(db), admHandler.Update)
	admin.DELETE(":id", middleware.AuthMiddleware(db), admHandler.Delete)
	admin.POST("/login", admHandler.Login)

	router.Run() // listen and serve on 0.0.0.0:8080

}
