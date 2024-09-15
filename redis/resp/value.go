package resp

import "strconv"

const (
	STRING = '+'
	ERROR  = '-'
	BULK   = '$'
	ARRAY  = '*'
	NULL   = 'n'
)

type Value struct {
	Typ   rune
	Str   string
	Num   int
	Bulk  string
	Array []Value
}

func (v Value) Marshal() []byte {
	switch v.Typ {
	case ARRAY:
		return v.marshalArray()
	case BULK:
		return v.marshalBulk()
	case STRING:
		return v.marshalString()
	case NULL:
		return v.marshalNull()
	case ERROR:
		return v.marshalError()
	default:
		return []byte{}
	}
}

func (v Value) marshalArray() (bytes []byte) {
	length := len(v.Array)
	bytes = append(bytes, ARRAY)
	bytes = append(bytes, strconv.Itoa(length)...)
	bytes = append(bytes, '\r', '\n')
	for _, elem := range v.Array {
		bytes = append(bytes, elem.Marshal()...)
	}
	return bytes
}

func (v Value) marshalBulk() (bytes []byte) {
	bytes = append(bytes, BULK)
	bytes = append(bytes, strconv.Itoa(len(v.Bulk))...)
	bytes = append(bytes, '\r', '\n')
	bytes = append(bytes, v.Bulk...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) marshalString() (bytes []byte) {
	bytes = append(bytes, STRING)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) marshalError() (bytes []byte) {
	bytes = append(bytes, ERROR)
	bytes = append(bytes, v.Str...)
	bytes = append(bytes, '\r', '\n')
	return bytes
}

func (v Value) marshalNull() []byte {
	return []byte("$-1\r\n")
}
