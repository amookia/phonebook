package Rest

import (
	"github.com/gin-gonic/gin"
	"phonebook/Config"
)


func Register(c *gin.Context){
	var user User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400,gin.H{"error":"Check your input!"})
		return

	}else{
		hashedpw := Config.HashGen(user.Password) // Genereate hash
		user.Password = hashedpw
		dberr := Config.DB.Table("users").Create(&user).Error
		if dberr != nil{
			c.JSON(400,gin.H{"error":"Email exists!"})
			return
		}else {
			//Generate jwt
			c.JSON(200,gin.H{
				"status":"Registered",
				"token":"salamazizamtokenmoken",
			})
			return
		}

	}
}


func Login(c *gin.Context){
	var user User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(400,gin.H{"error":"Check your input!"})
		return

	}else{
		password := user.Password
		Config.DB.Table("users").
			Where("email = ?",user.Email).Scan(&user)
		compare := Config.HashCompare(password,user.Password)
		if compare == true {
			//Generate jwt
			c.JSON(200,gin.H{
				"status": "Logged in",
				"token" : "toazizedelami",
			})
		}else {
			c.JSON(401,gin.H{
				"error":"Email or Password is wrong!",
			})
		}
	}
}
