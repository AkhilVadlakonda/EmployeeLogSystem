package models

type FrontEmployee struct {
	Id         int    `json:"id" gorm:"primaryKey;autoIncrement:true;unique"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Gender     string `json:"gender"`
	Status     int    `json:"status"`
}
