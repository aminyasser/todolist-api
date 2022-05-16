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
	
    // repositories
	userRepo = repository.NewUserRepository(db)
	taskRepo = repository.NewTaskRepository(db)
	// services
	authService = service.NewAuthService(userRepo)
	userService = service.NewUserService(userRepo)
	taskService = service.NewTaskService(taskRepo)
	jwtService = service.NewJWTService()
	// controllers
	auth = controller.NewAuthController(authService , jwtService)
	user = controller.NewUserController(userService , jwtService)
	task = controller.NewTaskController(taskService , jwtService)
)

func main() {
	defer database.CloseDb(db)

    route := gin.Default()
    
	authRoutes := route.Group("/api/auth") 
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
	}
    
	profileRoutes := route.Group("/api" , middleware.JWT(jwtService)) 
	{
		profileRoutes.GET("/profile" , user.Profile )
		profileRoutes.POST("/profile" , user.Update )
	}

    taskRoutes := route.Group("/api" , middleware.JWT(jwtService))
	{
		taskRoutes.GET("/tasks", task.Index)
		taskRoutes.GET("/tasks/:id", task.Show)
		// taskRoutes.POST("/tasks", task.Create)
		// taskRoutes.PATCH("/tasks/:id", task.Update)
		// taskRoutes.DELETE("/tasks/:id", task.Destroy)
	}

	route.Run() // listen and serve on 0.0.0.0:8080 
}