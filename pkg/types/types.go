package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const BinaryType uint8 = iota + 1

type Payload interface {
	io.ReaderFrom
	io.WriterTo
	Bytes() Binary
}

type Binary []byte

func (m Binary) Bytes() Binary {
	return m
}

func (m Binary) WriteTo(w io.Writer) (int64, error) {
	err := binary.Write(w, binary.BigEndian, BinaryType)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	err = binary.Write(w, binary.BigEndian, int32(len(m)))
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	n, err := w.Write(m)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return int64(n + 5), err
}

func (m *Binary) ReadFrom(r io.Reader) (int64, error) {
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var size uint32
	err = binary.Read(r, binary.BigEndian, &size)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	*m = make(Binary, size)
	n, err := r.Read(*m)
	return int64(n + 5), err
}

func Decode(r io.Reader) (Payload, error) {
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	payload := new(Binary)
	_, err = payload.ReadFrom(io.MultiReader(bytes.NewReader(Binary{typ}), r))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return payload, err
}
