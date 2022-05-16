package service

import (
	"github.com/aminyasser/todo-list/entity/response"
	"github.com/aminyasser/todo-list/repository"
)

type TaskService interface {
	GetAll(string) (*[]response.Task , error)
	Get(id string) (*response.Task , error) 
}
type taskService struct {
	repository repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *taskService {
	return &taskService{repo}
}


func (task *taskService) GetAll(id string) (*[]response.Task , error)  {
    tasks, err := task.repository.All(id)
	if err != nil {
		return nil, err
	}

	res := response.NewTaskArrayResponse(tasks)
	return &res, nil
}

func (task *taskService) Get(id string) (*response.Task , error)  {
    r, err := task.repository.FindBy("id" , id)
	if err != nil {
		return nil, err
	}

	res := response.NewTaskResponse(r)
	return &res, nil
}