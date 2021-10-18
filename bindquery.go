package main

import (
  "github.com/gin-gonic/gin"
  "log"
  "time"
  "net/http"
)

type Person struct{
  Name       string     `form:"name"`
  Address    string     `form:"address"`
  Birthday   time.Time  `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
  CreateTime time.Time  `form:"createTime" time_format:"unixNano"`
  UnixTime   time.Time  `form:"unixTime" time_format:"unix"`
}

func startPage(c *gin.Context){
  var person Person

  if c.ShouldBind(&person) == nil{
                log.Println(person.Name)
                log.Println(person.Address)
                log.Println(person.Birthday)
                log.Println(person.CreateTime)
                log.Println(person.UnixTime)
  }
  c.JSON(http.StatusCreated, gin.H{"name" : person.Name,
                                   "address" : person.Address,
                                   "birthday" : person.Birthday,
                                   "createTime" : person.CreateTime,
                                   "unixTime" : person.UnixTime,
                                 })
}

func main(){
  r := gin.Default()
  r.GET("/test", startPage)
  r.Run(":8080")
}
