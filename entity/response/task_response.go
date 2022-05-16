package response

import "github.com/aminyasser/todo-list/entity/model"




type Task struct {
	ID uint   `json:"id" binding:"required"`
	Body string  `json:"body" binding:"required"`
	Completed bool  `json:"completed" binding:"required"`
	User ProfileResponse `json:"user" binding:"required"`
} 

 func NewTaskArrayResponse(tasks []model.Task) []Task {
	  taskRes :=  []Task{}

     for _ , val := range tasks {
           tmp :=  Task{
			ID:    val.ID,
			Body:  val.Body,
			Completed: val.Completed,
			User: NewProfileResponse(val.User),
		}

		taskRes = append(taskRes , tmp)
	 } 
	 
	return taskRes
}

func NewTaskResponse(task model.Task) Task {
	return Task{
	ID:    task.ID,
	Body:  task.Body,
	Completed: task.Completed,
	User: NewProfileResponse(task.User),
	}
}

