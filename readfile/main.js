import mainwasm from "./mainwasm.ts";
import { Go } from "../wasm_exec.js";
import { decode } from "https://deno.land/std@0.92.0/encoding/base64.ts";

const bytes = decode(mainwasm);
const go = new Go();
const result = await WebAssembly.instantiate(bytes, go.importObject);
go.run(result.instance);

// run printUppercased
// const f = await Deno.open("./example.txt");
// printUppercased(f);
// f.close()

// run scaleImage
const img = await Deno.open("./syumai.png");
const scaled = scaleImage(img, 2);
Deno.copy(scaled, Deno.stdout);
img.close();
