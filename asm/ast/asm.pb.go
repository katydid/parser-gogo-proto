// Code generated by protoc-gen-gogo.
// source: asm.proto
// DO NOT EDIT!

/*
Package ast is a generated protocol buffer package.

It is generated from these files:
	asm.proto

It has these top-level messages:
	Rules
	Rule
	Root
	Init
	Transition
	IfExpr
	StateExpr
	Expr
	List
	Function
	Terminal
	Variable
	Keyword
	Space
*/
package ast

import proto "code.google.com/p/gogoprotobuf/proto"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"
import types "github.com/awalterschulze/katydid/types"

import fmt "fmt"
import strings "strings"
import code_google_com_p_gogoprotobuf_proto "code.google.com/p/gogoprotobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Rules struct {
	Rules            []*Rule `protobuf:"bytes,1,rep" json:"Rules,omitempty"`
	Final            *Space  `protobuf:"bytes,2,opt" json:"Final,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Rules) Reset()      { *m = Rules{} }
func (*Rules) ProtoMessage() {}

func (m *Rules) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *Rules) GetFinal() *Space {
	if m != nil {
		return m.Final
	}
	return nil
}

type Rule struct {
	Root             *Root       `protobuf:"bytes,1,opt" json:"Root,omitempty"`
	Init             *Init       `protobuf:"bytes,2,opt" json:"Init,omitempty"`
	Transition       *Transition `protobuf:"bytes,3,opt" json:"Transition,omitempty"`
	IfExpr           *IfExpr     `protobuf:"bytes,4,opt" json:"IfExpr,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Rule) Reset()      { *m = Rule{} }
func (*Rule) ProtoMessage() {}

func (m *Rule) GetRoot() *Root {
	if m != nil {
		return m.Root
	}
	return nil
}

func (m *Rule) GetInit() *Init {
	if m != nil {
		return m.Init
	}
	return nil
}

func (m *Rule) GetTransition() *Transition {
	if m != nil {
		return m.Transition
	}
	return nil
}

func (m *Rule) GetIfExpr() *IfExpr {
	if m != nil {
		return m.IfExpr
	}
	return nil
}

type Root struct {
	Before           *Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	Equal            *Keyword `protobuf:"bytes,2,opt" json:"Equal,omitempty"`
	BeforeQualId     *Space   `protobuf:"bytes,3,opt" json:"BeforeQualId,omitempty"`
	Package          string   `protobuf:"bytes,4,opt" json:"Package"`
	Message          string   `protobuf:"bytes,5,opt" json:"Message"`
	State            string   `protobuf:"bytes,6,opt" json:"State"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Root) Reset()      { *m = Root{} }
func (*Root) ProtoMessage() {}

func (m *Root) GetBefore() *Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *Root) GetEqual() *Keyword {
	if m != nil {
		return m.Equal
	}
	return nil
}

func (m *Root) GetBeforeQualId() *Space {
	if m != nil {
		return m.BeforeQualId
	}
	return nil
}

func (m *Root) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *Root) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Root) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type Init struct {
	Before           *Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	Package          string   `protobuf:"bytes,2,opt" json:"Package"`
	Message          string   `protobuf:"bytes,3,opt" json:"Message"`
	Equal            *Keyword `protobuf:"bytes,4,opt" json:"Equal,omitempty"`
	BeforeState      *Space   `protobuf:"bytes,5,opt" json:"BeforeState,omitempty"`
	State            string   `protobuf:"bytes,6,opt" json:"State"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Init) Reset()      { *m = Init{} }
func (*Init) ProtoMessage() {}

func (m *Init) GetBefore() *Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *Init) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *Init) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Init) GetEqual() *Keyword {
	if m != nil {
		return m.Equal
	}
	return nil
}

func (m *Init) GetBeforeState() *Space {
	if m != nil {
		return m.BeforeState
	}
	return nil
}

