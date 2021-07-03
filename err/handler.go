package err

import (
	"fmt"
	"github.com/NonTechCompany/go-kafka-restful/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.Last()
		fmt.Println(detectedErrors.Err)
		if detectedErrors.Err == user.UserNotFoundError {
			c.JSON(http.StatusOK, gin.H{
				"code": "1",
				"msg":  "User not found",
				"data": nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": "1",
				"msg":  detectedErrors,
				"data": nil,
			})
		}
		c.Abort()
		return

	}
}
