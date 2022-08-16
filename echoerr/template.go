package echoerr

import (
	"bytes"
	"github.com/aireet/protoc-gen-go-echo_error/errors"
	"text/template"
)

type E errors.Exception

const (
	errTemplate = `

	func New{{.ErrMessage}}Error() (err *errors.Exception) {
		err = &errors.Exception{
			Code:  {{.Code}},
			Message: "{{.Message}}",
			ErrMessage: "{{.ErrMessage}}",
		}
		return err
	}

	func Is{{.ErrMessage}}(err error) bool {
		if err == nil {
			return false
		}
		e := errors.FromError(err)
		return e.Code == {{.Code}} && e.ErrMessage == "{{.ErrMessage}}"
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
