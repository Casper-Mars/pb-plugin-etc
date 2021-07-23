package template

import (
	"bytes"
	"text/template"
)

var registryTemp = `
{{$sName := .ServiceName}}
type {{$sName}}HttpServer interface {
	{{- range .Methods }}
	{{.MethodName}}(ctx *gin.Context)
	{{- end }}
}

func Registry{{$sName}}HttpServer(engine *gin.Engine, srv {{$sName}}HttpServer) {
	{{- range .Methods }}
	engine.Handle("{{ .ReqMethod }}", "{{ .Path }}", srv.{{ .MethodName }})
	{{- end }}
}
`

type ServiceInfo struct {
	ServiceName string
	Methods     []HttpInfo
}

type HttpInfo struct {
	MethodName string
	ReqMethod  string
	Path       string
	HasBody    bool
}

func (receiver *ServiceInfo) Exec() string {
	b := new(bytes.Buffer)
	temp := template.Must(template.New("http").Parse(registryTemp))

	err := temp.Execute(b, receiver)
	if err != nil {
		panic(err)
	}
	return b.String()
}
