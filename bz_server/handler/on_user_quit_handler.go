package handler

import (
	"practise_go_net/bz_server/base"
	"practise_go_net/bz_server/msg"
	"practise_go_net/bz_server/network/broadcaster"
)

//用户断线

func OnUserQuitHandler(ctx base.MyCmdContext) {
	if ctx == nil ||
		ctx.GetUserId() <= 0 {
		return
	}

	result := &msg.UserQuitResult{
		QuitUserId: uint32(ctx.GetUserId()),
	}

	broadcaster.Broadcast(result)
}
