```py
# script.py
import ctypes
l = ctypes.CDLL('./liblib.so')
l.doTwoPlusTwo.argtypes = []
l.doTwoPlusTwo.restype = ctypes.c_int

print(l['doTwoPlusTwo']())
```
```c
// lib.c
int doTwoPlusTwo() {
    return 2 + 2
}
```
```sh
rm lib.o 2> /dev/null || rm liblib.so 2> /dev/null
cc -c lib.c
cc -shared -o liblib.so lib.o
python3 script.py
```
Results:
```json
{
    "Identifiers": [
        {"id": 0, "name":"build.sh", "type":"File", "loc":{"line":0, "col":0, "path":""}}, 
        {"id": 1, "name":"lib.c", "type":"File", "loc":{"line":0, "col":0, "path":""}}, 
        {"id": 2, "name":"script.py", "type":"File", "loc":{"line":0, "col":0, "path":""}}, 
        {"id": 3, "name":"ctypes", "type":"Opaque", "loc":{"line":0, "col":0, "path":""}},
        {"id": 4, "name":"doTwoPlusTwo", "type":"() -> Int", "loc":{"line":0, "col":0, "path":"lib.c"}},  
        {"id": 5, "name":"liblib.so", "type":"File", "loc":{"line":0, "col":0, "path":""}}, 
        {"id": 6, "name":"l", "type":"{ doTwoPlusTwo: { argtypes: List Any, restype: Any } } | String -> Unit -> Any", "loc":{"line":0, "col":0, "path":""}}, 
        {"id": 7, "name":"doTwoPlusTwo", "type":"String", "loc":{"line":0, "col":0, "path":"script.py"}},
        {"id": 8, "name":"CDLL", "type":"String -> Any", "loc":{"line":0, "col":0, "path":""}},
        {"id": 9, "name":"lib.o", "type":"File", "loc":{"line":0, "col":0, "path":""}},
    ],
    "Fragments": [
        {"id": -1, "env": [3, 5, 8], "sig":[], "in": "ctypes.CDLL('./liblib.so')", "loc":{"line":0, "col":0, "path":""}},
        {"id": -2, "env": [6, 7], "sig":[], "in": "l['doTwoPlusTwo']", "loc":{"line":0, "col":0, "path":""}},
        {"id": -3, "env": [1], "sig":[], "in": "cc -c lib.c", "loc":{"line":0, "col":0, "path":""}},
        {"id": -4, "env": [9], "sig":[5], "in": "cc -shared -o liblib.so lib.o", "loc":{"line":0, "col":0, "path":""}},
        {"id": -5, "env": [2], "sig":[], "in": "python3 script.py", "loc":{"line":0, "col":0, "path":""}},
        {"id": -6, "env": [0], "sig":[0], "in": "build.sh", "loc":{"line":0, "col":0, "path":""}},
        {"id": -7, "env": [1], "sig":[1], "in": "lib.c", "loc":{"line":0, "col":0, "path":""}},
        {"id": -8, "env": [2], "sig":[2], "in": "script.py", "loc":{"line":0, "col":0, "path":""}},
        {"id": -9, "env": [3], "sig":[], "in": "ctypes", "loc":{"line":0, "col":0, "path":""}},
        {"id": -10, "env": [4], "sig":[4], "in": "doTwoPlusTwo", "loc":{"line":0, "col":0, "path":"lib.c"}},
        {"id": -11, "env": [5], "sig":[], "in": "liblib.so", "loc":{"line":0, "col":0, "path":""}},
        {"id": -12, "env": [6], "sig":[], "in": "l", "loc":{"line":0, "col":0, "path":""}},
        {"id": -13, "env": [7], "sig":[], "in": "doTwoPlusTwo", "loc":{"line":0, "col":0, "path":"script.py"}},
        {"id": -14, "env": [8], "sig":[], "in": "CDLL", "loc":{"line":0, "col":0, "path":""}},
        {"id": -15, "env": [9], "sig":[], "in": "lib.o", "loc":{"line":0, "col":0, "path":""}},
    ],
    "Scopes": [
        {"term": -6, "scope": 0, "level": 0},
        {"term": -7, "scope": 0, "level": 0},
        {"term": -8, "scope": 0, "level": 0},
        {"term": -9, "scope": 0, "level": 0},
        {"term": -10, "scope": 0, "level": 0},
        {"term": -11, "scope": 0, "level": 0},
        {"term": -1, "scope": 1, "level": 1},
        {"term": -12, "scope": 1, "level": 1},
        {"term": -13, "scope": 1, "level": 1},
        {"term": -2, "scope": 1, "level": 1},
        {"term": -14, "scope": 1, "level": 1},
        {"term": -3, "scope": 2, "level": 1},
        {"term": -4, "scope": 2, "level": 1},
        {"term": -5, "scope": 2, "level": 1},
        {"term": -15, "scope": 2, "level": 1},
    ],
    "Monolingual links": [],
    "Crosslingual links": [],
}
```
```yaml
Note: If we somehow do know how ctypes work and have intermodule analysis,
than this analysis results are obtained

Fragments:
    [build.sh]: File
    [lib.c]: File
    [script.py]: File
    [ctypes]: Opaque | { CDLL: String -> Any }
    [doTwoPlusTwo:lib.c]: Unit -> Int
    [./liblib.so]: String
    [l]: { doTwoPlusTwo: { argtypes: List Any, restype: Any } } | String -> Unit -> Any
    [doTwoPlusTwo:script.py]: String
    [cc -c lib.c]: File -> File
    [cc -shared -o liblib.so lib.o]: File -> File
    [python3 script.py]: File -> Any
    [liblib.so]: File

Scope:
    [build.sh]: File
    [lib.c]: File
    [script.py]: File
    [liblib.so]: File
    [ctypes]: Opaque | { CDLL: String -> Any }
    [doTwoPlusTwo:lib.c]: Unit -> Int
        [./liblib.so]: String
        [l]: { doTwoPlusTwo: { argtypes: List Any, restype: Any } } | String -> Unit -> Any
        [doTwoPlusTwo:script.py]: String
        --
        [cc -c lib.c]: File -> File
        [cc -shared -o liblib.so lib.o]: File -> File
        [python3 script.py]: File -> Any

Links:
    [python3 script.py]: File -> Any :-
        [script.py]: File :-
            [ctypes]: File | { CDLL: String -> Any } :-
                [./liblib.so]: String
    [doTwoPlusTwo]: Unit -> Int :- 
        [lib.c]: File
    [l]: { doTwoPlusTwo: { argtypes: List Any, restype: Any } } | String -> Unit -> Any :-
        [ctypes]: { CDLL: String -> Any }
        [doTwoPlusTwo]: String
    [cc -c lib.c]: File -> File :- [lib.c]: File
    [cc -shared -o liblib.so lib.o]: File -> File :- [lib.o]: File
    [liblib.so]: File :- [build.sh]: File

==>
Heuristic about application:
    [l]: String -> Unit -> Any $ [doTwoPlusTwo]: String = ([l] [doTwoPlusTwo]): Unit -> Any
    ([l] [doTwoPlusTwo]): Unit -> Any :- [l]: { doTwoPlusTwo: { argtypes: List Any, restype: Any } }, [l]: String -> Unit -> Any, [doTwoPlusTwo]: String
    
([l] [doTwoPlusTwo:script.py]): Unit -> Any :- [doTwoPlusTwo:lib.c]: Unit -> Int
[script.py:Shell]: File :- [script.py:Python]: File

With probability:
[ctypes]: { CDLL: String -> Any } :- [liblib.so]: File under assumption
    [liblib.so]: File = [./liblib.so]: String 
```
