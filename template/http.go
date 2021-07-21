package template

import (
	"bytes"
	"text/template"
)

var registryTemp = `
{{$sName := .ServiceName}}
func (r *{{$sName}})Registry(engine *gin.Engine){
	{{- range .Methods }}
	engine.Handle("{{ .ReqMethod }}", "{{ .Path }}", receiver.{{ .MethodName }})
	{{- end }}
}
{{ range .Methods }}
func (r *{{$sName}}){{ .MethodName }}(ctx *gin.Context){

}
{{- end }}
`

type ServiceInfo struct {
	ServiceName string
	Methods     []HttpInfo
}

type HttpInfo struct {
	MethodName string
	ReqMethod  string
	Path       string
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
