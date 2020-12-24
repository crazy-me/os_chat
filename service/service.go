package service

import (
	"encoding/json"
	"github.com/crazy-me/os_chat/pojo"
	"github.com/gorilla/websocket"
)

func Start() {
	for {
		select { // 接收消息
		case msg := <-pojo.BroadcastMessage:
			data, _ := json.Marshal(msg)
			for client := range pojo.Clients {
				if msg.UserName == client.UserName {
					continue
				}
				_ = client.Connect.WriteMessage(websocket.TextMessage, data)
			}

		case user := <-pojo.ClientRegister:
			pojo.Clients[user] = true
			var resMsg pojo.ResultStruct
			resMsg.Code = 200
			resMsg.UserName = user.UserName
			resMsg.Total = len(pojo.Clients)
			resMsg.Msg = "加入聊天!"
			pojo.BroadcastMessage <- resMsg
		case exitClient := <-pojo.ClientUnregister:
			if pojo.Clients[exitClient] {
				delete(pojo.Clients, exitClient)
			}
			var resMsg pojo.ResultStruct
			resMsg.Code = 200
			resMsg.UserName = exitClient.UserName
			resMsg.Total = len(pojo.Clients)
			resMsg.Msg = "退出聊天!"
			pojo.BroadcastMessage <- resMsg

		}
	}
}
