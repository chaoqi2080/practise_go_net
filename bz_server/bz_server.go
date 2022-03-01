package main

import (
	"encoding/binary"
	"github.com/gorilla/websocket"
	"net/http"
	"practise_go_net/bz_server/handler"
	"practise_go_net/bz_server/msg"
	"practise_go_net/common/log"
	"practise_go_net/common/main_thread"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	if w == nil || r == nil {
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("websocket upgrade error:%v", err.Error())
		return
	}

	defer func() {
		conn.Close()
	}()

	log.Info("有新客户端连接")

	for {
		_, msgData, err := conn.ReadMessage()
		if err != nil {
			log.Error(err.Error())
			break
		}

		msgCode := binary.BigEndian.Uint16(msgData[2:4])

		newMsgX, err := msg.Decode(msgData[4:], int16(msgCode))
		if err != nil {
			log.Error("解码失败， %v", err.Error())
			continue
		}

		//log.Info(
		//	"收到消息: msgCode:%v, msgName:%v, newMsgX:%v",
		//	msgCode,
		//	newMsgX.Descriptor().Name(),
		//	newMsgX,
		//)

		cmdHandler := handler.CreateCmdHandler(msgCode)
		if cmdHandler == nil {
			log.Error(
				"没有找到消息处理器, msgCode:%v",
				msgCode,
			)
			continue
		}

		main_thread.Process(func() {
			cmdHandler(conn, newMsgX)
		})
	}
}

func main() {
	log.Config("/Users/chaoqi/gitcode/practise_go_net/log/game.log")
	log.Info("hello, the world.")

	http.HandleFunc("/websocket", websocketHandler)
	http.ListenAndServe(":54321", nil)
}
