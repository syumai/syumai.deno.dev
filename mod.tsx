/* @jsx h */
import {
  h,
  renderHTML,
} from "https://denopkg.com/syumai/deno-libs/jsx/renderer.ts";

import { listenAndServe } from "https://deno.land/std@0.111.0/http/server.ts";

const Body = () => (
  <body>
    <h1>Hello, world of Deno Deploy!</h1>
    <div>
      <img src="https://syum.ai/image/random" />
    </div>
    <div>
      <a href="https://github.com/syumai/deno-deploy-example">GitHub Repo</a>
    </div>
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

listenAndServe(":8080", () => (
  new Response(renderHTML(html), {
    status: 200,
    headers: {
      server: "denosr",
      "content-type": "text/html",
    },
  })
));
