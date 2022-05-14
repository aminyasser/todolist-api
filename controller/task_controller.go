package controller

import (
	"net/http"

	"github.com/aminyasser/todo-list/entity/model"
	"github.com/aminyasser/todo-list/entity/response"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)


type TaskController interface {
	GetAll(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type taskController struct {
	connection  *gorm.DB
} 

func NewTaskController(db *gorm.DB) *taskController {
   return &taskController{
	    connection: db,
   }
}

func (t *taskController) GetAll(c *gin.Context) {
    var tasks []model.Task
    result := t.connection.Find(&tasks)
    
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"message": result.Error.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Success",
			"tasks" : tasks,
		})
    }
}
func (t *taskController) Get(c *gin.Context) {
	id := c.Param("id")

	var task response.Task

	result := t.connection.Model(&model.Task{}).First(&task, id)

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"message": result.Error.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Success",
			"task": task,
		})
	}

}
func (t *taskController) Create(c *gin.Context) {
	var task response.Task
    err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	} else {
       t.connection.Create(&model.Task{Body: task.Body , Completed: task.Completed })

		c.JSON(200, gin.H{
			"message": "Task created successfully",
			"task": task,
	    })
	}
   
}
func (t *taskController) Update(c *gin.Context) {
	id := c.Param("id")
	var task response.Task
    err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
	} else {
       t.connection.Model(&model.Task{}).Where("id = ?", id).Updates(model.Task{Body: task.Body , Completed: task.Completed })
		c.JSON(200, gin.H{
			"message": "Task updated successfully",
			"task": task,
		})
	}
	
}
func (t *taskController) Delete(c *gin.Context) {

	id := c.Param("id")
	result :=t.connection.Delete(&model.Task{}, id)
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"message": "the task does not exist",
		})
	} else {
	c.JSON(200, gin.H{
		"message": "Task deleted succesfully",
	})
    } 
}