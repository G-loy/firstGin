package databases

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open("mysql", "root:Gg13506800412@tcp(121.41.8.129:3306)/zmz")
	if err != nil {
		log.Println("err:", err.Error())
	} else {
		log.Println("connect to database")
	}
}
