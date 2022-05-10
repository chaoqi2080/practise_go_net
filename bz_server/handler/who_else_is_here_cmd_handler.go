package handler

import (
	"google.golang.org/protobuf/types/dynamicpb"
	"practise_go_net/bz_server/base"
	"practise_go_net/bz_server/mod/user/userdata"
	"practise_go_net/bz_server/msg"
	"practise_go_net/common/log"
)

func init() {
	cmdHandlerMap[uint16(msg.MsgCode_WHO_ELSE_IS_HERE_CMD.Number())] = userWhoElseIsHereCmdHandler
}

//还有谁指令处理
func userWhoElseIsHereCmdHandler(ctx base.MyCmdContext, _ *dynamicpb.Message) {
	if ctx == nil ||
		ctx.GetUserId() <= 0 {
		return
	}

	log.Info(
		"收到<还有谁>消息, userId = %v",
		ctx.GetUserId(),
	)

	whoElseIsHereResult := &msg.WhoElseIsHereResult{}

	userAll := userdata.GetUserGroup().GetUserAll()

	for _, user := range userAll {
		if user == nil {
			continue
		}

		curMoveState := &msg.WhoElseIsHereResult_UserInfo_MoveState{
			FromPosX:  user.MoveState.MoveFromX,
			FromPosY:  user.MoveState.MoveFromY,
			ToPosX:    user.MoveState.MoveToX,
			ToPosY:    user.MoveState.MoveToY,
			StartTime: uint64(user.MoveState.StartTime),
		}

		whoElseIsHereResult.UserInfo = append(
			whoElseIsHereResult.UserInfo,
			&msg.WhoElseIsHereResult_UserInfo{
				UserId:     uint32(user.UserId),
				UserName:   user.UserName,
				HeroAvatar: user.HeroAvatar,
				MoveState:  curMoveState,
			},
		)
	}

	ctx.Write(whoElseIsHereResult)
}
