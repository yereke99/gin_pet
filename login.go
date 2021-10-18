package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

type Login struct{
  User string `form:"user" json:"user" xml:"user" binding:"required"`
  Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func loginJSON(c *gin.Context){
  var json Login

  if err := c.ShouldBindJSON(&json); err != nil{
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
  }
  if json.User != "yerek" || json.Password != "123456" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
  }
  c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func loginXML(c *gin.Context){
  var xml Login

  if err := c.ShouldBindXML(&xml); err != nil{
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
  }
  if xml.User != "yerek" || xml.Password != "123456" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
  }
  c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func loginForm(c *gin.Context){
  var form Login

  if err := c.ShouldBindXML(&form); err != nil{
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
  }
  if form.User != "yerek" || form.Password != "123456" {
      c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
  }
  c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}



func main(){
  r := gin.Default()

  r.POST("/loginJSON", loginJSON)
  r.POST("/loginXML", loginXML)
  r.POST("/loginForm", loginForm)

  r.Run(":8080")

}
