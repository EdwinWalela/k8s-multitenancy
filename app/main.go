package main

import "github.com/gin-gonic/gin"

func main() {
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

	r.Run(":3000")
}
