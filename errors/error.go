package errors

import "errors"

const (
	UnknownCode   = -1
	UnknowMessage = ""
)

type Exception struct {
	Code      int64  `json:"code"`
	ErrorName string `json:"ErrorName"`
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
		Code: UnknownCode,
	}
}
