package simple_factory

// Codec 编解码类型
type Codec int

const (
	CodecJson Codec = iota
	CodecYaml
	CodecGob
)

type CodecI interface {
	Marshal(v interface{}) ([]byte, error)
}
