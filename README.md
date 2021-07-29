# 自定义pb插件

## 要点

* pb文件的解析
* text包的使用

## 生成gin的handler

实战生成gin框架使用的handler

## 环境

在本地测试前，先执行命令 `make init` 进行环境初始化

改完代码后，执行命令 `make run` 来调用protoc编译测试的 `test.proto`文件

## 处理流程

### 原理图

从proto文件获取信息，转化为构建gin的handler所需的信息，最后使用text/template包按模板生成文件内容。

### 模块

* 模板模块：定义生成文件内容所需的信息结构体，提供生成文件内容的方法。
* 解析模块：解析proto文件，获取信息并组装成模板模块所需的信息结构体。

#### 解析模块

此模块需要用到Google的protobuf包(`google.golang.org/protobuf`)提供的api对proto文件的信息进行提取操作([参考链接](https://pkg.go.dev/google.golang.org/protobuf))










