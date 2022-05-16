package controller

import (
	"fmt"


	"github.com/aminyasser/todo-list/entity/response"
	"github.com/aminyasser/todo-list/service"
	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)


type TaskController interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

type taskController struct {
	taskService service.TaskService
	jwtService  service.JWTService
} 

func NewTaskController(taskService service.TaskService, jwtService service.JWTService) *taskController {
   return &taskController{taskService, jwtService}
}


func (t *taskController) Index(c *gin.Context) {
    
    id := t.getIdFromHeader(c)
	result , err := t.taskService.GetAll(id)
	if err != nil {
		response := response.Error(err.Error())
		c.JSON( 401 , response)
		return
	}
    res := response.Success("Success" , result)
	c.JSON(200, res)
}


func (t *taskController) Show(c *gin.Context) {
	id := c.Param("id")
	
	result , err := t.taskService.Get(id)
	if err != nil {
		response := response.Error(err.Error())
		c.JSON( 401 , response)
		return
	}

	res := response.Success("Success" , result)
	c.JSON(200, res)
}


// func (t *taskController) Create(c *gin.Context) {
// 	var task response.Task
//     err := c.ShouldBindJSON(&task)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
// 	} else {
//        t.connection.Create(&model.Task{Body: task.Body , Completed: task.Completed })

// 		c.JSON(200, gin.H{
// 			"message": "Task created successfully",
// 			"task": task,
// 	    })
// 	}
   
// }
// func (t *taskController) Update(c *gin.Context) {
// 	id := c.Param("id")
// 	var task response.Task
//     err := c.ShouldBindJSON(&task)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
// 	} else {
//        t.connection.Model(&model.Task{}).Where("id = ?", id).Updates(model.Task{Body: task.Body , Completed: task.Completed })
// 		c.JSON(200, gin.H{
// 			"message": "Task updated successfully",
// 			"task": task,
// 		})
// 	}
	
// }
// func (t *taskController) Destroy(c *gin.Context) {

// 	id := c.Param("id")
// 	result :=t.connection.Delete(&model.Task{}, id)
// 	if result.RowsAffected == 0 {
// 		c.JSON(404, gin.H{
// 			"message": "the task does not exist",
// 		})
// 	} else {
// 	c.JSON(200, gin.H{
// 		"message": "Task deleted succesfully",
// 	})
//     } 
// }

func (c *taskController) getIdFromHeader(ctx *gin.Context) string {
	header := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(header, ctx)
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}