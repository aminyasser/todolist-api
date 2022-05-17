package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aminyasser/todo-list/entity/request"
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
	result, err := t.taskService.GetAll(id)
	if err != nil {
		response := response.Error(err.Error())
		c.JSON(401, response)
		return
	}
	res := response.Success("Success", result)
	c.JSON(200, res)
}

func (t *taskController) Show(c *gin.Context) {
	id := c.Param("id")

	result, err := t.taskService.Get(id)
	if err != nil {
		response := response.Error(err.Error())
		c.JSON(401, response)
		return
	}

	res := response.Success("Success", result)
	c.JSON(200, res)
}

func (t *taskController) Create(c *gin.Context) {
	var task request.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		id := t.getIdFromHeader(c)
		res , _ := t.taskService.CreateTask(task , id)

		response := response.Success("task created successfully", res)
		c.JSON(200, response)
	}
}

func (t *taskController) Update(c *gin.Context) {
	id := c.Param("id")
	var task request.UpdateTask
    err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	} else {
		taskId , _ := strconv.Atoi(id)
        task.ID = uint(taskId)
        
		userId := t.getIdFromHeader(c)
		res , err := t.taskService.UpdateTask(task , userId)
		if err != nil {
			response := response.Error(err.Error())
			c.JSON(404 , response)
			return
		}

		response := response.Success("task updated successfully", res)
		c.JSON(200, response)
	}

}

func (t *taskController) Destroy(c *gin.Context) {

	id := c.Param("id")
	userId := t.getIdFromHeader(c)


	err  := t.taskService.DeleteTask(id , userId)
	if err != nil {
		response := response.Error(err.Error())
		c.JSON(404 , response)
		return
	}
	
	response := response.Message("task deleted successfully")
	c.JSON(200, response)
}

func (c *taskController) getIdFromHeader(ctx *gin.Context) string {
	header := ctx.GetHeader("Authorization")
	token := c.jwtService.ValidateToken(header, ctx)
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
