package handlers

import (
	"net/http"

	"Employee-Log-System/models"

	"github.com/gin-gonic/gin"
)

func (h handler) AllEmployee(c *gin.Context) {
	var re []models.Employee
	h.DB.Find(&re)
	c.IndentedJSON(http.StatusOK, re)
}

func (h handler) AddEmployee(c *gin.Context) {
	var employee models.Employee
	err := c.BindJSON(&employee)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	h.DB.Create(&employee)
	c.JSON(http.StatusOK, &employee)
}

func (h handler) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee
	if err := h.DB.Where("id = ?", id).First(&employee).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&employee)
	h.DB.Save(&employee)
	c.JSON(http.StatusOK, &employee)
}

func (h handler) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	var employee models.Employee

	err := h.DB.Where("id = ?", id).First(&employee).Error

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	h.DB.Delete(&employee)
	c.JSON(http.StatusOK, gin.H{
		"status": "Deleted Successfully",
	})
}
