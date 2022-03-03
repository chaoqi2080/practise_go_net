package handler

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"practise_go_net/bz_server/msg"
	"practise_go_net/common/log"
)

func init() {
	MsgCodeAndHandlerMap[uint16(msg.MsgCode_USER_LOGIN_CMD.Number())] = userLoginCmdHandler
}

func userLoginCmdHandler(ctx MyCmdContext, message *dynamicpb.Message) {
	cmd := &msg.UserLoginCmd{}

	message.Range(func(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		cmd.ProtoReflect().Set(descriptor, value)
		return true
	})

	log.Info(
		"收到登录消息 userName:%v, password:%v",
		cmd.GetUserName(),
		cmd.GetPassword(),
	)

	result := &msg.UserLoginResult{
		UserId:     1,
		UserName:   cmd.GetUserName(),
		HeroAvatar: "Hero_Shaman",
	}

	ctx.BindUserId(1)
	ctx.Write(result)
}
