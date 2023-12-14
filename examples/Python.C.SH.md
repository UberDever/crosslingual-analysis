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
    "Expressions": [
        {"id": -1, "used": [3, 5, 8], "in": "ctypes.CDLL('./liblib.so')"},
        {"id": -2, "used": [6, 7], "in": "l['doTwoPlusTwo']"},
        {"id": -3, "used": [1], "in": "cc -c lib.c"},
        {"id": -4, "used": [9], "in": "cc -shared -o liblib.so lib.o"},
        {"id": -5, "used": [2], "in": "python3 script.py"},
    ],
    "Scopes": [
        {"term": 0, "scope": 0, "level": 0},
        {"term": 1, "scope": 0, "level": 0},
        {"term": 2, "scope": 0, "level": 0},
        {"term": 3, "scope": 0, "level": 0},
        {"term": 4, "scope": 0, "level": 0},
        {"term": 5, "scope": 0, "level": 0},
        {"term": -1, "scope": 1, "level": 1},
        {"term": 6, "scope": 1, "level": 1},
        {"term": 7, "scope": 1, "level": 1},
        {"term": -2, "scope": 1, "level": 1},
        {"term": 8, "scope": 1, "level": 1},
        {"term": -3, "scope": 2, "level": 1},
        {"term": -4, "scope": 2, "level": 1},
        {"term": -5, "scope": 2, "level": 1},
    ],
    "Monolingual links": [

    ],
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
