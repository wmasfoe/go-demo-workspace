package main

import (
	"fmt"
	"strconv"

	"example.com/data"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(data.Students)
	r := gin.Default()

	r.GET("/students", func(c *gin.Context) {
		c.JSON(200, data.Students)
	})

	r.GET("/students/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			c.JSON(400, gin.H{"error": "id must be int"})
		}
		// TODO id 转 int64
		c.JSON(200, data.GetStudent(id))
	})

	r.Run(":8080")
}
