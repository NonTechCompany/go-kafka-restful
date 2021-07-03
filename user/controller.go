package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strconv"
)

func Router(r *gin.RouterGroup) {
	r.POST("/", addUser())
	r.GET("/:id", getUserById())
}

func getUserById() func(c *gin.Context) {
	return func(c *gin.Context) {

		logrus.WithFields(logrus.Fields{"hostname": "staging-1"}).Info("ggwp")
		id := c.Param("id")
		value, _ := strconv.Atoi(id)
		user, err := getUserFromId(value)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(200, user)
	}
}

func addUser() func(c *gin.Context) {
	return func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		var userDTO UserDTO
		_ = json.Unmarshal(body, &userDTO)
		c.JSON(200, gin.H{
			"id": insertUser(userDTO),
		})
	}
}
