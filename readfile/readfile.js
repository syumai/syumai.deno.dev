import readfilewasm from "./readfilewasm.ts";
import { decode } from "../vendor/https/deno.land/x/std/encoding/base64.ts";
import { Go } from "../wasm_exec.js";

const bytes = decode(readfilewasm);
const go = new Go();
const result = await WebAssembly.instantiate(bytes, go.importObject);
go.run(result.instance);

const f = await Deno.open("./example.txt");

printUppercased(f);

