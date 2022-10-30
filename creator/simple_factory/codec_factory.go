package simple_factory

import (
	"fmt"
)

// NewCodec 序列化器工厂方法
func NewCodec(codec Codec) (CodecI, error) {
	switch codec {
	case CodecJson:
		return &jsonCodec{}, nil
	case CodecYaml:
		return &yamlCodec{}, nil
	case CodecGob:
		return &gobCodec{}, nil
		// 违反开闭原则,新增实现需要修改工厂方法
	}
	return nil, fmt.Errorf("invalid codec: %v", codec)
}
