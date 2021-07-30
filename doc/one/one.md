# 自定义protoc插件（1）

## 目录

- [目录](##目录)
- [项目说明](##项目说明)
- [初始化项目](##初始化项目)
    - [建立项目](####建立项目)
    - [添加依赖](####添加依赖)
    - [创建用户protobuf描述文件](####创建用户protobuf描述文件)

## 项目说明

需求：自定义protoc插件生成gin需要的handler。

## 初始化项目

#### 建立项目

```shell
go mod init protoc-gen-myplugin
```

#### 添加依赖

* gin

```shell
go get -u github.com/gin-gonic/gin
```

* protobuf

```shell
go get google.golang.org/protobuf@v1.27.1 
```

* genproto

```shell
go get google.golang.org/genproto/...
```

#### 创建用户protobuf描述文件

```shell
mkdir api
touch api/user.proto
```

内容如下：

```protobuf
syntax = "proto3";

package api;


import "google/api/annotations.proto";
import "third_party/google/api/annotations.proto";

// 定义包目录，其中";"后面定义的是额外指定的包名称
option go_package = "./;out";

message UserInfo {
    uint64 id = 1;
    string account = 2;
    string name = 3;
    string avatar = 4;
}

message UserListRequest {

}

message UserList {
    repeated UserListItem list = 1;
}

message UserListItem {

}

message UserInfoRequest {
    uint64 id = 1;
}

service user {
    rpc GetUserInfo(UserInfoRequest) returns(UserInfo) {
        option (google.api.http) = {
            get: "/user/info"
        };
    };
    rpc List(UserListRequest) returns(UserList) {
        option (google.api.http) = {
            get: "/user/list"
        };
    }
}
```

#### 创建模板

模板定义handler对象的信息和渲染生成的源文件。

基础阶段，此处只定义handler的参数。

```go
// 服务信息
type ServiceInfo struct {
// 服务名称
ServiceName string
// 接口列表
Methods     []HttpInfo
}

// 接口信息
type HttpInfo struct {
// 接口名称
MethodName string
// 接口请求方法：GET、POST等
ReqMethod  string
// 接口请求url
Path       string
// 是否有请求体
HasBody    bool
}
```

定义模板，使用`text/template`包处理模板渲染的问题。`text/template`包的使用参考[官方文档](https://pkg.go.dev/text/template)

``` go
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
```












