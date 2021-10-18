package main

import (
  "github.com/gin-gonic/gin"
  "fmt"
  "time"
)

func logger(p gin.LogFormatterParams)string{
  return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
                p.ClientIP,
                p.TimeStamp.Format(time.RFC1123),
                p.Method,
                p.Path,
                p.Request,
                p.Request.Proto,
                p.StatusCode,
                p.Latency,
                p.Request.UserAgent(),
                p.ErrorMessage,
              )

}

func main(){
    router := gin.New()

    router.Use(gin.LoggerWithFormatter(logger))

    router.Use(gin.Recovery())

    router.GET("/", func(c *gin.Context){
      c.String(200, "Helloo")
    })
    router.Run(":8080")
}
