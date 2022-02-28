package msg

import (
	"errors"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
	"sync"
)

var msgNameAndMsgCodeMap = make(map[string]int16)
var msgCodeAndMsgDescMap = make(map[int16]protoreflect.MessageDescriptor)
var locker = &sync.Mutex{}

func getMsgCodeByName(msgName string) (int16, error) {
	if len(msgName) <= 0 {
		return -1, errors.New("输入参数不合法")
	}

	if len(msgNameAndMsgCodeMap) <= 0 {
		init2Map()
	}

	return msgNameAndMsgCodeMap[msgName], nil
}

func getMsgDescriptorByMsgCode(msgCode int16) (protoreflect.MessageDescriptor, error) {
	if msgCode < 0 {
		return nil, errors.New("输入参数不合法")
	}

	if len(msgCodeAndMsgDescMap) <= 0 {
		init2Map()
	}

	return msgCodeAndMsgDescMap[msgCode], nil
}

func init2Map() {
	locker.Lock()
	defer locker.Unlock()

	if len(msgNameAndMsgCodeMap) > 0 &&
		len(msgCodeAndMsgDescMap) > 0 {
		return
	}

	for k, v := range MsgCode_value {
		msgName := strings.ToLower(
			strings.Replace(k, "_", "", -1),
		)

		msgNameAndMsgCodeMap[msgName] = int16(v)
	}

	desLst := File_GameMsgProtocol_proto.Messages()

	for i := 0; i < desLst.Len(); i++ {
		desc := desLst.Get(i)

		msgName := strings.ToLower(
			strings.Replace(string(desc.Name()), "_", "", -1),
		)

		msgCode := msgNameAndMsgCodeMap[msgName]

		msgCodeAndMsgDescMap[msgCode] = desc
	}
}
