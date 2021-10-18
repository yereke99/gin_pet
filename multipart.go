package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

type ProfileForm struct{
  Name string                   `form:"name" binding:"required"`
  Avatar *multipart.FileHeader  `form:"avatar" binding:"required"`
}

func main(){
  r := gin.Default()

  r.POST("/profile", func(c *gin.Context){
    var form ProfileForm

    if err := c.ShouldBind(&form); err != nil{
      c.String(http.StatusBadRequest, "bad request")
      return
    }

    err := c.SaveUploadedFile(form.Avatar, form.Avatar.Filename)
    if err != nil{
      c.String(http.StatusInternalServerError, "unknown error")
			return
    }

    c.String(200, "ok")

  })

  r.Run(":8080")


}
