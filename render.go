package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main(){
  r := gin.Default()

  r.GET("/someJSON", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
  })

  r.GET("/moreJSON", func(c *gin.Context){
    var msg struct{
      Name    string `json:"user"`
      Message string
      Number  int
    }
    msg.Name = "yerek"
    msg.Message = "Helllo"
    msg.Number = 87471850499

    c.JSON(200, msg)
  })

  r.GET("/someXML", func(c *gin.Context){
    c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
  })

  r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

  r.GET("/someProtoBuff", func(c *gin.Context){
    resp := []int64{int64(1), int64(2)}
    label := "test"

    data := &protoexample.Test{
      Label: label,
      Resp:  resp,
    }

    c.ProtoBuf(200, data)
  })

  r.Run(":8080")
}
