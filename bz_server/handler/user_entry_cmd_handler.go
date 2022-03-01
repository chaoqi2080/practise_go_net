package handler

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"practise_go_net/bz_server/msg"
	"practise_go_net/common/log"
)

func init() {
	MsgCodeAndHandlerMap[uint16(msg.MsgCode_USER_ENTRY_CMD.Number())] = userEnterCmdHandler
}

func userEnterCmdHandler(ctx MyCmdContext, message *dynamicpb.Message) {
	cmd := &msg.UserEntryCmd{}

	message.Range(func(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		cmd.ProtoReflect().Set(descriptor, value)
		return true
	})

	log.Info(
		"收到用户进入消息",
	)
}
