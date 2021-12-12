package internal

import (
	"github.com/sonirico/parco/internal/utils"
	"io"
)

type VarcharType struct {
	head Type
	// TODO: IoC parser, typeFactory
}

func (v VarcharType) Length() int {
	return v.head.Length()
}

func (v VarcharType) Compile(x interface{}, w io.Writer) (err error) {
	var bites []byte
	switch val := x.(type) {
	case string:
		bites = utils.String2Bytes(val)
	case []byte:
		bites = val
	}
	// TODO: Check whether header type and actual value do not overflow
	if err = v.head.Compile(len(bites), w); err != nil {
		return err
	}
	_, err = w.Write(bites)
	return err
}

func (v VarcharType) Parse(r io.Reader) (interface{}, error) {
	rawLength, err := v.head.Parse(r)
	if err != nil {
		return nil, err
	}

	le, _ := utils.AnyIntToInt(rawLength)
	data := make([]byte, le)

	if _, err := r.Read(data); err != nil {
		return nil, err
	}

	return utils.Bytes2String(data), nil
}

func Varchar(head Type) VarcharType {
	return VarcharType{head: head}
}

func SmallVarchar() VarcharType {
	return VarcharType{head: UInt8()}
}
