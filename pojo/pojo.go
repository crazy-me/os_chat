package pojo

import "github.com/gorilla/websocket"

var (
	Clients          = make(map[Client]bool)
	BroadcastMessage = make(chan ResultStruct, 100) // 消息广播
	ClientRegister   = make(chan Client, 100)       // 客户端注册
	ClientUnregister = make(chan Client, 100)       // 客户端注销
)

// 客户端信息
type Client struct {
	Connect  *websocket.Conn
	Msg      chan []byte
	UserName string
}

// 消息体
type ResultStruct struct {
	Code     int
	Total    int
	UserName string
	Msg      string
}
