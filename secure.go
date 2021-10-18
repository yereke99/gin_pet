package main

import (
  "github.com/gin-gonic/gin"
)

func main(){
  r := gin.Default()

  r.GET("/", func(c *gin.Context){
    names := []string{"lena", "paul", "jake"}
    c.SecureJSON(200, names)
  })

  r.Run(":8080")
}
