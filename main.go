package main

import (
	"github.com/aminyasser/todo-list/controller"
	"github.com/aminyasser/todo-list/database"
	"github.com/aminyasser/todo-list/repository"
	"github.com/aminyasser/todo-list/service"
	"github.com/gin-gonic/gin"
)

var (
	db , _ = database.InitDb()
	task = controller.NewTaskController(db)

	authRepo = repository.NewUserRepository(db)
	authService = service.NewAuthService(authRepo)
	auth = controller.NewAuthController(authService)
)

func main() {
	defer database.CloseDb(db)

    route := gin.Default()
    
	authRoutes := route.Group("/api/auth") 
	{
		authRoutes.POST("/login", auth.Login)
		authRoutes.POST("/register", auth.Register)
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