func (m *Init) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type Transition struct {
	Before           *Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	Src              string   `protobuf:"bytes,2,opt" json:"Src"`
	BeforeInput      *Space   `protobuf:"bytes,3,opt" json:"BeforeInput,omitempty"`
	Input            string   `protobuf:"bytes,4,opt" json:"Input"`
	Equal            *Keyword `protobuf:"bytes,5,opt" json:"Equal,omitempty"`
	BeforeDst        *Space   `protobuf:"bytes,6,opt" json:"BeforeDst,omitempty"`
	Dst              string   `protobuf:"bytes,7,opt" json:"Dst"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Transition) Reset()      { *m = Transition{} }
func (*Transition) ProtoMessage() {}

func (m *Transition) GetBefore() *Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *Transition) GetSrc() string {
	if m != nil {
		return m.Src
	}
	return ""
}

func (m *Transition) GetBeforeInput() *Space {
	if m != nil {
		return m.BeforeInput
	}
	return nil
}

func (m *Transition) GetInput() string {
	if m != nil {
		return m.Input
	}
	return ""
}

func (m *Transition) GetEqual() *Keyword {
	if m != nil {
		return m.Equal
	}
	return nil
}

func (m *Transition) GetBeforeDst() *Space {
	if m != nil {
		return m.BeforeDst
	}
	return nil
}

func (m *Transition) GetDst() string {
	if m != nil {
		return m.Dst
	}
	return ""
}

type IfExpr struct {
	Before           *Space     `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	Condition        *Expr      `protobuf:"bytes,2,opt" json:"Condition,omitempty"`
	ThenWord         *Keyword   `protobuf:"bytes,3,opt" json:"ThenWord,omitempty"`
	ThenClause       *StateExpr `protobuf:"bytes,4,opt" json:"ThenClause,omitempty"`
	ElseWord         *Keyword   `protobuf:"bytes,5,opt" json:"ElseWord,omitempty"`
	ElseClause       *StateExpr `protobuf:"bytes,6,opt" json:"ElseClause,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *IfExpr) Reset()      { *m = IfExpr{} }
func (*IfExpr) ProtoMessage() {}

func (m *IfExpr) GetBefore() *Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *IfExpr) GetCondition() *Expr {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *IfExpr) GetThenWord() *Keyword {
	if m != nil {
		return m.ThenWord
	}
	return nil
}

func (m *IfExpr) GetThenClause() *StateExpr {
	if m != nil {
		return m.ThenClause
	}
	return nil
}

func (m *IfExpr) GetElseWord() *Keyword {
	if m != nil {
		return m.ElseWord
	}
	return nil
}

func (m *IfExpr) GetElseClause() *StateExpr {
	if m != nil {
		return m.ElseClause
	}
	return nil
}

type StateExpr struct {
	Before           *Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	State            *string  `protobuf:"bytes,2,opt" json:"State,omitempty"`
	IfExpr           *IfExpr  `protobuf:"bytes,3,opt" json:"IfExpr,omitempty"`
	CloseCurly       *Keyword `protobuf:"bytes,4,opt" json:"CloseCurly,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StateExpr) Reset()      { *m = StateExpr{} }
func (*StateExpr) ProtoMessage() {}

func (m *StateExpr) GetBefore() *Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *StateExpr) GetState() string {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ""
}

func (m *StateExpr) GetIfExpr() *IfExpr {
	if m != nil {
		return m.IfExpr
	}
	return nil
}

func (m *StateExpr) GetCloseCurly() *Keyword {
	if m != nil {
		return m.CloseCurly
	}
	return nil
}

