package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string
	Name  string
}

func main() {

	dbUri := os.Getenv("DB_URI")
	port := os.Getenv("PORT")

	db, err := gorm.Open(mysql.Open(dbUri), &gorm.Config{})

	if err != nil {
		panic("failed to connect to db")
	}

	db.AutoMigrate(&User{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "hello",
		})
	})

	r.POST("/", func(c *gin.Context) {
		email := c.PostForm("email")
		name := c.PostForm("name")

		c.JSON(201, gin.H{
			"msg":   "created",
			"email": email,
			"name":  name,
		})
	})

	r.Run(fmt.Sprintf(":%s", port))
}
