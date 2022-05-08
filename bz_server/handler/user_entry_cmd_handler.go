package handler

import (
	"google.golang.org/protobuf/types/dynamicpb"
	"practise_go_net/bz_server/mod/user/userdata"
	"practise_go_net/bz_server/msg"
	"practise_go_net/common/log"
)

func init() {
	MsgCodeAndHandlerMap[uint16(msg.MsgCode_USER_ENTRY_CMD.Number())] = userEntryCmdHandler
}

//用户入场指令处理器
func userEntryCmdHandler(ctx MyCmdContext, message *dynamicpb.Message) {
	if ctx == nil || ctx.GetUserId() <= 0 {
		return
	}

	log.Info(
		"收到用户进入消息 userId = %v",
		ctx.GetUserId(),
	)

	user := userdata.GetUserGroup().GetByUserId(ctx.GetUserId())

	if user == nil {
		log.Error(
			"未找到用户数据, userId = %d",
			ctx.GetUserId(),
		)
		return
	}

	userEntryResult := &msg.UserEntryResult{
		UserId:     uint32(ctx.GetUserId()),
		UserName:   user.UserName,
		HeroAvatar: user.HeroAvatar,
	}

	ctx.Write(userEntryResult)
}
