package Config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 2020/12/16 12:34 AM

var DB *gorm.DB

var DBconfig struct{
	Host 		 string
	User         string
	Password     string
	Database 	 string
}


func ConnectToDB(){
	//Configs
	DBconfig.Host = "mysql:3306" //DB Host
	DBconfig.User = "root"			 //DB USERNAME
	DBconfig.Password = "password"   //DB PASSWORD
	DBconfig.Database = "taskapi"	 //DB NAME

	//URI String format
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DBconfig.User,
		DBconfig.Password,
		DBconfig.Host,
		DBconfig.Database)
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("FAILED TO CONECT DATABASE!")}



	fmt.Println("Migrations DONE!")

	DB = db
}

