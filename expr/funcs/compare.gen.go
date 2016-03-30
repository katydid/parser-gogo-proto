// Code generated by funcs-gen.
// DO NOT EDIT!

package funcs

import (
	"bytes"
)

type doubleGE struct {
	V1 Double
	V2 Double
}

func (this *doubleGE) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 >= v2, nil
}

func init() {
	Register("ge", new(doubleGE))
}

// DoubleGE returns a new greater than or equal function.
func DoubleGE(a, b Double) Bool {
	return &doubleGE{V1: a, V2: b}
}

type intGE struct {
	V1 Int
	V2 Int
}

func (this *intGE) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 >= v2, nil
}

func init() {
	Register("ge", new(intGE))
}

// IntGE returns a new greater than or equal function.
func IntGE(a, b Int) Bool {
	return &intGE{V1: a, V2: b}
}

type uintGE struct {
	V1 Uint
	V2 Uint
}

func (this *uintGE) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 >= v2, nil
}

func init() {
	Register("ge", new(uintGE))
}

// UintGE returns a new greater than or equal function.
func UintGE(a, b Uint) Bool {
	return &uintGE{V1: a, V2: b}
}

type bytesGE struct {
	V1 Bytes
	V2 Bytes
}

func (this *bytesGE) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return bytes.Compare(v1, v2) >= 0, nil
}

func init() {
	Register("ge", new(bytesGE))
}

// BytesGE returns a new greater than or equal function.
func BytesGE(a, b Bytes) Bool {
	return &bytesGE{V1: a, V2: b}
}

type doubleGt struct {
	V1 Double
	V2 Double
}

func (this *doubleGt) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 > v2, nil
}

func init() {
	Register("gt", new(doubleGt))
}

// DoubleGt returns a new greater than function.
func DoubleGt(a, b Double) Bool {
	return &doubleGt{V1: a, V2: b}
}

type intGt struct {
	V1 Int
	V2 Int
}

func (this *intGt) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 > v2, nil
}

func init() {
	Register("gt", new(intGt))
}

// IntGt returns a new greater than function.
func IntGt(a, b Int) Bool {
	return &intGt{V1: a, V2: b}
}

type uintGt struct {
	V1 Uint
	V2 Uint
}

func (this *uintGt) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 > v2, nil
}

func init() {
	Register("gt", new(uintGt))
}

// UintGt returns a new greater than function.
func UintGt(a, b Uint) Bool {
	return &uintGt{V1: a, V2: b}
}

type bytesGt struct {
	V1 Bytes
	V2 Bytes
}

func (this *bytesGt) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return bytes.Compare(v1, v2) > 0, nil
}

func init() {
	Register("gt", new(bytesGt))
}

// BytesGt returns a new greater than function.
func BytesGt(a, b Bytes) Bool {
	return &bytesGt{V1: a, V2: b}
}

type doubleLE struct {
	V1 Double
	V2 Double
}

func (this *doubleLE) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 <= v2, nil
}

func init() {
	Register("le", new(doubleLE))
}

// DoubleLE returns a new less than or equal function.
func DoubleLE(a, b Double) Bool {
	return &doubleLE{V1: a, V2: b}
}

type intLE struct {
	V1 Int
	V2 Int
}

func (this *intLE) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 <= v2, nil
}

func init() {
	Register("le", new(intLE))
}

// IntLE returns a new less than or equal function.
func IntLE(a, b Int) Bool {
	return &intLE{V1: a, V2: b}
}

type uintLE struct {
	V1 Uint
	V2 Uint
}

func (this *uintLE) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 <= v2, nil
}

func init() {
	Register("le", new(uintLE))
}

// UintLE returns a new less than or equal function.
func UintLE(a, b Uint) Bool {
	return &uintLE{V1: a, V2: b}
}

type bytesLE struct {
	V1 Bytes
	V2 Bytes
}

func (this *bytesLE) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return bytes.Compare(v1, v2) <= 0, nil
}

func init() {
	Register("le", new(bytesLE))
}

// BytesLE returns a new less than or equal function.
func BytesLE(a, b Bytes) Bool {
	return &bytesLE{V1: a, V2: b}
}

type doubleLt struct {
	V1 Double
	V2 Double
}

func (this *doubleLt) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 < v2, nil
}

func init() {
	Register("lt", new(doubleLt))
}

// DoubleLt returns a new less than function.
func DoubleLt(a, b Double) Bool {
	return &doubleLt{V1: a, V2: b}
}

type intLt struct {
	V1 Int
	V2 Int
}

func (this *intLt) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 < v2, nil
}

func init() {
	Register("lt", new(intLt))
}

