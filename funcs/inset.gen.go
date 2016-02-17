// Code generated by funcs-gen.
// DO NOT EDIT!

package funcs

type inSetInt struct {
	Elem Int
	List ConstInts
	set  map[int64]struct{}
}

func (this *inSetInt) Init() error {
	l, err := this.List.Eval()
	if err != nil {
		return err
	}
	this.set = make(map[int64]struct{})
	for i := range l {
		this.set[l[i]] = struct{}{}
	}
	return nil
}

func (this *inSetInt) Eval() (bool, error) {
	v, err := this.Elem.Eval()
	if err != nil {
		return false, err
	}
	_, ok := this.set[v]
	return ok, nil
}

func init() {
	Register("contains", new(inSetInt))
}

func ContainsInt(e Int, l ConstInts) Bool {
	return &inSetInt{e, l, nil}
}

type inSetUint struct {
	Elem Uint
	List ConstUints
	set  map[uint64]struct{}
}

func (this *inSetUint) Init() error {
	l, err := this.List.Eval()
	if err != nil {
		return err
	}
	this.set = make(map[uint64]struct{})
	for i := range l {
		this.set[l[i]] = struct{}{}
	}
	return nil
}

func (this *inSetUint) Eval() (bool, error) {
	v, err := this.Elem.Eval()
	if err != nil {
		return false, err
	}
	_, ok := this.set[v]
	return ok, nil
}

func init() {
	Register("contains", new(inSetUint))
}

func ContainsUint(e Uint, l ConstUints) Bool {
	return &inSetUint{e, l, nil}
}

type inSetString struct {
	Elem String
	List ConstStrings
	set  map[string]struct{}
}

func (this *inSetString) Init() error {
	l, err := this.List.Eval()
	if err != nil {
		return err
	}
	this.set = make(map[string]struct{})
	for i := range l {
		this.set[l[i]] = struct{}{}
	}
	return nil
}

func (this *inSetString) Eval() (bool, error) {
	v, err := this.Elem.Eval()
	if err != nil {
		return false, err
	}
	_, ok := this.set[v]
	return ok, nil
}

func init() {
	Register("contains", new(inSetString))
}

func ContainsString(e String, l ConstStrings) Bool {
	return &inSetString{e, l, nil}
}