/* @jsx h */
import {
  h,
  renderHTML,
} from "https://denopkg.com/syumai/deno-libs/jsx/renderer.ts";

const Body = () => (
  <body>
    <h1>Hello, world of Deno Deploy!</h1>
    <img src="https://syum.ai/image/random" />
  </body>
);

const html = (
  <html>
    <head>
      <title>Hello, world of Deno Deploy!</title>
    </head>
    <Body />
  </html>
);

interface Responder {
  respondWith(res: Response): void;
}

addEventListener("fetch", (event) => {
  const e = (event as unknown) as Responder;
  e.respondWith(
    new Response(renderHTML(html), {
      status: 200,
      headers: {
        server: "denosr",
        "content-type": "text/html",
      },
    })
  );
});
