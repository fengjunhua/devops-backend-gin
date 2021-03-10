package mysql

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)



type User struct {
	ID         uint     `gorm:"id"         json:"id"`
	Username   string   `gorm:"username"   json:"username"`
	Firstname  string   `gorm:"firstname"  json:"firstname"`
	Lastname   string   `gorm:"lastname"   json:"lastname"`
	Telephone  uint     `gorm:"telephone"  json:"telephone"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
}

//var db *gorm.DB

func init() {



	p1 := User{ID:1,Username: "蔡安宁",Firstname: "蔡",Lastname: "安宁",Telephone: 13520859533}
	p2 := User{ID:2,Username: "冯君华",Firstname: "冯",Lastname: "君华",Telephone: 13520859533}

	fmt.Println(p1)
	fmt.Println(p2)

}

func GetUserById(c *gin.Context) {
	db, err := gorm.Open("mysql", "root:120225fjhFJH!@tcp(182.92.236.162:3306)/kubernetes?charset=utf8")
	defer db.Close()
	if err != nil{
		panic(err)
	}
	//userid := c.Param("userid")
	//result := map[string]interface{}{}
	result := []User{}
	//var result User
	//db.Table("users").Where("username = ?", "蔡安宁").First(&result)
	db.Table("users").Where("id = ?", 2).Find(&result)
	fmt.Println(result)
	c.JSON(http.StatusOK,result)

}