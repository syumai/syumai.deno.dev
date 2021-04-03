package main

import (
	"io"
	"syscall/js"
)

type ReadWriteSeekCloser interface {
	io.Reader
	io.Writer
	io.Closer
	io.Seeker
}

type DenoFile struct {
	fileValue js.Value
}

var _ ReadWriteSeekCloser = (*DenoFile)(nil)

func NewDenoFile(v js.Value) *DenoFile {
	return &DenoFile{v}
}

func (f *DenoFile) Read(p []byte) (int, error) {
	ua := NewUint8Array(len(p))
	result := f.fileValue.Call("readSync", ua)
	if result.IsNull() {
		return 0, io.EOF
	}
	_ = js.CopyBytesToGo(p, ua)
	return result.Int(), nil
}

func (f *DenoFile) Write(p []byte) (int, error) {
	ua := NewUint8Array(len(p))
	_ = js.CopyBytesToJS(ua, p)
	result := f.fileValue.Call("writeSync", ua)
	return result.Int(), nil
}

// Seek
// whence: SeekStart = 0 / SeekCurrent = 1 / SeekEnd = 2
func (f *DenoFile) Seek(offset int64, whence int) (int64, error) {
	result := f.fileValue.Call("seekSync", js.ValueOf(offset), js.ValueOf(whence))
	return int64(result.Int()), nil
}

func (f *DenoFile) Close() error {
	f.fileValue.Call("close")
	return nil
}
