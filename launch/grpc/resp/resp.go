package resp

import (
	"bytes"
	"comment/infrastructure/common/errcode"
	"comment/infrastructure/util/util/highperf"
	"comment/interfaces/dto"
	"comment/interfaces/proto"
	"github.com/bytedance/sonic"
	"strconv"
)

func Success(data any) (*proto.Response, error) {
	res := &proto.Response{
		Code: 0,
		Msg:  "success",
	}
	switch data.(type) {
	case int8:
		res.Data = strconv.FormatInt(int64(data.(int8)), 10)
	case int16:
		res.Data = strconv.FormatInt(int64(data.(int16)), 10)
	case int32:
		res.Data = strconv.FormatInt(int64(data.(int32)), 10)
	case int64:
		res.Data = strconv.FormatInt(data.(int64), 10)
	case string:
		res.Data = data.(string)
	case nil:
		res.Data = ""
	case dto.H:
		res.Data = ""
	default:
		buf := bytes.NewBuffer(nil)
		enc := sonic.ConfigDefault.NewEncoder(buf)
		if err := enc.Encode(data); err != nil {
			return nil, err
		}

		res.Data = highperf.Bytes2str(buf.Bytes())
	}

	return res, nil
}

func FailWithErrCode(err *errcode.ErrCode) (*proto.Response, error) {
	return &proto.Response{
		Code: err.Code,
		Msg:  err.Msg,
	}, err
}
