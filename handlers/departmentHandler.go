package handlers

import (
	"net/http"

	"Employee-Log-System/models"

	"github.com/gin-gonic/gin"
)

func (h handler) AllDepartment(c *gin.Context) {
	var department []models.Department
	h.DB.Find(&department)
	c.IndentedJSON(http.StatusOK, department)
}

func (h handler) AddDepartment(c *gin.Context) {
	var department models.Department
	err := c.BindJSON(&department)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	h.DB.Create(&department)
	c.IndentedJSON(http.StatusOK, &department)
}

func (h handler) UpdateDepartment(c *gin.Context) {
	id := c.Param("id")
	var department models.Department

	err := h.DB.Where("id = ?", id).First(&department).Error

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.BindJSON(&department)
	h.DB.Save(&department)
	c.JSON(http.StatusOK, &department)
}

func (h handler) DeleteDepartment(c *gin.Context) {
	id := c.Param("id")
	var department models.Department

	err := h.DB.Where("id = ?", id).First(&department).Error

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	h.DB.Delete(&department)
	c.JSON(http.StatusOK, gin.H{
		"status": "Deleted Successfully",
	})
}
