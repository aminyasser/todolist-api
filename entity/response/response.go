package response


type successResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type errorResponse struct {
	Status  bool        `json:"status"`
	Error  string       `json:"error"`
}
type messageResponse struct {
	Status  bool        `json:"status"`
	Message  string       `json:"message"`
}


func Error( err string) errorResponse {
   return errorResponse{
      Status: false,
	  Error: err,
   }
}

func Success(message string , data interface{}) successResponse {
	return 	successResponse{
		Status: true,
		Message: message,
		Data: data,
	}
 }
 func Message(message string) messageResponse {
	return 	messageResponse{
		Status: true,
		Message: message,
	}
 }