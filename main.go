package main 

import (
	"github.com/gin-gonic/gin"
	"github.com/aminyasser/todo-list/controller"
	"github.com/aminyasser/todo-list/database"

)

var (
	db , _ = database.InitDb()
	task = controller.NewTaskController(db)
)

func main() {
	defer database.CloseDb(db)

    route := gin.Default()

    r := route.Group("/api")
	{
		r.GET("/tasks", task.GetAll)
		r.GET("/tasks/:id", task.Get)
		r.POST("/tasks", task.Create)
		r.PATCH("/tasks/:id", task.Update)
		r.DELETE("/tasks/:id", task.Delete)
	}

	route.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}