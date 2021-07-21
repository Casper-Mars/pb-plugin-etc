package template

import (
	"fmt"
	"testing"
)

func TestServiceInfo_Exec(t *testing.T) {

	info := ServiceInfo{
		ServiceName: "user",
		Methods: []HttpInfo{
			{
				MethodName: "list",
				ReqMethod:  "get",
				Path:       "/user/list",
			},
			{
				MethodName: "detail",
				ReqMethod:  "get",
				Path:       "/user/detail",
			},
		},
	}
	exec := info.Exec()
	fmt.Println(exec)
}
