package controllers

import (
	"Employee-Log-System/handlers"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func Controller(Port string, DB *gorm.DB) {
	h := handlers.New(DB)
	r := gin.Default()
	r.GET("/api", h.Home)

	r.Use(static.Serve("/asset", static.LocalFile("views/frontend/assets", true)))
	r.LoadHTMLGlob("views/frontend/*.gohtml")

	department := r.Group("/api")
	{
		department.GET("/department", h.AllDepartment)
		department.POST("/department", h.AddDepartment)
		department.PUT("/department/:id", h.UpdateDepartment)
		department.DELETE("/department/:id", h.DeleteDepartment)
	}
	employee := r.Group("/api")
	{
		employee.GET("/employee", h.AllEmployee)
		employee.POST("/employee", h.AddEmployee)
		employee.PUT("/employee/:id", h.UpdateEmployee)
		employee.DELETE("/employee/:id", h.DeleteEmployee)
	}

	user := r.Group("/api")
	{
		user.GET("/user", h.AllUser)
		user.POST("/user", h.AddUser)

		user.PUT("/user/:id", h.UpdateUser)
		user.DELETE("/user/:id", h.DeleteUser)
	}

	attendance := r.Group("/api")
	{
		attendance.GET("/attendance", h.Attendance)
		attendance.POST("/attendance", h.AddAttendance)
		attendance.PUT("/attendance/:id", h.UpdateAttendance)
		attendance.DELETE("/attendance/:id", h.DeleteAttendance)
	}

	// frontend
	frontend := r.Group("/")
	{
		frontend.GET("/", h.FrontHome)
		frontend.GET("/supervisor", h.FrontSupervisor)
		frontend.GET("/employee", h.FrontEmployee)
		frontend.POST("/login", h.Login)
		frontend.POST("/logout", h.Logout)
		frontend.GET("/profile", h.Profile)
		frontend.GET("/editprofile", h.EditProfile)
		frontend.GET("/editattendance", h.EditAttendance)

		frontend.GET("/attendance", h.FrontAttendance)

	}

	r.Run(Port)
}
