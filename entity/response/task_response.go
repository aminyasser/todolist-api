package response




type Task struct {
	Body string  `json:"body" binding:"required"`
	Completed bool  `json:"completed" binding:"required"`
 } 