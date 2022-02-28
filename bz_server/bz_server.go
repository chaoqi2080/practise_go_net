package main

import (
	"encoding/binary"
	"github.com/gorilla/websocket"
	"net/http"
	"practise_go_net/bz_server/msg"
	"practise_go_net/common/log"
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

		cmd, err := msg.Decode(msgData[4:], int16(msgCode))
		if err != nil {
			log.Error("解码失败， %v", err.Error())
			continue
		}

		log.Info(
			"收到消息: msgCode:%v, msgName:%v, cmd:%v",
			msgCode,
			cmd.Descriptor().Name(),
			cmd,
		)
	}

}

func main() {
	log.Config("/Users/chaoqi/gitcode/practise_go_net/log/game.log")
	log.Info("hello, the world.")

	http.HandleFunc("/websocket", websocketHandler)
	http.ListenAndServe(":54321", nil)
}
