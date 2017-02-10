package main

import (
	"github.com/gin-gonic/gin"
	"mytube"
	"os"
)

func main() {
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8001"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/channels", func(c *gin.Context) {
		c.JSON(200, mytube.Channels())
	})

	router.GET("/videos/:channel", func(c *gin.Context) {
		channel := c.Params.ByName("channel")
		c.JSON(200, mytube.VideosByChannel(channel))
	})

	router.Run(":" + PORT)
}
