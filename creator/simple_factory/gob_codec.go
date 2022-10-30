package simple_factory

import (
	"bytes"
	"encoding/gob"
)

// gobCodec
// @Description: gob编码器
type gobCodec struct {
}

func (g gobCodec) Marshal(v interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
