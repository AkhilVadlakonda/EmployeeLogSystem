package handlers

import (
	"net/http"

	"Employee-Log-System/models"

	"github.com/gin-gonic/gin"
)

func (h handler) Attendance(c *gin.Context) {
	var attendance []models.Attendance

	if err := h.DB.Find(&attendance).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.IndentedJSON(http.StatusOK, attendance)
}

func (h handler) AddAttendance(c *gin.Context) {
	var attendance models.Attendance

	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.DB.Create(&attendance).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "Created Successfully",
		"message": attendance,
	})
}

func (h handler) UpdateAttendance(c *gin.Context) {
	id := c.Param("id")
	var attendance models.Attendance

	if err := h.DB.Where("id = ?", id).First(&attendance).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := c.ShouldBindJSON(&attendance); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := h.DB.Save(&attendance).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Updated Successfully",
		"message": attendance,
	})
}

func (h handler) DeleteAttendance(c *gin.Context) {
	id := c.Param("id")
	var attendance models.Attendance

	if err := h.DB.Where("id = ?", id).First(&attendance).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := h.DB.Delete(&attendance).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Deleted Successfully",
	})
}
