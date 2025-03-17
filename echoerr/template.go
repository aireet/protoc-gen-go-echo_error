package echoerr

import (
	"bytes"
	"github.com/aireet/protoc-gen-go-echo_error/errors"
	"text/template"
)

type E errors.Exception

const (
	errTemplate = `

	func New{{.ErrorName}}Error(msg string) (err error) {
		err = status.Errorf(codes.Code({{.Code}}),msg)
		return err
	}

	func Is{{.ErrorName}}Error(err error) bool {
		status := status.Convert(err)
		if status == nil {
			return false
		}
		return  status.Code() == codes.Code({{.Code}})		
	}

	`
)

func (e *E) Gen() string {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("errors").Parse(errTemplate)
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, e); err != nil {
		panic(err)
	}

	return buf.String()

}
