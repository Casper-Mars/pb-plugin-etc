package main

import (
	"flag"

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
			_, err := generate(gen, f)
			if err != nil {
				panic(err)
			}
		}
		return nil
	})
}
