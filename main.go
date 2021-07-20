package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	flag.Parse()
	var flags flag.FlagSet
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {

		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generate(gen, f)
		}
		return nil
	})
}

func generate(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_test.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-myplugin. DO NOT EDIT .")
	g.P()
	g.P(fmt.Sprintf("package %s", file.GoPackageName))
	g.P()
	generateFunc(g, file)
	return g
}

func generateFunc(g *protogen.GeneratedFile, file *protogen.File) {

}
