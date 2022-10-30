package simple_factory

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCodec(t *testing.T) {
	data := map[string]string{"name": "zhangsan"}
	type args struct {
		codec Codec
	}
	gobBytes, _ := hex.DecodeString("0eff81040102ff8200010c010c000012ff820001046e616d65087a68616e6773616e")
	tests := []struct {
		name           string
		args           args
		wantNewErr     bool
		wantBytes      []byte
		wantMarshalErr bool
	}{
		{
			name: "json codec",
			args: args{
				codec: CodecJson,
			},
			wantNewErr:     false,
			wantBytes:      []byte(`{"name":"zhangsan"}`),
			wantMarshalErr: false,
		},
		{
			name: "yaml codec",
			args: args{
				codec: CodecYaml,
			},
			wantNewErr:     false,
			wantBytes:      []byte("name: zhangsan\n"),
			wantMarshalErr: false,
		},
		{
			name: "gob codec",
			args: args{
				codec: CodecGob,
			},
			wantNewErr:     false,
			wantBytes:      gobBytes,
			wantMarshalErr: false,
		},
		{
			name: "invalid codec",
			args: args{
				codec: -1,
			},
			wantNewErr:     true,
			wantBytes:      nil,
			wantMarshalErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCodec(tt.args.codec)
			if (err != nil) != tt.wantNewErr {
				t.Errorf("NewCodec() error = %v, wantErr %v", err, tt.wantNewErr)
				return
			}
			if err != nil {
				return
			}
			bytes, err := got.Marshal(data)
			if (err != nil) != tt.wantMarshalErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantMarshalErr)
				return
			}
			assert.Equal(t, tt.wantBytes, bytes)
		})
	}
}
