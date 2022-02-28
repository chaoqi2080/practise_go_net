package msg

import (
	"encoding/binary"
	"errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

func Encode(msgObj protoreflect.ProtoMessage) ([]byte, error) {
	if msgObj == nil {
		return nil, errors.New("传入参数有误")
	}

	msgCode, err := getMsgCodeByName(string(msgObj.ProtoReflect().Descriptor().Name()))
	if err != nil {
		return nil, err
	}

	msgSizeByteArray := make([]byte, 2)
	binary.BigEndian.PutUint16(msgSizeByteArray, uint16(msgCode))

	msgCodeByteArray := make([]byte, 2)
	binary.BigEndian.PutUint16(msgCodeByteArray, uint16(msgCode))

	msgBodyArray, err := proto.Marshal(msgObj)
	if err != nil {
		return nil, err
	}

	completedMsg := append(msgSizeByteArray, msgCodeByteArray...)
	completedMsg = append(completedMsg, msgBodyArray...)

	return completedMsg, nil
}

func Decode(msgData []byte, msgCode int16) (protoreflect.Message, error) {
	if len(msgData) <= 0 || msgCode < 0 {
		return nil, errors.New("传入参数为空")
	}

	desc, err := getMsgDescriptorByMsgCode(msgCode)

	if err != nil {
		return nil, err
	}

	msg := dynamicpb.NewMessage(desc)

	err = proto.Unmarshal(msgData, msg)
	if err != nil {
		return nil, err
	}

	return msg, nil
}