type Expr struct {
	Comma            *Keyword  `protobuf:"bytes,1,opt" json:"Comma,omitempty"`
	Terminal         *Terminal `protobuf:"bytes,2,opt" json:"Terminal,omitempty"`
	List             *List     `protobuf:"bytes,3,opt" json:"List,omitempty"`
	Function         *Function `protobuf:"bytes,4,opt" json:"Function,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Expr) Reset()      { *m = Expr{} }
func (*Expr) ProtoMessage() {}

func (m *Expr) GetComma() *Keyword {
	if m != nil {
		return m.Comma
	}
	return nil
}

func (m *Expr) GetTerminal() *Terminal {
	if m != nil {
		return m.Terminal
	}
	return nil
}

func (m *Expr) GetList() *List {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *Expr) GetFunction() *Function {
	if m != nil {
		return m.Function
	}
	return nil
}

type List struct {
	Before           *Space     `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	Type             types.Type `protobuf:"varint,2,opt,enum=types.Type" json:"Type"`
	OpenCurly        *Keyword   `protobuf:"bytes,3,opt" json:"OpenCurly,omitempty"`
	Elems            []*Expr    `protobuf:"bytes,4,rep" json:"Elems,omitempty"`
	CloseCurly       *Keyword   `protobuf:"bytes,5,opt" json:"CloseCurly,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *List) Reset()      { *m = List{} }
func (*List) ProtoMessage() {}

func (m *List) GetBefore() *Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *List) GetType() types.Type {
	if m != nil {
		return m.Type
	}
	return types.UNKNOWN
}

func (m *List) GetOpenCurly() *Keyword {
	if m != nil {
		return m.OpenCurly
	}
	return nil
}

func (m *List) GetElems() []*Expr {
	if m != nil {
		return m.Elems
	}
	return nil
}

func (m *List) GetCloseCurly() *Keyword {
	if m != nil {
		return m.CloseCurly
	}
	return nil
}

type Function struct {
	Before           *Space   `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	Name             string   `protobuf:"bytes,2,opt" json:"Name"`
	OpenParen        *Keyword `protobuf:"bytes,3,opt" json:"OpenParen,omitempty"`
	Params           []*Expr  `protobuf:"bytes,4,rep" json:"Params,omitempty"`
	CloseParen       *Keyword `protobuf:"bytes,5,opt" json:"CloseParen,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Function) Reset()      { *m = Function{} }
func (*Function) ProtoMessage() {}

func (m *Function) GetBefore() *Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *Function) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Function) GetOpenParen() *Keyword {
	if m != nil {
		return m.OpenParen
	}
	return nil
}

func (m *Function) GetParams() []*Expr {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Function) GetCloseParen() *Keyword {
	if m != nil {
		return m.CloseParen
	}
	return nil
}

type Terminal struct {
	Before           *Space    `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	Literal          string    `protobuf:"bytes,2,opt" json:"Literal"`
	DoubleValue      *float64  `protobuf:"fixed64,3,opt" json:"DoubleValue,omitempty"`
	FloatValue       *float32  `protobuf:"fixed32,4,opt" json:"FloatValue,omitempty"`
	Int64Value       *int64    `protobuf:"varint,5,opt" json:"Int64Value,omitempty"`
	Uint64Value      *uint64   `protobuf:"varint,6,opt" json:"Uint64Value,omitempty"`
	Int32Value       *int32    `protobuf:"varint,7,opt" json:"Int32Value,omitempty"`
	BoolValue        *bool     `protobuf:"varint,8,opt" json:"BoolValue,omitempty"`
	StringValue      *string   `protobuf:"bytes,9,opt" json:"StringValue,omitempty"`
	BytesValue       []byte    `protobuf:"bytes,10,opt" json:"BytesValue,omitempty"`
	Uint32Value      *uint32   `protobuf:"varint,11,opt" json:"Uint32Value,omitempty"`
	Variable         *Variable `protobuf:"bytes,50,opt" json:"Variable,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Terminal) Reset()      { *m = Terminal{} }
func (*Terminal) ProtoMessage() {}

func (m *Terminal) GetBefore() *Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *Terminal) GetLiteral() string {
	if m != nil {
		return m.Literal
	}
	return ""
}

func (m *Terminal) GetDoubleValue() float64 {
	if m != nil && m.DoubleValue != nil {
		return *m.DoubleValue
	}
	return 0
}

func (m *Terminal) GetFloatValue() float32 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *Terminal) GetInt64Value() int64 {
	if m != nil && m.Int64Value != nil {
		return *m.Int64Value
	}
	return 0
}

func (m *Terminal) GetUint64Value() uint64 {
	if m != nil && m.Uint64Value != nil {
		return *m.Uint64Value
	}
	return 0
}

func (m *Terminal) GetInt32Value() int32 {
	if m != nil && m.Int32Value != nil {
		return *m.Int32Value
	}
	return 0
}

func (m *Terminal) GetBoolValue() bool {
	if m != nil && m.BoolValue != nil {
		return *m.BoolValue
	}
	return false
}

func (m *Terminal) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Terminal) GetBytesValue() []byte {
	if m != nil {
		return m.BytesValue
	}
	return nil
}

func (m *Terminal) GetUint32Value() uint32 {
	if m != nil && m.Uint32Value != nil {
		return *m.Uint32Value
	}
	return 0
}

func (m *Terminal) GetVariable() *Variable {
	if m != nil {
		return m.Variable
	}
	return nil
}

type Variable struct {
	Name             string     `protobuf:"bytes,1,opt" json:"Name"`
	Type             types.Type `protobuf:"varint,2,opt,enum=types.Type" json:"Type"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Variable) Reset()      { *m = Variable{} }
func (*Variable) ProtoMessage() {}

func (m *Variable) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Variable) GetType() types.Type {
	if m != nil {
		return m.Type
	}
	return types.UNKNOWN
}

