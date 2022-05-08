package websocket

import "google.golang.org/protobuf/reflect/protoreflect"

type cmdCtxImplGroup struct {
	innerMap map[int32]*CmdContextImpl
}

var cmdContextImplGroupInstance = &cmdCtxImplGroup{
	innerMap: make(map[int32]*CmdContextImpl),
}

func GetCmdContextImplGroup() *cmdCtxImplGroup {
	return cmdContextImplGroupInstance
}

func (group *cmdCtxImplGroup) Add(ctx *CmdContextImpl) {
	if ctx == nil {
		return
	}

	group.innerMap[ctx.SessionId] = ctx
}

func (group *cmdCtxImplGroup) RemoveBySessionId(sessionId int32) {
	if sessionId <= 0 {
		return
	}

	delete(group.innerMap, sessionId)
}

func (group *cmdCtxImplGroup) Broadcast(msgObj protoreflect.ProtoMessage) {
	if msgObj == nil {
		return
	}

	for _, ctx := range group.innerMap {
		if ctx != nil {
			ctx.Write(msgObj)
		}
	}
}
