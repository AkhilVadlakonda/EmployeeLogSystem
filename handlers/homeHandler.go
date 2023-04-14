package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h handler) Home(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"App":     "Employee-Log-System",
		"Version": "1.0",
		"Author":  "EmployeeLogSystem",
		"Url":     "https://github.com/EmployeeLogSystem/EmployeeLogSystem",
	})

}
