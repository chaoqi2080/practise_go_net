package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	mywebsocket "practise_go_net/bz_server/network/websocket"
	"practise_go_net/common/log"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var sessionId int32 = 0

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

	//递增 sessionId
	sessionId += 1

	ctx := &mywebsocket.CmdContextImpl{
		Conn:      conn,
		SessionId: sessionId,
	}

	mywebsocket.GetCmdContextImplGroup().Add(ctx)
	defer mywebsocket.GetCmdContextImplGroup().RemoveBySessionId(ctx.SessionId)

	//循环发送消息
	ctx.LoopSendMsg()
	//循环读取消息
	ctx.LoopReadMsg()
}

func main() {
	log.Config("/Users/chaoqi/gitcode/practise_go_net/log/game")
	log.Info("hello, the world.")

	http.HandleFunc("/websocket", websocketHandler)
	http.ListenAndServe(":54321", nil)
}
