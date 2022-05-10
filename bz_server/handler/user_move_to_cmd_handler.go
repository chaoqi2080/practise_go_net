package handler

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
	"practise_go_net/bz_server/base"
	"practise_go_net/bz_server/mod/user/userdata"
	"practise_go_net/bz_server/msg"
	"practise_go_net/bz_server/network/broadcaster"
	"time"
)

func init() {
	cmdHandlerMap[uint16(msg.MsgCode_USER_MOVE_TO_RESULT.Number())] = userMoveToCmdHandler
}

func userMoveToCmdHandler(ctx base.MyCmdContext, pbMsgObj *dynamicpb.Message) {
	if ctx == nil ||
		ctx.GetUserId() <= 0 {
		return
	}

	user := userdata.GetUserGroup().GetByUserId(ctx.GetUserId())

	if user == nil {
		return
	}

	userMoveToCmd := &msg.UserMoveToCmd{}

	pbMsgObj.Range(func(f protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		userMoveToCmd.ProtoReflect().Set(f, v)
		return true
	})

	if user.MoveState == nil {
		user.MoveState = &userdata.MoveState{}
	}

	tNow := time.Now().UnixMilli()

	user.MoveState.MoveFromX = userMoveToCmd.MoveFromPosX
	user.MoveState.MoveFromY = userMoveToCmd.MoveFromPosY
	user.MoveState.MoveToX = userMoveToCmd.MoveToPosX
	user.MoveState.MoveToY = userMoveToCmd.MoveToPosY
	user.MoveState.StartTime = tNow

	userMoveToResult := &msg.UserMoveToResult{
		MoveUserId:    uint32(ctx.GetUserId()),
		MoveFromPosX:  userMoveToCmd.MoveFromPosX,
		MoveFromPosY:  userMoveToCmd.MoveFromPosY,
		MoveToPosX:    userMoveToCmd.MoveToPosX,
		MoveToPosY:    userMoveToCmd.MoveToPosY,
		MoveStartTime: uint64(tNow),
	}

	broadcaster.Broadcast(userMoveToResult)
}
