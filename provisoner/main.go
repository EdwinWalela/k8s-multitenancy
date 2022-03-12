package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func deployCluster(tenantName string) error {
	log.Printf("Deploying application to /%s", tenantName)

	res := exec.Command("/bin/sh", "./apply.sh", tenantName, "/"+tenantName)

	if err := res.Start(); err != nil {
		return err
	}

	return nil
}

func main() {

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		tenantName := c.PostForm("tenant")

		err := deployCluster(tenantName)

		if err != nil {
			c.JSON(501, gin.H{
				"msg":   "failed",
				"error": err.Error(),
			})
		} else {
			c.JSON(201, gin.H{
				"msg": "Application deployed",
				"url": "https://crafted.co.ke/" + tenantName,
			})
		}

	})

	r.Run(fmt.Sprintf(":%d", 7000))

}
