package echoerr

import (
	"google.golang.org/protobuf/compiler/protogen"
	"strings"
)

func GenError(g *protogen.GeneratedFile, file *protogen.File) {

	errList := make([]*E, 0)

	var hasMessage bool

	for _, message := range file.Messages {

		if !strings.Contains(message.GoIdent.GoName, "EchoError") {
			continue
		}

		for _, enum := range message.Enums {

			hasMessage = true

			for _, value := range enum.Values {
				errList = append(errList, packException("", value))
			}
		}

	}

	if !hasMessage {
		return
	}

	genErrors(g, errList)
}

func genErrors(g *protogen.GeneratedFile, errList []*E) {
	for _, e := range errList {
		if e == nil {
			continue
		}
		g.P()
		g.P(e.Gen())
	}

}

func packException(msgName string, enumValue *protogen.EnumValue) *E {

	if enumValue.Desc.Number() == 0 {
		return nil
	}

	code := int64(enumValue.Desc.Number())

	e := &E{
		Code:      code,
		ErrorName: enumValue.GoIdent.GoName,
	}

	return e

}
