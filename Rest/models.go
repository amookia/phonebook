package Rest

type User struct {
	ID 		   uint   `json:"id" gorm:"primaryKey"`
	Email      string `json:"email" gorm:"unique" binding:"required,min=6,max=35"`
	Password   string `json:"password" binding:"required,min=8,max=35"`
}

type Contact struct {
	ID 		         uint   `json:"id" gorm:"primaryKey"`
	Name 	         string `json:"name" gorm:"not null" binding:"required,min=6,max=35"`
	Phone_number     string `json:"phone_number" gorm:"not null"`
	Description 	 string `json:"description"`
	User_Id 		 int    `json:"user_id"`
}


type AddContactForm struct{
	ID 				 uint 			`json:"id"`
	Name             string	        `json:"name" binding:"required,min=6,max=35"`
	Phone_numbers    []string       `json:"phone_numbers" binding:"required"`
	Description 	 string         `json:"description"`
}