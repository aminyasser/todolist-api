package main

import (
	"github.com/aminyasser/todo-list/controller"
	"github.com/aminyasser/todo-list/database"
	"github.com/aminyasser/todo-list/repository"
	"github.com/aminyasser/todo-list/service"
	"github.com/aminyasser/todo-list/middleware"
	"github.com/gin-gonic/gin"
)

var (
	db , _ = database.InitDb()
	task = controller.NewTaskController(db)
    // repositories
	userRepo = repository.NewUserRepository(db)
	// services
	authService = service.NewAuthService(userRepo)
	userService = service.NewUserService(userRepo)
	jwtService = service.NewJWTService()

	auth = controller.NewAuthController(authService , jwtService)
	user = controller.NewUserController(userService , jwtService)

)

func main() {
	defer database.CloseDb(db)

    route := gin.Default()
    
	authRoutes := route.Group("/api/auth") 
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
	}
    
	profileRoutes := route.Group("/api" , middleware.AuthorizeJWT(jwtService)) 
	{
		profileRoutes.GET("/profile" , user.Profile )
		profileRoutes.POST("/profile" , user.Update )
	}

    taskRoutes := route.Group("/api")
	{
		taskRoutes.GET("/tasks", task.GetAll)
		taskRoutes.GET("/tasks/:id", task.Get)
		taskRoutes.POST("/tasks", task.Create)
		taskRoutes.PATCH("/tasks/:id", task.Update)
		taskRoutes.DELETE("/tasks/:id", task.Delete)
	}

	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}