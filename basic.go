package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "fmt"
)

func main(){
  r := gin.Default()
  r.GET("/user/:name", func(c *gin.Context){
    name := c.Param("name")
    c.String(http.StatusOK, "Hello %s", name)
  })

  r.GET("/user/:name/*action", func(c *gin.Context){
    name := c.Param("name")
    action := c.Param("action")
    message := name + " is " + action
    c.String(http.StatusOK, message)
  })

  r.POST("/user/:name/*action", func(c *gin.Context){
    b := c.FullPath() == "/user/:name/*action"
    c.String(http.StatusOK, "%t", b)
  })
  r.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})

  r.GET("/welcome", func(c *gin.Context){
    firstname := c.DefaultQuery("firstname", "Guest")
    lastname := c.Query("lastname")

    c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
  })

  r.POST("/form_post", func(c *gin.Context){
    message := c.PostForm("message")
    nick := c.DefaultPostForm("nick", "anonymous")

    c.String(http.StatusOK, "status: posted, nick = %s, message = %s",nick, message)
  })
  r.POST("/post", func(c *gin.Context){
    id := c.Query("id")
    page := c.DefaultQuery("1", "0")
    name := c.PostForm("name")
    message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
    c.String(http.StatusOK,  id, page, name, message)

  })

  r.POST("/posts", func(c *gin.Context) {
    ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
    c.String(http.StatusOK, "%s %s", ids, names)
	})

  r.Run(":8080")

}
