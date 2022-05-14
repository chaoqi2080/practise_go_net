package handler

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"practise_go_net/bz_server/base"
	"practise_go_net/bz_server/mod/user/userdata"
	"practise_go_net/bz_server/mod/user/userlso"
	"practise_go_net/bz_server/msg"
	"practise_go_net/bz_server/network/broadcaster"
	"practise_go_net/common/lazy_save"
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

	user := userdata.GetUserGroup().GetByUserId(int64(userAttkCmd.TargetUserId))
	if user == nil {
		return
	}

	userAttkResult := &msg.UserAttkResult{
		AttkUserId:   uint32(ctx.GetUserId()),
		TargetUserId: userAttkCmd.TargetUserId,
	}

	broadcaster.Broadcast(userAttkResult)

	//minus hp
	var subtractHp int32 = 10

	userSubtractHpResult := &msg.UserSubtractHpResult{
		SubtractHp:   uint32(subtractHp),
		TargetUserId: userAttkCmd.TargetUserId,
	}

	user.CurrHp -= subtractHp

	broadcaster.Broadcast(userSubtractHpResult)

	lso := &userlso.UserLso{
		User: user,
	}

	lazy_save.SaveOrUpdate(lso)
}
