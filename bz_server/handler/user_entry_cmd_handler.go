package handler

import (
	"google.golang.org/protobuf/types/dynamicpb"
	"practise_go_net/bz_server/base"
	"practise_go_net/bz_server/mod/user/userdata"
	"practise_go_net/bz_server/msg"
	"practise_go_net/bz_server/network/broadcaster"
	"practise_go_net/common/log"
)

func init() {
	cmdHandlerMap[uint16(msg.MsgCode_USER_ENTRY_CMD.Number())] = userEntryCmdHandler
}

//用户入场指令处理器
func userEntryCmdHandler(ctx base.MyCmdContext, _ *dynamicpb.Message) {
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

	//互相引用，增加一个 c 包
	//websocket.GetCmdContextImplGroup().Broadcast(userEntryResult)
	broadcaster.Broadcast(userEntryResult)
}
