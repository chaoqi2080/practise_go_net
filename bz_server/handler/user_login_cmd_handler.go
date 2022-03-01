package handler

import (
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"practise_go_net/bz_server/msg"
	"practise_go_net/common/log"
)

func init() {
	MsgCodeAndHandlerMap[uint16(msg.MsgCode_USER_LOGIN_CMD.Number())] = userLoginCmdHandler
}

func userLoginCmdHandler(conn *websocket.Conn, message *dynamicpb.Message) {
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
}
