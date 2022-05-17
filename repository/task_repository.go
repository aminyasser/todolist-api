package repository

import (
	"github.com/aminyasser/todo-list/entity/model"
	"gorm.io/gorm"
)


type TaskRepository interface {
	All(id string) ([]model.Task , error)
	FindBy( string ,  string) (model.Task, error)
	Insert( model.Task) (model.Task, error)
	Update( model.Task) model.Task
	Delete( string) error
}
type taskRepository struct {
	connection  *gorm.DB
} 

func NewTaskRepository(db *gorm.DB) *taskRepository {
   return &taskRepository{
	    connection: db,
   }
}

func (task *taskRepository) All(id string) ([]model.Task , error){
	tasks := []model.Task{}
	task.connection.Preload("User").Where("user_id = ?", id).Find(&tasks)
	return tasks, nil
}

func (task *taskRepository) FindBy(colm string , value string) (model.Task, error) {
	var taskModel model.Task
	res := task.connection.Preload("User").Where(colm +" = ?", value).Take(&taskModel)
	if res.Error != nil {
		return taskModel, res.Error
	}
	return taskModel, nil
}

func (task *taskRepository) Insert(taskModel model.Task) (model.Task, error) {
	task.connection.Save(&taskModel)
	task.connection.Preload("User").Find(&taskModel)
	return taskModel, nil
}

func (task *taskRepository) Update(taskModel model.Task) model.Task {
	task.connection.Where("id = ?", taskModel.ID).Updates(&taskModel)
	task.connection.Preload("User").Find(&taskModel)
	return taskModel
}

func (task *taskRepository) Delete(id string) error {
	res := task.connection.Delete(&model.Task{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}