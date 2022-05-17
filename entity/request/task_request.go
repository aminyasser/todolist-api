package request

type Task struct {
	Body    string  `json:"body" bindding:"required"`
	Completed bool  `json:"completed" bindding:"required"`
}