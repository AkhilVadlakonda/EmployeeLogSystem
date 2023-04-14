package models

type User struct {
	Id          int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name        string `json:"name"`
	UserName    string `json:"username"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	DateOfBirth string `json:"dob"`
	Designation string `json:"designation"`
	Password    string `json:"password"`
	UserType    string `json:"type"`
	Status      int    `json:"status"`
	SessionID   string
}

type Employee struct {
	Id         int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Gender     string `json:"gender"`
	Status     int    `json:"status"`
}

type Department struct {
	Id         int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Department string `json:"department"`
	Supervisor int    `json:"supervisor" ` //gorm:"foreignKey:Id"
}

type Attendance struct {
	Id                 int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name               string `json:"name"` //`json:"name" gorm:"foreignKey:Name"`
	EmployeeAttendance int    `json:"employeeAttendance"`
}
