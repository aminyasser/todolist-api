package request

type Task struct {
	Body    string  `json:"body" bindding:"required"`
	Completed bool  `json:"completed" bindding:"required"`
}

type UpdateTask struct {
	ID      uint   `json:"id"`
	Body    string  `json:"body" bindding:"required"`
	Completed bool  `json:"completed" bindding:"required"`
}