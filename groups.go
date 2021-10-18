package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main(){
  r := gin.Default()

  v1 := r.Group("/api")
  {
    v1.GET("/sign-up", signUp)
    v1.GET("/sign-in", signIn)
  }
  v2 := r.Group("/info")
  {
    v2.GET("/read", read)
    v2.GET("/submit", submit)
  }
  r.Run(":8080")
}

func signIn(c *gin.Context){
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Регистрация"})
}

func signUp(c *gin.Context){
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Вход"})
}

func read(c *gin.Context){
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Чтение"})
}

func submit(c *gin.Context){
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Submit"})
}
