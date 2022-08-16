package errors

import "errors"

const (
	UnknownCode   = -1
	UnknowMessage = ""
)

type Exception struct {
	Code       int64  `json:"code"`       // 错误码， 低五位是子错误码，高位是业务码
	Message    string `json:"message"`    // 中文描述 //现在前端大部分展示的是这个错误码
	ErrMessage string `json:"errMessage"` // 服务端内部的错误码，方便定位问题
	MeteData   map[string]interface{}
}

func (*Exception) Error() string {
	return ""
}

func FromError(err error) *Exception {
	if err == nil {
		return nil
	}
	if se := new(Exception); errors.As(err, &se) {
		return se
	}
	// todo grpc error

	// todo micro error
	return &Exception{
		Code:       UnknownCode,
		Message:    UnknowMessage,
		ErrMessage: UnknowMessage,
	}
}
