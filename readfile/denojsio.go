package main

import (
	"io"
	"syscall/js"
)

type DenoJSReader struct {
	rd io.Reader
}

func NewDenoJSReader(rd io.Reader) js.Value {
	djr := &DenoJSReader{ rd }
	readFunc := js.FuncOf(func (_ js.Value, args []js.Value) interface{} {
		return djr.Read(args[0])
	})
	obj := map[string]interface{}{
		"read": readFunc,
	}
	return js.ValueOf(obj)
}

func (djr *DenoJSReader) Read(ua js.Value) js.Value {
	b := make([]byte, ua.Length())
	n, err := djr.rd.Read(b)
	if err == io.EOF {
		return js.Null()
	}
	if err != nil {
		panic(err)
	}
	_ = js.CopyBytesToJS(ua, b)
	return js.ValueOf(n)
}