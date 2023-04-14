package handlers

import (
	"Employee-Log-System/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) Supervisor(c *gin.Context) {
	var re []models.Department
	h.DB.Find(&re)
	c.IndentedJSON(http.StatusOK, re)
}
