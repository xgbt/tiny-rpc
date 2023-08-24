package serializer

import (
	"errors"

	"google.golang.org/protobuf/proto"
)

// ErrNotImplementedProtoMessage refers to param not implemented by proto.Message
var ErrNotImplementedProtoMessage = errors.New("param does not implement proto.Message")

var Proto = ProtoSerializer{}

// ProtoSerializer implements the Serializer interface
type ProtoSerializer struct {
}

func (ProtoSerializer) Marshal(message interface{}) ([]byte, error) {
	var body proto.Message
	if message == nil {
		return []byte{}, nil
	}

	var ok bool
	if body, ok = message.(proto.Message); !ok {
		return nil, ErrNotImplementedProtoMessage
	}
	
	return proto.Marshal(body)
}

func (ProtoSerializer) Unmarshal(data []byte, message interface{}) error {
	var body proto.Message
	if message == nil {
		return nil
	}

	var ok bool
	body, ok = message.(proto.Message)
	if !ok {
		return ErrNotImplementedProtoMessage
	}

	return proto.Unmarshal(data, body)
}
