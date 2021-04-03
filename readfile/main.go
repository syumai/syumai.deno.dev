//go:generate sh -c "tinygo build -opt=s -o readfile.wasm -target wasm ./main.go && cat readfile.wasm | deno run https://denopkg.com/syumai/binpack/mod.ts > readfilewasm.ts && rm readfile.wasm"
package main

import (
	"fmt"
	"io"
	"strings"
	"syscall/js"
)

var global = js.Global()

func main() {
	global.Set("printUppercased",
		js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
			f := NewDenoFile(args[0])
			defer f.Close()
			printUppercased(f)
			return nil
		}))
	select {}
}

func printUppercased(rd io.Reader) {
	b, err := io.ReadAll(rd)
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.ToUpper(string(b)))
}

type DenoFile struct {
	fileValue js.Value
}

var _ io.ReadWriteCloser = (*DenoFile)(nil)

func NewDenoFile(v js.Value) *DenoFile {
	return &DenoFile{v}
}

func (f *DenoFile) Read(p []byte) (int, error) {
	ua := NewUint8Array(len(p))
	result := f.fileValue.Call("readSync", ua)
	if result.IsNull() {
		return 0, io.EOF
	}
	n := js.CopyBytesToGo(p, ua)
	return n, nil
}

func (f *DenoFile) Write(p []byte) (int, error) {
	ua := NewUint8Array(len(p))
	_ = js.CopyBytesToJS(ua, p)
	result := f.fileValue.Call("writeSync", ua)
	return result.Int(), nil
}

func (f *DenoFile) Close() error {
	f.fileValue.Call("close")
	return nil
}

func NewUint8Array(size int) js.Value {
	ua := global.Get("Uint8Array")
	return ua.New(size)
}
