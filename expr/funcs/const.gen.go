// Code generated by funcs-gen.
// DO NOT EDIT!

package funcs

import (
	"fmt"
	"reflect"
	"strings"
)

type ConstDouble interface {
	Double
}

var typConstDouble reflect.Type = reflect.TypeOf((*ConstDouble)(nil)).Elem()

type constDouble struct {
	v float64
}

//DoubleConst returns a new constant function of type Double
func DoubleConst(v float64) ConstDouble {
	return &constDouble{v}
}

func (this *constDouble) IsConst() {}

func (this *constDouble) Eval() (float64, error) {
	return this.v, nil
}

func (this *constDouble) String() string {
	return fmt.Sprintf("double(%f)", this.v)
}

type ConstInt interface {
	Int
}

var typConstInt reflect.Type = reflect.TypeOf((*ConstInt)(nil)).Elem()

type constInt struct {
	v int64
}

//IntConst returns a new constant function of type Int
func IntConst(v int64) ConstInt {
	return &constInt{v}
}

func (this *constInt) IsConst() {}

func (this *constInt) Eval() (int64, error) {
	return this.v, nil
}

func (this *constInt) String() string {
	return fmt.Sprintf("int(%d)", this.v)
}

type ConstUint interface {
	Uint
}

var typConstUint reflect.Type = reflect.TypeOf((*ConstUint)(nil)).Elem()

type constUint struct {
	v uint64
}

//UintConst returns a new constant function of type Uint
func UintConst(v uint64) ConstUint {
	return &constUint{v}
}

func (this *constUint) IsConst() {}

func (this *constUint) Eval() (uint64, error) {
	return this.v, nil
}

func (this *constUint) String() string {
	return fmt.Sprintf("uint(%d)", this.v)
}

type ConstBool interface {
	Bool
}

var typConstBool reflect.Type = reflect.TypeOf((*ConstBool)(nil)).Elem()

type constBool struct {
	v bool
}

//BoolConst returns a new constant function of type Bool
func BoolConst(v bool) ConstBool {
	return &constBool{v}
}

func (this *constBool) IsConst() {}

func (this *constBool) Eval() (bool, error) {
	return this.v, nil
}

func (this *constBool) String() string {
	return fmt.Sprintf("%v", this.v)
}

type ConstString interface {
	String
}

var typConstString reflect.Type = reflect.TypeOf((*ConstString)(nil)).Elem()

type constString struct {
	v string
}

//StringConst returns a new constant function of type String
func StringConst(v string) ConstString {
	return &constString{v}
}

func (this *constString) IsConst() {}

func (this *constString) Eval() (string, error) {
	return this.v, nil
}

func (this *constString) String() string {
	return fmt.Sprintf("`%s`", this.v)
}

type ConstBytes interface {
	Bytes
}

var typConstBytes reflect.Type = reflect.TypeOf((*ConstBytes)(nil)).Elem()

type constBytes struct {
	v []byte
}

//BytesConst returns a new constant function of type Bytes
func BytesConst(v []byte) ConstBytes {
	return &constBytes{v}
}

func (this *constBytes) IsConst() {}

func (this *constBytes) Eval() ([]byte, error) {
	return this.v, nil
}

func (this *constBytes) String() string {
	return fmt.Sprintf("%#v", this.v)
}

type ConstDoubles interface {
	Doubles
}

var typConstDoubles reflect.Type = reflect.TypeOf((*ConstDoubles)(nil)).Elem()

type constDoubles struct {
	v []float64
}

//DoublesConst returns a new constant function of type Doubles
func DoublesConst(v []float64) ConstDoubles {
	return &constDoubles{v}
}

func (this *constDoubles) IsConst() {}

func (this *constDoubles) Eval() ([]float64, error) {
	return this.v, nil
}

func (this *constDoubles) String() string {
	ss := make([]string, len(this.v))
	for i := range this.v {
		ss[i] = fmt.Sprintf("double(%f)", this.v[i])
	}
	return "[]double{" + strings.Join(ss, ",") + "}"
}

type ConstInts interface {
	Ints
}

var typConstInts reflect.Type = reflect.TypeOf((*ConstInts)(nil)).Elem()

type constInts struct {
	v []int64
}

//IntsConst returns a new constant function of type Ints
func IntsConst(v []int64) ConstInts {
	return &constInts{v}
}

func (this *constInts) IsConst() {}

func (this *constInts) Eval() ([]int64, error) {
	return this.v, nil
}

func (this *constInts) String() string {
	ss := make([]string, len(this.v))
	for i := range this.v {
		ss[i] = fmt.Sprintf("int(%d)", this.v[i])
	}
	return "[]int{" + strings.Join(ss, ",") + "}"
}

type ConstUints interface {
	Uints
}

var typConstUints reflect.Type = reflect.TypeOf((*ConstUints)(nil)).Elem()

type constUints struct {
	v []uint64
}

//UintsConst returns a new constant function of type Uints
func UintsConst(v []uint64) ConstUints {
	return &constUints{v}
}

func (this *constUints) IsConst() {}

func (this *constUints) Eval() ([]uint64, error) {
	return this.v, nil
}

func (this *constUints) String() string {
	ss := make([]string, len(this.v))
	for i := range this.v {
		ss[i] = fmt.Sprintf("uint(%d)", this.v[i])
	}
	return "[]uint{" + strings.Join(ss, ",") + "}"
}

type ConstBools interface {
	Bools
}

var typConstBools reflect.Type = reflect.TypeOf((*ConstBools)(nil)).Elem()

type constBools struct {
	v []bool
}

//BoolsConst returns a new constant function of type Bools
func BoolsConst(v []bool) ConstBools {
	return &constBools{v}
}

func (this *constBools) IsConst() {}

func (this *constBools) Eval() ([]bool, error) {
	return this.v, nil
}

func (this *constBools) String() string {
	ss := make([]string, len(this.v))
	for i := range this.v {
		ss[i] = fmt.Sprintf("%v", this.v[i])
	}
	return "[]bool{" + strings.Join(ss, ",") + "}"
}

type ConstStrings interface {
	Strings
}

var typConstStrings reflect.Type = reflect.TypeOf((*ConstStrings)(nil)).Elem()

type constStrings struct {
	v []string
}

//StringsConst returns a new constant function of type Strings
func StringsConst(v []string) ConstStrings {
	return &constStrings{v}
}

func (this *constStrings) IsConst() {}

func (this *constStrings) Eval() ([]string, error) {
	return this.v, nil
}

func (this *constStrings) String() string {
	ss := make([]string, len(this.v))
	for i := range this.v {
		ss[i] = fmt.Sprintf("`%s`", this.v[i])
	}
	return "[]string{" + strings.Join(ss, ",") + "}"
}

type ConstListOfBytes interface {
	ListOfBytes
}

var typConstListOfBytes reflect.Type = reflect.TypeOf((*ConstListOfBytes)(nil)).Elem()

type constListOfBytes struct {
	v [][]byte
}

//ListOfBytesConst returns a new constant function of type ListOfBytes
func ListOfBytesConst(v [][]byte) ConstListOfBytes {
	return &constListOfBytes{v}
}

func (this *constListOfBytes) IsConst() {}

func (this *constListOfBytes) Eval() ([][]byte, error) {
	return this.v, nil
}

func (this *constListOfBytes) String() string {
	ss := make([]string, len(this.v))
	for i := range this.v {
		ss[i] = fmt.Sprintf("%#v", this.v[i])
	}
	return "[][]byte{" + strings.Join(ss, ",") + "}"
}
