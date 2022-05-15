package response

import "github.com/gin-gonic/gin"


func Error( err string) gin.H {
   return 	gin.H{
      "status" : false, 
	  "error": err,
	}
}

func Success(message string , data interface{}) gin.H {
	return 	gin.H{
	   "status" : true, 
	   "message": message,
	   "data": data,
	 }
 }