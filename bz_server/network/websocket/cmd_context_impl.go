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
const oneSecond = 1000
const msgLimitSize = 1024

type CmdContextImpl struct {
	userId       int64
	remoteIpAddr string
	Conn         *websocket.Conn
	sendMsgQ     chan protoreflect.ProtoMessage
	SessionId    int32
}

func (ctx *CmdContextImpl) BindUserId(userId int64) {
	ctx.userId = userId
}

func (ctx *CmdContextImpl) GetUserId() (userId int64) {
	return ctx.userId
}

func (ctx *CmdContextImpl) GetClientIpAddr() string {
	return ctx.Conn.RemoteAddr().String()
}

func (ctx *CmdContextImpl) Write(msgObj protoreflect.ProtoMessage) {
	if msgObj == nil {
		return
	}

	ctx.sendMsgQ <- msgObj
}

func (ctx *CmdContextImpl) SendError(errorCode int, errorInfo string) {

}

func (ctx *CmdContextImpl) Disconnect() {

}

func (ctx *CmdContextImpl) LoopSendMsg() {
	//构建发送队列
	if ctx.sendMsgQ == nil {
		ctx.sendMsgQ = make(chan protoreflect.ProtoMessage, 64)
	}

	go func() {
		for {
			msgObj := <-ctx.sendMsgQ

			if msgObj == nil {
				continue
			}

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
	}()
}

func (ctx *CmdContextImpl) LoopReadMsg() {
	if ctx.Conn == nil {
		return
	}

	//控制每一个消息大小
	ctx.Conn.SetReadLimit(64 * msgLimitSize)

	//控制每秒的消息数量
	t0 := int64(0)
	counter := 0

	for {
		_, msgData, err := ctx.Conn.ReadMessage()
		if err != nil {
			log.Error(err.Error())
			break
		}

		t1 := time.Now().UnixMilli()
		if (t1 - t0) > oneSecond {
			counter = 0
			t0 = t1
		}

		if counter > msgPerSecond {
			log.Error("消息过于频繁")
			return
		}

		counter++

		msgCode := binary.BigEndian.Uint16(msgData[2:4])

		newMsgX, err := msg.Decode(msgData[4:], int16(msgCode))
		if err != nil {
			log.Error("解码失败， %v", err.Error())
			continue
		}

		log.Info(
			"收到消息: msgCode:%v, msgName:%v",
			msgCode,
			newMsgX.Descriptor().Name(),
		)

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

	//广播用户断线消息
	handler.OnUserQuitHandler(ctx)
}
