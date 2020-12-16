package Rest

import (
	"fmt"
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
			token := Config.GenrateJWT(user.Email)
			c.JSON(200,gin.H{
				"status":"Registered",
				"token":token,
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
			token := Config.GenrateJWT(user.Email)
			c.JSON(200,gin.H{
				"status": "Logged in",
				"token" : token,
			})
		}else {
			c.JSON(401,gin.H{
				"error":"Email or Password is wrong!",
			})
		}
	}
}

func AuthRequired(c *gin.Context){
	token := c.GetHeader("Authorization")
	if len(token) != 0 {
		verify := Config.CheckJWT(token)
		if verify {
			c.Next()
		}else {
			c.AbortWithStatusJSON(401,gin.H{
				"error":"Invalid Token!",
			})
		}
	}else {
		c.AbortWithStatusJSON(401,gin.H{
			"error":"Token Required!",
		})
	}

}

func ContactsAdd(c *gin.Context){
	var ContactForm AddContactForm
	var contact 	Contact
	token := c.GetHeader("Authorization")
	email := Config.ClaimJWT(token)
	fmt.Println(email)
	form := c.ShouldBindJSON(&ContactForm)
	lenofnums := len(ContactForm.Phone_numbers)
	if form != nil || lenofnums == 0 {
		c.JSON(400,gin.H{
			"error":"Check your input!",
		})
	}else{
		var str string
		for _,nums := range ContactForm.Phone_numbers{
			fmt.Println(nums)
			str += nums + ","
		}
		Config.DB.Table("users").Select("id as user_id").Where("email = ?",email).Scan(&contact)
		contact.Name = ContactForm.Name
		contact.Phone_number = str
		contact.Description = ContactForm.Description
		res := Config.DB.Table("contacts").Create(&contact).Error
		if res != nil {
			c.JSON(500,gin.H{
				"error":"IDK",
			})
			return
		}
		c.JSON(200,gin.H{"status":"Success"})
	}
}
