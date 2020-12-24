package main

import (
	"fmt"
	"github.com/crazy-me/os_chat/initialization"
	"github.com/crazy-me/os_chat/service"
	"github.com/crazy-me/os_chat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	port     int
	host, ws string
)

func init() {
	go service.Start()
}

func main() {
	var r = gin.Default()
	port = 8800
	host = utils.GetOutboundIP()
	ws = "//" + host + ":" + strconv.Itoa(port)
	r.StaticFS("/static", http.Dir("resource/static"))
	r.StaticFile("/favicon.ico", "resource/static/icon/favicon.ico")
	r.LoadHTMLFiles("resource/view/index.html")
	r.GET("/ws", initialization.WsHandler)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ws": ws,
		})
	})
	fmt.Printf("run server success! http://%s:%d\n", host, port)
	r.Run(":8800")

}
