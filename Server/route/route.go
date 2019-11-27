package route

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func GetAction(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + " is " + action
	c.String(http.StatusOK, message)
}

func GetWelcome(c *gin.Context) {
	// /welcome?firstname=Jane&lastname=Doe
	firstname := c.DefaultQuery("firstname", "Guest") // If not exists, return Guest in this case
	lastname := c.Query("lastname")                   // if not exists, return "" empty string
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

type Login struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func PostEcho(c *gin.Context) {
	// {
	// 	"User": "naoto",
	// 	"Password": "passwordDesu"
	// }
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(json)
	c.JSON(http.StatusOK, gin.H{
		"User":     json.User,
		"Password": json.Password,
	})
}
