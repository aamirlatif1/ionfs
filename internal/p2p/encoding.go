package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, any) error
}

type GOBDecoder struct{}

func (d GOBDecoder) Decode(r io.Reader, v any) error {
	return gob.NewDecoder(r).Decode(v)
}

type StringDecoder struct{}

func (d StringDecoder) Decode(r io.Reader, v any) error {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		return err
	}
	if strPtr, ok := v.(*string); ok {
		*strPtr = string(buf[:n])
		return nil
	}
	return nil
}
