import addwasm from "./addwasm.ts";
import { decode } from "../vendor/https/deno.land/x/std/encoding/base64.ts";
import { Go } from "./wasm_exec.js";

const bytes = decode(addwasm);
const go = new Go();
const result = await WebAssembly.instantiate(bytes, go.importObject);
const add = result.instance.exports.add as (a: number, b: number) => number;

console.log(add(1, 2));
