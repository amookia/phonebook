package Rest

type User struct {
	ID 		   uint   `json:"id" gorm:"primaryKey"`
	Email      string `json:"email" gorm:"unique" binding:"required,min=6,max=35"`
	Password   string `json:"password" binding:"required,min=8,max=35"`
}
