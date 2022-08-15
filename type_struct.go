package parco

type (
	StructType[T any] struct {
		ParserType[T]
		CompilerType[T]
	}
)

func (s StructType[T]) ByteLength() int {
	panic("not implemented")
}

func Struct[T any](b ModelBuilder[T]) StructType[T] {
	parser, compiler := b.ParCo()
	return StructParco[T](parser, compiler)
}

func StructPar[T any](parser ParserType[T]) StructType[T] {
	return StructType[T]{
		ParserType: parser,
	}
}

func StructCo[T any](compiler CompilerType[T]) StructType[T] {
	return StructType[T]{
		CompilerType: compiler,
	}
}

func StructParco[T any](parser ParserType[T], compiler CompilerType[T]) StructType[T] {
	return StructType[T]{
		ParserType:   parser,
		CompilerType: compiler,
	}
}
