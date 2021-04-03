//go:generate sh -c "tinygo build -opt=s -o main.wasm -target wasm ./ && cat main.wasm | deno run https://denopkg.com/syumai/binpack/mod.ts > mainwasm.ts && rm main.wasm"
package main

import (
	"bytes"
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	_ "image/png"
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

	global.Set("scaleImage",
		js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
			if len(args) < 2 {
				panic("two args must be given")
			}
			f := NewDenoFile(args[0])
			rd, err := scaleImage(f, args[1].Float())
			if err != nil {
				panic(err)
			}
			return NewDenoJSReader(rd)
		}))
	select {}
}

func scaleImage(rd io.Reader, scale float64) (io.Reader, error) {
	m, t, err := image.Decode(rd)
	if err != nil {
		return nil, err
	}
	b := m.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, int(float64(b.Dx()) * scale), int(float64(b.Dy()) * scale)))
	draw.BiLinear.Scale(dst, dst.Bounds(), m, m.Bounds(), draw.Over, nil)

	var buf bytes.Buffer
	switch t {
	case "jpeg":
		if err := jpeg.Encode(&buf, dst, &jpeg.Options{Quality: 100}); err != nil {
			return nil, err
		}
	case "gif":
		if err := gif.Encode(&buf, dst, nil); err != nil {
			return nil, err
		}
	case "png":
		if err := png.Encode(&buf, dst); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown format: %s", t)
	}
	return &buf, nil
}

func printUppercased(rd io.Reader) {
	b, err := io.ReadAll(rd)
	if err != nil {
		panic(err)
	}
	fmt.Printf(strings.ToUpper(string(b)))
}

func NewUint8Array(size int) js.Value {
	ua := global.Get("Uint8Array")
	return ua.New(size)
}
