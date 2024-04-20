package iofundamentals

import (
	"testing"
)

// https://medium.com/@andreiboar/fundamentals-of-i-o-in-go-c893d3714deb

func Test_tryFmtPrintF(t *testing.T) {
	in := []byte("Hello, World!")
	out := tryFmtPrintF(in)
	if string(in) != string(out) {
		t.Errorf("tryIoReadAll() = %v, want %v", out, in)
	}
}

func Test_tryIoCopy(t *testing.T) {
	in := []byte("Hello, World!")
	out := tryIoCopy(in)
	if string(in) != string(out) {
		t.Errorf("tryIoReadAll() = %v, want %v", out, in)
	}
}

func Test_tryIoReadAll(t *testing.T) {
	in := []byte("Hello, World!")
	out := tryIoReadAll(in)
	if string(in) != string(out) {
		t.Errorf("tryIoReadAll() = %v, want %v", out, in)
	}
}

// https://medium.com/@andreiboar/fundamentals-of-i-o-in-go-part-2-e7bb68cd5608

func Test_tryIoLimitReader(t *testing.T) {
	in := []byte("Hello, World!")
	out := tryIoLimitReader(in)
	want := "Hello"
	if string(out) != string(want) {
		t.Errorf("tryIoReadAll() = %s, want %s", out, want)
	}
}
