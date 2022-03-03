package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	mysocket "practise_go_net/bz_server/network/websocket"
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

	ctx := &mysocket.WebsocketCmdContext{
		Conn: conn,
	}

	go ctx.LoopWriteMsg()
	ctx.LoopReadMsg()
}

func main() {
	log.Config("/Users/chaoqi/gitcode/practise_go_net/log/game.log")
	log.Info("hello, the world.")

	http.HandleFunc("/websocket", websocketHandler)
	http.ListenAndServe(":54321", nil)
}
