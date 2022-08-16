package example

import (
	pb "github.com/aireet/protoc-gen-go-echo_error/proto"
)

func example() {
	//  断言
	// err = doRpc
	var err error
	// 通过 proto 生成的代码断言错误
	if pb.IsMallErrNoPermission(err) {
		//...
	}

	// 赋值
	err = pb.NewMallErrNoPermissionError()

}
