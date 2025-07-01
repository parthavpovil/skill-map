package main

import(
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){

	fmt.Println("Skill Map")

	router := gin.Default()
  	router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "skill-map-home-page",
    })
  	})
	router.GET("/login", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "login required",
    })
  })
  router.Run()
}