package base

import "google.golang.org/protobuf/reflect/protoreflect"

type MyCmdContext interface {
	BindUserId(userId int64)
	GetUserId() (userId int64)
	GetClientIpAddr() string
	Write(msgObj protoreflect.ProtoMessage)
	SendError(errorCode int, errorInfo string)
	Disconnect()
}
