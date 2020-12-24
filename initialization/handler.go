package initialization

import (
	"github.com/crazy-me/os_chat/pojo"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	u = &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func WsHandler(c *gin.Context) {
	username := c.Query("username")
	conn, err := u.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	client := pojo.Client{
		Connect:  conn,
		Msg:      make(chan []byte),
		UserName: username,
	}

	// 新用户
	if !pojo.Clients[client] {
		pojo.ClientRegister <- client
	}

	// 断开用户连接
	defer func() {
		pojo.ClientUnregister <- client
		client.Connect.Close()
	}()

	// 读取消息
	for {
		_, message, err := client.Connect.ReadMessage()
		if err != nil {
			break
		}

		var msg pojo.ResultStruct
		msg.Code = 200
		msg.UserName = client.UserName
		msg.Msg = string(message)
		msg.Total = len(pojo.Clients)
		pojo.BroadcastMessage <- msg
	}
}
