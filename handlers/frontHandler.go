package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"Employee-Log-System/models"

	"github.com/gin-gonic/gin"
)

func (h handler) FrontHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.gohtml", nil)
}

func (h handler) FrontSupervisor(c *gin.Context) {
	req, err := http.Get("http://localhost:8082/api/user")
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	var supervisor []models.Supervisor
	json.Unmarshal(body, &supervisor)

	c.HTML(http.StatusOK, "supervisor.gohtml", supervisor)
}

func (h handler) FrontEmployee(c *gin.Context) {
	req, err := http.Get("http://localhost:8082/api/employee")
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	var employee []models.FrontEmployee
	json.Unmarshal(body, &employee)

	c.HTML(http.StatusOK, "employee.gohtml", employee)
}

func (h handler) FrontAttendance(c *gin.Context) {
	req, err := http.Get("http://localhost:8082/api/attendance")
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	var attendance []models.Attendance
	json.Unmarshal(body, &attendance)
	fmt.Println("Attendance", req)

	c.HTML(http.StatusOK, "attendance.gohtml", attendance)
}

func (h handler) Login(c *gin.Context) {
	// Parse the username and password from the request body
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	// Use the username field from the request

	loginRequest.Username = user.UserName
	loginRequest.Password = user.Password
	fmt.Println("Username" + loginRequest.Username + "Password" + loginRequest.Password)

	// Find the user with the given username and password
	if err := h.DB.Where("user_name = ? AND password = ?", loginRequest.Username, loginRequest.Password).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate a random session ID
	sessionID := strconv.FormatInt(rand.Int63(), 16)

	// Store the session ID in a secure HTTP-only cookie
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})

	// Store the session ID in the database for future verification
	h.DB.Model(&user).Update("session_id", sessionID)
	fmt.Println("sessionID" + sessionID)

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func (h handler) Logout(c *gin.Context) {
	// Get the session ID from the cookie
	cookie, err := c.Request.Cookie("session_id")
	fmt.Println("Session ID1:", cookie.Value)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Session ID not found"})
		return
	}

	// Delete the session ID from the database
	var user models.User
	if err := h.DB.Where("session_id = ?", cookie.Value).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}
	h.DB.Model(&user).Update("session_id", "")

	// Delete the session ID cookie
	cookie.MaxAge = -1
	cookie.Path = "/"
	http.SetCookie(c.Writer, cookie)

	// Return a success message
	c.Redirect(http.StatusFound, "/")

}
func (h handler) EditProfile(c *gin.Context) {
	var user models.User

	c.HTML(http.StatusOK, "editprofile.gohtml", gin.H{
		"user": user,
	})
}
func (h handler) EditAttendance(c *gin.Context) {
	var attendance models.Attendance

	c.HTML(http.StatusOK, "editattendance.gohtml", gin.H{
		"attendance": attendance,
	})
}

func (h handler) Profile(c *gin.Context) {
	// Get the session ID from the cookie
	cookie, err := c.Request.Cookie("session_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	sessionID := cookie.Value

	// Find the user with the given session ID
	var user models.User
	if err := h.DB.Where("session_id = ?", sessionID).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	c.HTML(http.StatusOK, "profile.gohtml", gin.H{
		"user": user,
	})
}
