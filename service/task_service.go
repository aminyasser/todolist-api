package service

import (
	"errors"
	"strconv"

	"github.com/aminyasser/todo-list/entity/model"
	"github.com/aminyasser/todo-list/entity/request"
	"github.com/aminyasser/todo-list/entity/response"
	"github.com/aminyasser/todo-list/repository"
	"github.com/mashingan/smapping"
)

type TaskService interface {
	GetAll(string) (*[]response.Task, error)
	Get(string) (*response.Task, error)
	CreateTask( request.Task ,  string) (*response.Task, error) 
	UpdateTask( request.UpdateTask ,  string) (*response.Task, error)
	DeleteTask( string ,  string) error
}
type taskService struct {
	repository repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *taskService {
	return &taskService{repo}
}

func (task *taskService) GetAll(id string) (*[]response.Task, error) {
	tasks, err := task.repository.All(id)
	if err != nil {
		return nil, err
	}

	res := response.NewTaskArrayResponse(tasks)
	return &res, nil
}

func (task *taskService) Get(id string) (*response.Task, error) {
	r, err := task.repository.FindBy("id", id)
	if err != nil {
		return nil, err
	}

	res := response.NewTaskResponse(r)
	return &res, nil
}

func (task *taskService) CreateTask(taskReq request.Task , userId string) (*response.Task, error) {
	taskModel := model.Task{}
	err := smapping.FillStruct(&taskModel, smapping.MapFields(&taskReq))
	if err != nil {
		return nil, err
	}
   
    taskModel.UserId , _ = strconv.Atoi(userId)

	createdTask , _ := task.repository.Insert(taskModel)

	res := response.NewTaskResponse(createdTask)
	return &res, nil
}

func (task *taskService) UpdateTask(taskReq request.UpdateTask , userId string) (*response.Task, error) {
	
	taskCheck , _ := task.repository.FindBy("id" , strconv.Itoa(int(taskReq.ID)))
	if strconv.Itoa(taskCheck.UserId) != userId {
		return nil, errors.New("you can't update someone else task")
	}

	taskModel := model.Task{}
	err := smapping.FillStruct(&taskModel, smapping.MapFields(&taskReq))
	taskModel.ID = taskReq.ID
	if err != nil {
		return nil, err
	}
	
    taskModel.UserId , _ = strconv.Atoi(userId)

	updatedTask := task.repository.Update(taskModel)

	res := response.NewTaskResponse(updatedTask)
	return &res, nil
}

func (task *taskService) DeleteTask(id string , userId string) error {
	taskCheck , _ := task.repository.FindBy("id" , id)
	if strconv.Itoa(taskCheck.UserId) != userId {
		return errors.New("you can't delete someone else task")
	}

	err := task.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