// IntLt returns a new less than function.
func IntLt(a, b Int) Bool {
	return &intLt{V1: a, V2: b}
}

type uintLt struct {
	V1 Uint
	V2 Uint
}

func (this *uintLt) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 < v2, nil
}

func init() {
	Register("lt", new(uintLt))
}

// UintLt returns a new less than function.
func UintLt(a, b Uint) Bool {
	return &uintLt{V1: a, V2: b}
}

type bytesLt struct {
	V1 Bytes
	V2 Bytes
}

func (this *bytesLt) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return bytes.Compare(v1, v2) < 0, nil
}

func init() {
	Register("lt", new(bytesLt))
}

// BytesLt returns a new less than function.
func BytesLt(a, b Bytes) Bool {
	return &bytesLt{V1: a, V2: b}
}

type doubleEq struct {
	V1 Double
	V2 Double
}

func (this *doubleEq) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 == v2, nil
}

func init() {
	Register("eq", new(doubleEq))
}

// DoubleEq returns a new equal function.
func DoubleEq(a, b Double) Bool {
	return &doubleEq{V1: a, V2: b}
}

type intEq struct {
	V1 Int
	V2 Int
}

func (this *intEq) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 == v2, nil
}

func init() {
	Register("eq", new(intEq))
}

// IntEq returns a new equal function.
func IntEq(a, b Int) Bool {
	return &intEq{V1: a, V2: b}
}

type uintEq struct {
	V1 Uint
	V2 Uint
}

func (this *uintEq) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 == v2, nil
}

func init() {
	Register("eq", new(uintEq))
}

// UintEq returns a new equal function.
func UintEq(a, b Uint) Bool {
	return &uintEq{V1: a, V2: b}
}

type boolEq struct {
	V1 Bool
	V2 Bool
}

func (this *boolEq) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 == v2, nil
}

func init() {
	Register("eq", new(boolEq))
}

// BoolEq returns a new equal function.
func BoolEq(a, b Bool) Bool {
	return &boolEq{V1: a, V2: b}
}

type stringEq struct {
	V1 String
	V2 String
}

func (this *stringEq) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 == v2, nil
}

func init() {
	Register("eq", new(stringEq))
}

// StringEq returns a new equal function.
func StringEq(a, b String) Bool {
	return &stringEq{V1: a, V2: b}
}

type bytesEq struct {
	V1 Bytes
	V2 Bytes
}

func (this *bytesEq) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return bytes.Equal(v1, v2), nil
}

func init() {
	Register("eq", new(bytesEq))
}

// BytesEq returns a new equal function.
func BytesEq(a, b Bytes) Bool {
	return &bytesEq{V1: a, V2: b}
}

type doubleNe struct {
	V1 Double
	V2 Double
}

func (this *doubleNe) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 != v2, nil
}

func init() {
	Register("ne", new(doubleNe))
}

// DoubleNe returns a new not equal function.
func DoubleNe(a, b Double) Bool {
	return &doubleNe{V1: a, V2: b}
}

type intNe struct {
	V1 Int
	V2 Int
}

func (this *intNe) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 != v2, nil
}

func init() {
	Register("ne", new(intNe))
}

// IntNe returns a new not equal function.
func IntNe(a, b Int) Bool {
	return &intNe{V1: a, V2: b}
}

type uintNe struct {
	V1 Uint
	V2 Uint
}

func (this *uintNe) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 != v2, nil
}

func init() {
	Register("ne", new(uintNe))
}

// UintNe returns a new not equal function.
func UintNe(a, b Uint) Bool {
	return &uintNe{V1: a, V2: b}
}

type boolNe struct {
	V1 Bool
	V2 Bool
}

func (this *boolNe) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 != v2, nil
}

func init() {
	Register("ne", new(boolNe))
}

// BoolNe returns a new not equal function.
func BoolNe(a, b Bool) Bool {
	return &boolNe{V1: a, V2: b}
}

type stringNe struct {
	V1 String
	V2 String
}

func (this *stringNe) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return v1 != v2, nil
}

func init() {
	Register("ne", new(stringNe))
}

// StringNe returns a new not equal function.
func StringNe(a, b String) Bool {
	return &stringNe{V1: a, V2: b}
}

type bytesNe struct {
	V1 Bytes
	V2 Bytes
}

func (this *bytesNe) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return false, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return false, nil
	}
	return !bytes.Equal(v1, v2), nil
}

func init() {
	Register("ne", new(bytesNe))
}

// BytesNe returns a new not equal function.
func BytesNe(a, b Bytes) Bool {
	return &bytesNe{V1: a, V2: b}
}
