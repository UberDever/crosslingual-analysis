```ts
//file.ts
export default function (a: number) => (b: number) => plus(a, b) // assuming plus is declared in stdlib and opaque/implicit

// file.js
f = require("file.ts")
let a = f(1, 2)

This example proves that all entities, that interact in code are always either:
 - named
 - nested/arranged in some pattern 
```
Results:
```yaml
Fragments:
    [file.ts]: Int -> Int -> Any | File
    [file.js]: File
    [f]: Opaque | Num -> Num -> Any
    ["file.ts"]: String
    [a]: Any

Scope:
    [file.ts]: Int -> Int -> Any | File
    [file.js]: File
        [f]: Opaque | Num -> Num -> Any
        ["file.ts"]: String
        [a]: Any

Links:
[a]: Any :-
    [f]: Opaque | Num -> Num -> Any :-
        ["file.ts"]: String
        [file.js]: File :- 
            [file.ts]: Int -> Int -> Any | File
```
