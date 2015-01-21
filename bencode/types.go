package bencode

type Type int

const (
	String Type = iota
	List
	Dictionary
	Integer
)

type BType interface {
	Typ() Type
	Decode() []byte
}

type String string

func (s String) Typ() { return String }

func (s String) Decode() []byte {
	return []byte(fmt.Sprintf("%d:%s", len(s), string(s)))
}

type List []Btype

func (l List) Type() { return List }

func (l List) Decode() []byte {
	out := []byte{'l'}
	for _, e := range List {
		out = append(out, e.Decode())
	}
	return append(out, []byte{'e'})
}



