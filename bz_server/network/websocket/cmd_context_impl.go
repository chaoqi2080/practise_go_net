package websocket

import (
	"encoding/binary"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/reflect/protoreflect"
	"practise_go_net/bz_server/handler"
	"practise_go_net/bz_server/msg"
	"practise_go_net/common/log"
	"practise_go_net/common/main_thread"
	"time"
)

const msgPerSecond = 16
const msgLimitSize = 1024

type WebsocketCmdContext struct {
	userId       int64
	remoteIpAddr string
	Conn         *websocket.Conn
	Chan         chan protoreflect.ProtoMessage
}

func (ctx *WebsocketCmdContext) BindUserId(userId int64) {
	ctx.userId = userId
}

func (ctx *WebsocketCmdContext) GetUserId() (userId int64, err error) {
	return ctx.userId, nil
}

func (ctx *WebsocketCmdContext) GetClientIpAddr() string {
	return "127.0.0.1"
}

func (ctx *WebsocketCmdContext) Write(msgObj protoreflect.ProtoMessage) {
	if msgObj == nil {
		return
	}

	ctx.Chan <- msgObj
}

func (ctx *WebsocketCmdContext) SendError(errorCode int, errorInfo string) {

}

func (ctx *WebsocketCmdContext) Disconnect() {

}

func (ctx *WebsocketCmdContext) LoopReadMsg() {
	if ctx.Conn == nil {
		return
	}

	for {
		_, msgData, err := ctx.Conn.ReadMessage()
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
			cmdHandler(ctx, newMsgX)
		})
	}
}

func (ctx *WebsocketCmdContext) LoopWriteMsgBack() {
	if ctx.Chan == nil {
		ctx.Chan = make(chan protoreflect.ProtoMessage, 64)
	}

	//控制每一个消息大小
	ctx.Conn.SetReadLimit(64 * msgLimitSize)

	//控制每秒的消息数量
	tStart := int64(0)
	tMsgNum := 0

	for {
		msgObj := <-ctx.Chan

		if msgObj == nil {
			continue
		}

		tNow := time.Now().UnixMilli()
		if (tNow - tStart) > 1000 {
			tMsgNum = 0
			tStart = tNow
		}

		if tMsgNum > msgPerSecond {
			return
		}

		tMsgNum++

		resultArray, err := msg.Encode(msgObj)

		if err != nil {
			log.Error(
				"组合消息失败:%v",
				err.Error(),
			)
		}

		err = ctx.Conn.WriteMessage(websocket.BinaryMessage, resultArray)

		if err != nil {
			log.Error(
				"发送消息失败:%v",
				err.Error(),
			)
		}
	}
}
