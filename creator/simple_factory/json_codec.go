package simple_factory

import "encoding/json"

// jsonCodec
// @Description: json编码器
type jsonCodec struct {
}

func (j jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
