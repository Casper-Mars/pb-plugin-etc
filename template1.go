package main

import (
	"bytes"
	"text/template"
)

var stringTemplate = `
func (t {{.MsgName}})ToString() string{
	return "test"
}
`

type StringFunc struct {
	MsgName string
}

func (receiver StringFunc) ToString() string {

	buf := new(bytes.Buffer)
	tmp := template.Must(template.New("string").Parse(stringTemplate))

	err := tmp.Execute(buf, receiver)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
