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