package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func deployCluster(ns string, port string) error {
	log.Printf("Deploying resources to [%s] namespace", ns)

	res := exec.Command("/bin/sh", "./apply.sh", ns, port)

	if err := res.Start(); err != nil {
		return err
	}

	return nil
}

func main() {

	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		ns := c.PostForm("ns")
		port := c.PostForm("port")

		err := deployCluster(ns, port)

		if err != nil {
			c.JSON(501, gin.H{
				"msg":   "failed",
				"error": err.Error(),
			})
		} else {
			c.JSON(201, gin.H{
				"msg":  "cluster deployed",
				"port": port,
			})
		}

	})

	r.Run(fmt.Sprintf(":%d", 7000))

}
