package handler

import (
	"google.golang.org/protobuf/types/dynamicpb"
)

type CmdHandlerFun func(ctx MyCmdContext, message *dynamicpb.Message)

//消息 id 跟处理器对应 map
var MsgCodeAndHandlerMap = make(map[uint16]CmdHandlerFun)

func CreateCmdHandler(msgCode uint16) CmdHandlerFun {
	if msgCode < 0 {
		return nil
	}

	return MsgCodeAndHandlerMap[msgCode]
}