package main

import (
	"Employee-Log-System/controllers"
	"Employee-Log-System/db"

	"fmt"
)

var Port string = ":8082"

func main() {
	DB := db.Init()
	fmt.Println("Listing to port : ", Port)
	controllers.Controller(Port, DB)
}
