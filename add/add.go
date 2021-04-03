//go:generate sh -c "tinygo build -opt=s -o add.wasm -target wasm add.go && cat add.wasm | deno run https://denopkg.com/syumai/binpack/mod.ts > addwasm.ts && rm add.wasm"
package main

func main() {}

//export add
func add(a, b int32) int32 {
	return a + b
}
