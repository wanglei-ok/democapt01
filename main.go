package main

import (
	"democapt01/gohksdk"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"os"
	"strconv"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.Use(cors())
	router.GET("/captuerPicture", CapturePicture)
	router.Run(":10086")
}

func CapturePicture(c *gin.Context) {

	channel := 1
	port := 8000
	if v, ok := c.GetQuery("channel"); ok {
		channel, _ = strconv.Atoi(v)
	}

	if v, ok := c.GetQuery("port"); ok {
		port, _ = strconv.Atoi(v)
	}
	ip, _ := c.GetQuery("ip")
	username, _ := c.GetQuery("username")
	password, _ := c.GetQuery("password")

	id := xid.New()
	saveFile := fmt.Sprintf("temp/%v.jpg", id.String())
	ret := gohksdk.CapturePicture(channel, port, ip, username, password, saveFile)
	if ret != 0 {
		c.AbortWithStatus(http.StatusBadGateway)
		return
	}

	defer os.Remove(saveFile)
	c.File(saveFile)
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
