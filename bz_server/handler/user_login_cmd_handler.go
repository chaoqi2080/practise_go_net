package handler

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"practise_go_net/bz_server/mod/login/loginsrv"
	"practise_go_net/bz_server/mod/user/userdata"
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
		cmd.UserName,
		cmd.Password,
	)

	returnedObj := loginsrv.LoginByPasswordAsync(cmd.UserName, cmd.Password)

	if returnedObj == nil {
		log.Error(
			"返回的 obj 为空",
			cmd.UserName,
		)
		return
	}

	returnedObj.OnComplete(func() {
		user := returnedObj.GetReturnedObj().(*userdata.User)
		if user == nil {
			log.Error(
				"用户不存在 :%v",
				cmd.UserName,
			)
			return
		}

		result := &msg.UserLoginResult{
			UserId:     uint32(user.UserId),
			UserName:   user.UserName,
			HeroAvatar: user.HeroAvatar,
		}

		ctx.BindUserId(user.UserId)
		ctx.Write(result)
	})
}
