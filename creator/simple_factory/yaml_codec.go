package simple_factory

import "gopkg.in/yaml.v3"

// yamlCodec
// @Description: yaml编码器
type yamlCodec struct {
}

func (y yamlCodec) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}
