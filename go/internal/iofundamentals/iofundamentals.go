package iofundamentals

import (
	"bytes"
	"fmt"
	"io"
)

func tryFmtPrintF(in []byte) []byte {
	s := string(in)
	var buff bytes.Buffer // bytes.Buffer implements io.ReadWriter
	fmt.Fprintf(&buff, s)
	out := buff.Bytes()
	return out
}

func tryIoCopy(in []byte) []byte {
	src := bytes.NewReader(in) // bytes.Reader implements io.Reader
	var dst bytes.Buffer
	io.Copy(&dst, src)
	out := dst.Bytes()
	return out
}

func tryIoReadAll(in []byte) []byte {
	r := bytes.NewReader(in)
	out, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return out
}

func tryIoLimitReader(in []byte) []byte {
	r := bytes.NewReader(in)
	lr := io.LimitReader(r, 5)
	out, err := io.ReadAll(lr)
	if err != nil {
		panic(err)
	}
	return out
}
