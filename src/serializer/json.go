package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtobufToJSON converts protocol buffer message to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		UseProtoNames:   true,   // 对应 OrigName: true 
		EmitUnpopulated: true,   // 对应 EmitDefaults: true 
		Indent:          "  ",   // 缩进两个空格 
		Multiline:       true,   // 多行格式化（需与 Indent 配合）
	}

	// 将 Protobuf 消息转换为字节切片 
	jsonBytes, err := marshaler.Marshal(message) 
	if err!= nil { 
		return "", err 
	} 
 
	// 将字节切片转换为字符串 
	return string(jsonBytes), nil 
}

// JSONToProtobufMessage converts JSON string to protocol buffer message
func JSONToProtobufMessage(data string, message proto.Message) error {
	opts := protojson.UnmarshalOptions{
        AllowPartial:   true,   // 允许部分解析（非必填字段缺失时不报错）
        DiscardUnknown: true,   // 忽略 JSON 中的未知字段 
    }
    return opts.Unmarshal([]byte(data), message)
}
