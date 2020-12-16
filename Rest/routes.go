package Rest

import "github.com/gin-gonic/gin"


func RegisterRest(app *gin.Engine){
	v1 := app.Group("/api/v1")
	{
		v1.POST("/register",Register) // Register User
		v1.POST("/login",Login)       // Login User
	}

	contact := app.Group("/api/v1/contact")
	contact.Use(AuthRequired)
	{
		contact.POST("/add",ContactsAdd)
	}
}
