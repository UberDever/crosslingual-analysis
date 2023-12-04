
```c
// file.c
#include "some_file.h"

int foo() {
    a += 2;
}
```
```js
// file.js
a = require("some_module")

export function bar() {
    a.foo()
}
```
```
Note: file.js has interlinked dependencies,
but they are not listed here because this is not a tool concern for now

_ (some_file.h): File |- 
    [C:0] file.c: File |-
        _ [a]: Int |- 
            _ [foo]: Unit -> Int

_ (some_module.js): File |-
    [Js:0] file.js: File |-
        _ [a]: Opaque |-
            _ [foo]: Unit -> Any |-
                _ [bar]: Unit -> Int
```
Linear form:
```
_ (some_file.h): File |- [C:0] file.c: File
_ [a]: Int |- _ [foo]: Unit -> Int

_ (some_module.js): File |- [Js:0] file.js: File
_ [a]: Opaque |- _ [foo]: Unit -> Any
_ [foo]: Unit -> Any |- _ [bar]: Unit -> Unit
```