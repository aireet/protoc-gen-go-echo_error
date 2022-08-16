package echoerr

import (
	"google.golang.org/protobuf/compiler/protogen"
	"log"
	"os"
	"strings"
)

func GenError(g *protogen.GeneratedFile, file *protogen.File) {

	errList := make([]*E, 0)

	for _, message := range file.Messages {

		if message.GoIdent.GoName != "EchoError" {
			continue
		}

		for _, enum := range message.Enums {

			log.Println(enum.GoIdent.GoName)

			if enum.GoIdent.GoName != "EchoError_Code" {
				continue
			}

			for _, value := range enum.Values {
				errList = append(errList, packException(31, value))
			}
		}

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

func packException(projectId int64, enumValue *protogen.EnumValue) *E {

	if enumValue.Desc.Number() == 0 {
		return nil
	}

	code := projectId*100000 + int64(enumValue.Desc.Number())

	em := strings.Trim(enumValue.GoIdent.GoName, "EchoError_")

	if enumValue.Comments.Leading.String() == "" {
		log.Fatal("comments is required")
		os.Exit(1)
	}

	m := strings.Trim(enumValue.Comments.Leading.String(), "//")
	m = strings.Replace(m, " ", "", -1)
	m = strings.Replace(m, "\n", "", -1)

	e := &E{
		Code:       code,
		Message:    m,
		ErrMessage: em,
	}

	return e

}
