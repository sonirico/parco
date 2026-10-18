package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"reflect"
	"strings"
)

// wire is what a field becomes on the model: the parco.Type expression and
// the getter and setter bodies around the Go field.
type wire struct {
	typeExpr string
	valueTyp string
	get      string
	set      string
}

type tagOptions struct {
	skip     bool
	bigEnd   bool
	header   string
	location bool
}

func parseTag(raw string) (tagOptions, error) {
	var opts tagOptions
	tag, ok := reflect.StructTag(raw).Lookup("parco")
	if !ok {
		return opts, nil
	}
	for _, opt := range strings.Split(tag, ",") {
		switch {
		case opt == "":
		case opt == "-":
			opts.skip = true
		case opt == "le":
			opts.bigEnd = false
		case opt == "be":
			opts.bigEnd = true
		case opt == "location":
			opts.location = true
		case strings.HasPrefix(opt, "hdr="):
			opts.header = strings.TrimPrefix(opt, "hdr=")
		default:
			return opts, fmt.Errorf("unknown tag option %q", opt)
		}
	}
	return opts, nil
}

func (o tagOptions) order() string {
	if o.bigEnd {
		return "BE"
	}
	return "LE"
}

var headers = map[string]string{
	"uint8":  "parco.UInt8Header()",
	"int8":   "parco.Int8Header()",
	"uint16": "parco.UInt16Header%s()",
	"int16":  "parco.Int16%sHeader()",
	"uint32": "parco.UInt32%sHeader()",
	"int32":  "parco.Int32%sHeader()",
	"uint64": "parco.UInt64%sHeader()",
	"int64":  "parco.Int64%sHeader()",
	"uint":   "parco.UInt%sHeader()",
	"int":    "parco.Int%sHeader()",
}

func (o tagOptions) headerExpr() (string, error) {
	name := o.header
	if name == "" {
		name = "uint16"
	}
	tmpl, ok := headers[name]
	if !ok {
		return "", fmt.Errorf("unknown header type %q", name)
	}
	if strings.Contains(tmpl, "%s") {
		return fmt.Sprintf(tmpl, o.order()), nil
	}
	return tmpl, nil
}

var scalars = map[string]string{
	"bool":    "parco.Bool()",
	"byte":    "parco.Byte()",
	"uint8":   "parco.UInt8()",
	"int8":    "parco.Int8()",
	"uint16":  "parco.UInt16%s()",
	"int16":   "parco.Int16%s()",
	"uint32":  "parco.UInt32%s()",
	"int32":   "parco.Int32%s()",
	"uint64":  "parco.UInt64%s()",
	"int64":   "parco.Int64%s()",
	"uint":    "parco.UInt%s()",
	"int":     "parco.Int%s()",
	"float32": "parco.Float32%s()",
	"float64": "parco.Float64%s()",
}

func (s *source) exprText(expr ast.Expr) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, s.fset, expr); err != nil {
		panic(err)
	}
	return buf.String()
}

// typeExpr returns the parco.Type expression for a Go type, or an error when
// parco-gen has no mapping for it.
func (s *source) typeExpr(expr ast.Expr, opts tagOptions) (string, error) {
	switch t := expr.(type) {
	case *ast.Ident:
		if tmpl, ok := scalars[t.Name]; ok {
			if strings.Contains(tmpl, "%s") {
				return fmt.Sprintf(tmpl, opts.order()), nil
			}
			return tmpl, nil
		}
		if t.Name == "string" {
			switch opts.header {
			case "":
				return "parco.Varchar()", nil
			case "uint8":
				return "parco.SmallVarchar()", nil
			}
			hdr, err := opts.headerExpr()
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("parco.String(%s)", hdr), nil
		}
		if _, ok := s.structs[t.Name]; ok {
			return modelFunc(t.Name) + "()", nil
		}
		return "", fmt.Errorf("unsupported type %s", t.Name)
	case *ast.SelectorExpr:
		if s.exprText(t) == "time.Time" {
			if opts.location {
				return "parco.TimeLocation()", nil
			}
			return "parco.TimeUTC()", nil
		}
		return "", fmt.Errorf("unsupported type %s", s.exprText(t))
	case *ast.StarExpr:
		inner, err := s.typeExpr(t.X, opts)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("parco.Option(%s)", inner), nil
	case *ast.ArrayType:
		if t.Len == nil && s.exprText(t.Elt) == "byte" {
			hdr, err := opts.headerExpr()
			if err != nil {
				return "", err
			}
			return fmt.Sprintf("parco.Blob(%s)", hdr), nil
		}
		elem, err := s.typeExpr(t.Elt, opts)
		if err != nil {
			return "", err
		}
		if t.Len != nil {
			return fmt.Sprintf("parco.Array(%s, %s)", s.exprText(t.Len), elem), nil
		}
		hdr, err := opts.headerExpr()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("parco.Slice(%s, %s)", hdr, elem), nil
	case *ast.MapType:
		key, err := s.typeExpr(t.Key, opts)
		if err != nil {
			return "", err
		}
		value, err := s.typeExpr(t.Value, opts)
		if err != nil {
			return "", err
		}
		hdr, err := opts.headerExpr()
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("parco.MapType(%s, %s, %s)", hdr, key, value), nil
	}
	return "", fmt.Errorf("unsupported type %s", s.exprText(expr))
}

func (s *source) wireOf(name string, expr ast.Expr, opts tagOptions) (wire, error) {
	typeExpr, err := s.typeExpr(expr, opts)
	if err != nil {
		return wire{}, err
	}
	goType := s.exprText(expr)
	w := wire{
		typeExpr: typeExpr,
		valueTyp: goType,
		get:      fmt.Sprintf("return x.%s", name),
		set:      fmt.Sprintf("x.%s = v", name),
	}
	arr, isArray := expr.(*ast.ArrayType)
	if !isArray || s.exprText(arr.Elt) == "byte" && arr.Len == nil {
		return w, nil
	}
	elem := s.exprText(arr.Elt)
	w.valueTyp = fmt.Sprintf("parco.Iterable[%s]", elem)
	if arr.Len == nil {
		w.get = fmt.Sprintf("return parco.SliceView[%s](x.%s)", elem, name)
		w.set = fmt.Sprintf("x.%s = v.Unwrap()", name)
		return w, nil
	}
	w.get = fmt.Sprintf("return parco.SliceView[%s](x.%s[:])", elem, name)
	w.set = fmt.Sprintf("copy(x.%s[:], v.Unwrap())", name)
	return w, nil
}

func modelFunc(typeName string) string {
	if token.IsExported(typeName) {
		return "New" + typeName + "Model"
	}
	return "new" + strings.ToUpper(typeName[:1]) + typeName[1:] + "Model"
}
