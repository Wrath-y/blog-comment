package resp

import (
	"bytes"
	"comment/infrastructure/util/errcode"
	"comment/infrastructure/util/util/highperf"
	"comment/interfaces/proto"
	"encoding/json"
)

func Success[T comparable](data T) (*proto.Response, error) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(data); err != nil {
		return nil, err
	}

	return &proto.Response{
		Code: 0,
		Msg:  "success",
		Data: highperf.Bytes2str(buf.Bytes()),
	}, nil
}

func FailWithErrCode(err *errcode.ErrCode) (*proto.Response, error) {
	return &proto.Response{
		Code: err.Code,
		Msg:  err.Msg,
	}, err
}
