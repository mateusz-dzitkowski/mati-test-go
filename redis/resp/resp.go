package resp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Resp struct {
	reader *bufio.Reader
	writer io.Writer
}

func NewResp(reader io.Reader, writer io.Writer) *Resp {
	return &Resp{
		reader: bufio.NewReader(reader),
		writer: writer,
	}
}

func (r *Resp) Write(v Value) error {
	bytes := v.Marshal()
	_, err := r.writer.Write(bytes)
	return err
}

func (r *Resp) Read() (Value, error) {
	_type, err := r.reader.ReadByte()
	if err != nil {
		return Value{}, err
	}

	switch _type {
	case ARRAY:
		return r.readArray()
	case BULK:
		return r.readBulk()
	default:
		return Value{}, errors.New(fmt.Sprintf("unknown type: %v", _type))
	}
}

func (r *Resp) readLine() (line []byte, n int, err error) {
	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}
		n += 1
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}
	return line[:len(line)-2], n, nil
}

func (r *Resp) readInteger() (output int, n int, err error) {
	line, n, err := r.readLine()
	if err != nil {
		return 0, 0, err
	}
	i64output, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return int(i64output), n, nil
}

func (r *Resp) readArray() (Value, error) {
	v := Value{}
	v.Typ = ARRAY
	length, _, err := r.readLine()
	if err != nil {
		return v, err
	}
	intLength, err := strconv.Atoi(string(length))
	if err != nil {
		return v, err
	}

	v.Array = make([]Value, 0)
	for range intLength {
		val, err := r.Read()
		if err != nil {
			return v, err
		}
		v.Array = append(v.Array, val)
	}

	return v, nil
}

func (r *Resp) readBulk() (Value, error) {
	v := Value{}
	v.Typ = BULK

	length, _, err := r.readInteger()
	if err != nil {
		return v, err
	}

	bulk := make([]byte, length)
	_, err = r.reader.Read(bulk)
	if err != nil {
		return v, err
	}

	v.Bulk = string(bulk)

	// read the trailing CRLF
	_, _, _ = r.readLine()
	return v, nil
}