type Keyword struct {
	Before           *Space `protobuf:"bytes,1,opt" json:"Before,omitempty"`
	Value            string `protobuf:"bytes,2,opt" json:"Value"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Keyword) Reset()      { *m = Keyword{} }
func (*Keyword) ProtoMessage() {}

func (m *Keyword) GetBefore() *Space {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *Keyword) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Space struct {
	Space            []string `protobuf:"bytes,1,rep" json:"Space,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Space) Reset()      { *m = Space{} }
func (*Space) ProtoMessage() {}

func (m *Space) GetSpace() []string {
	if m != nil {
		return m.Space
	}
	return nil
}

func init() {
}
func (this *Rules) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Rules{` + `Rules:` + fmt.Sprintf("%#v", this.Rules), `Final:` + fmt.Sprintf("%#v", this.Final), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Rule) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Rule{` + `Root:` + fmt.Sprintf("%#v", this.Root), `Init:` + fmt.Sprintf("%#v", this.Init), `Transition:` + fmt.Sprintf("%#v", this.Transition), `IfExpr:` + fmt.Sprintf("%#v", this.IfExpr), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Root) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Root{` + `Before:` + fmt.Sprintf("%#v", this.Before), `Equal:` + fmt.Sprintf("%#v", this.Equal), `BeforeQualId:` + fmt.Sprintf("%#v", this.BeforeQualId), `Package:` + fmt.Sprintf("%#v", this.Package), `Message:` + fmt.Sprintf("%#v", this.Message), `State:` + fmt.Sprintf("%#v", this.State), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Init) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Init{` + `Before:` + fmt.Sprintf("%#v", this.Before), `Package:` + fmt.Sprintf("%#v", this.Package), `Message:` + fmt.Sprintf("%#v", this.Message), `Equal:` + fmt.Sprintf("%#v", this.Equal), `BeforeState:` + fmt.Sprintf("%#v", this.BeforeState), `State:` + fmt.Sprintf("%#v", this.State), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Transition) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Transition{` + `Before:` + fmt.Sprintf("%#v", this.Before), `Src:` + fmt.Sprintf("%#v", this.Src), `BeforeInput:` + fmt.Sprintf("%#v", this.BeforeInput), `Input:` + fmt.Sprintf("%#v", this.Input), `Equal:` + fmt.Sprintf("%#v", this.Equal), `BeforeDst:` + fmt.Sprintf("%#v", this.BeforeDst), `Dst:` + fmt.Sprintf("%#v", this.Dst), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *IfExpr) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.IfExpr{` + `Before:` + fmt.Sprintf("%#v", this.Before), `Condition:` + fmt.Sprintf("%#v", this.Condition), `ThenWord:` + fmt.Sprintf("%#v", this.ThenWord), `ThenClause:` + fmt.Sprintf("%#v", this.ThenClause), `ElseWord:` + fmt.Sprintf("%#v", this.ElseWord), `ElseClause:` + fmt.Sprintf("%#v", this.ElseClause), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *StateExpr) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.StateExpr{` + `Before:` + fmt.Sprintf("%#v", this.Before), `State:` + valueToGoStringAsm(this.State, "string"), `IfExpr:` + fmt.Sprintf("%#v", this.IfExpr), `CloseCurly:` + fmt.Sprintf("%#v", this.CloseCurly), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Expr) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Expr{` + `Comma:` + fmt.Sprintf("%#v", this.Comma), `Terminal:` + fmt.Sprintf("%#v", this.Terminal), `List:` + fmt.Sprintf("%#v", this.List), `Function:` + fmt.Sprintf("%#v", this.Function), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *List) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.List{` + `Before:` + fmt.Sprintf("%#v", this.Before), `Type:` + fmt.Sprintf("%#v", this.Type), `OpenCurly:` + fmt.Sprintf("%#v", this.OpenCurly), `Elems:` + fmt.Sprintf("%#v", this.Elems), `CloseCurly:` + fmt.Sprintf("%#v", this.CloseCurly), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Function) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Function{` + `Before:` + fmt.Sprintf("%#v", this.Before), `Name:` + fmt.Sprintf("%#v", this.Name), `OpenParen:` + fmt.Sprintf("%#v", this.OpenParen), `Params:` + fmt.Sprintf("%#v", this.Params), `CloseParen:` + fmt.Sprintf("%#v", this.CloseParen), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Terminal) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Terminal{` + `Before:` + fmt.Sprintf("%#v", this.Before), `Literal:` + fmt.Sprintf("%#v", this.Literal), `DoubleValue:` + valueToGoStringAsm(this.DoubleValue, "float64"), `FloatValue:` + valueToGoStringAsm(this.FloatValue, "float32"), `Int64Value:` + valueToGoStringAsm(this.Int64Value, "int64"), `Uint64Value:` + valueToGoStringAsm(this.Uint64Value, "uint64"), `Int32Value:` + valueToGoStringAsm(this.Int32Value, "int32"), `BoolValue:` + valueToGoStringAsm(this.BoolValue, "bool"), `StringValue:` + valueToGoStringAsm(this.StringValue, "string"), `BytesValue:` + valueToGoStringAsm(this.BytesValue, "byte"), `Uint32Value:` + valueToGoStringAsm(this.Uint32Value, "uint32"), `Variable:` + fmt.Sprintf("%#v", this.Variable), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Variable) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Variable{` + `Name:` + fmt.Sprintf("%#v", this.Name), `Type:` + fmt.Sprintf("%#v", this.Type), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Keyword) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Keyword{` + `Before:` + fmt.Sprintf("%#v", this.Before), `Value:` + fmt.Sprintf("%#v", this.Value), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func (this *Space) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ast.Space{` + `Space:` + fmt.Sprintf("%#v", this.Space), `XXX_unrecognized:` + fmt.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func valueToGoStringAsm(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringAsm(e map[int32]code_google_com_p_gogoprotobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
