package main

import (
	"google.golang.org/protobuf/compiler/protogen"
)

const (
	fmtPackage = protogen.GoImportPath("fmt")
)

func main() {
	//fmt.Printf("protoc-gen-go-errors %v\n", "aaa")
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			if !f.Generate {
				continue
			}
			generateFile(plugin, f)
		}
		return nil
	})
}

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	filename := file.GeneratedFilenamePrefix + "_errors.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-errors. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	g.P("// this is your father")
	g.QualifiedGoIdent(fmtPackage.Ident(""))
	return g
}