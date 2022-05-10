package handler

import (
	"google.golang.org/protobuf/types/dynamicpb"
	"practise_go_net/bz_server/base"
)

type CmdHandlerFun func(ctx base.MyCmdContext, pbMsgObj *dynamicpb.Message)

//消息 id 跟处理器对应 map
var cmdHandlerMap = make(map[uint16]CmdHandlerFun)

func CreateCmdHandler(msgCode uint16) CmdHandlerFun {
	if msgCode < 0 {
		return nil
	}

	return cmdHandlerMap[msgCode]
}
