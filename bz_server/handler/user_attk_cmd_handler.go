package handler

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"practise_go_net/bz_server/base"
	"practise_go_net/bz_server/msg"
	"practise_go_net/bz_server/network/broadcaster"
)

func init() {
	cmdHandlerMap[uint16(msg.MsgCode_USER_ATTK_CMD.Number())] = userAttkCmdHandler
}

func userAttkCmdHandler(ctx base.MyCmdContext, pbMsgObj *dynamicpb.Message) {
	if ctx == nil || ctx.GetUserId() <= 0 || pbMsgObj == nil {
		return
	}

	userAttkCmd := &msg.UserAttkCmd{}

	pbMsgObj.Range(func(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		userAttkCmd.ProtoReflect().Set(descriptor, value)
		return true
	})

	userAttkResult := &msg.UserAttkResult{
		AttkUserId:   uint32(ctx.GetUserId()),
		TargetUserId: userAttkCmd.TargetUserId,
	}

	broadcaster.Broadcast(userAttkResult)

	//minus hp
	userSubtractHpResult := &msg.UserSubtractHpResult{
		SubtractHp:   10,
		TargetUserId: userAttkCmd.TargetUserId,
	}

	broadcaster.Broadcast(userSubtractHpResult)
}
