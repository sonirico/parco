package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strconv"
	"strings"
)

// source is the package as parco-gen sees it: every struct type declared in
// it, in declaration order, and the imports its files use so a field typed
// time.Time can be resolved to its import path.
type source struct {
	pkg     string
	fset    *token.FileSet
	order   []string
	structs map[string]*ast.StructType
	imports map[string]string
}

func parseSource(dir string) (source, error) {
	src := source{
		fset:    token.NewFileSet(),
		structs: make(map[string]*ast.StructType),
		imports: make(map[string]string),
	}
	pkgs, err := parser.ParseDir(src.fset, dir, func(fi os.FileInfo) bool {
		return !strings.HasSuffix(fi.Name(), "_test.go")
	}, parser.ParseComments)
	if err != nil {
		return src, err
	}
	for name, pkg := range pkgs {
		if strings.HasSuffix(name, "_test") {
			continue
		}
		if src.pkg != "" {
			return src, fmt.Errorf("more than one package in %s: %s and %s", dir, src.pkg, name)
		}
		src.pkg = name
		for _, file := range pkg.Files {
			if ast.IsGenerated(file) {
				continue
			}
			src.collect(file)
		}
	}
	if src.pkg == "" {
		return src, fmt.Errorf("no Go package in %s", dir)
	}
	return src, nil
}

func (s *source) collect(file *ast.File) {
	for _, spec := range file.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			continue
		}
		name := path[strings.LastIndex(path, "/")+1:]
		if spec.Name != nil {
			name = spec.Name.Name
		}
		s.imports[name] = path
	}
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}
		for _, spec := range gen.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}
			s.order = append(s.order, ts.Name.Name)
			s.structs[ts.Name.Name] = st
		}
	}
}
