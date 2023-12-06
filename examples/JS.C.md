
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
Results:
```yaml

Fragments:
    [some_file.h]: File
    [foo:file.c]: Unit -> Int
    [a:file.c]: Int
    [some_module.js]: File
    [a:file.js]: Opaque
    [bar]: Unit -> Unit
    [foo:file.js]: Unit -> Unit
    [file.c]: File
    [file.js]: File
    
Scope:
    [some_file.h]: File
    [file.c]: File
    [file.js]: File
    [some_module.js]: File
    [foo:file.c]: Unit -> Int
    [bar]: Unit -> Unit
        [a:file.c]: Int
        --
        [a:file.js]: Opaque
        [a.foo]: Unit -> Unit

Links:
    [bar]: Unit -> Unit :-
        [a.foo]: Unit -> Unit :-
            [a:file.js]: Opaque :-
                [some_module.js]: File
    [a:file.js]: Opaque :- 
        [file.js]: File :-
            [some_module.js]: File
    [foo:file.c]: Unit -> Int :- 
        [a:file.c]: Int :-
            [file.c]: File :-
                [some_file.h]: File
   
```