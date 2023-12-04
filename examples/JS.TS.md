```ts
//file.ts
export default function (a: number) => (b: number) => plus(a, b) // assuming plus is declared in stdlib and opaque/implicit

// file.js
f = require("file.ts")
let a = f(1, 2)

This example proves that all entities, that interact in code are always either:
 - named
 - nested/arranged in some pattern 

[Ts:0] file.ts: Int -> Int -> Any

[Ts:0] file.ts: File |- 
    [Js:0] file.js: ()
```
