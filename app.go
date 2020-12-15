package main

import "github.com/gin-gonic/gin"
import "phonebook/Config"
import "phonebook/Rest"

func main(){
	r := gin.Default()

	//INIT DB
	Config.ConnectToDB()

	//Register
	Rest.RegisterRest(r)

	//Migrations
	Config.DB.AutoMigrate(&Rest.User{}) // Migrate User Table

	r.Run(":8000")
}