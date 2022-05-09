package broadcaster

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"practise_go_net/bz_server/base"
)

var innerMap = make(map[int32]base.MyCmdContext)

func AddCmdCtx(sessionId int32, ctx base.MyCmdContext) {
	if ctx == nil {
		return
	}

	innerMap[sessionId] = ctx
}

func RemoveCmdCtxBySessionId(sessionId int32) {
	if sessionId <= 0 {
		return
	}

	delete(innerMap, sessionId)
}

func Broadcast(msgObj protoreflect.ProtoMessage) {
	if msgObj == nil {
		return
	}

	for _, ctx := range innerMap {
		if ctx != nil {
			ctx.Write(msgObj)
		}
	}
}
