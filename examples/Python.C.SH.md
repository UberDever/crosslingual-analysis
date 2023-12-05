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
Logic form:
```
[]: Any |- [build.sh]: File
[]: Any |- [lib.c]: File
[ctypes]: File |- [script.py]: File
[lib.c]: File |- [doTwoPlusTwo]: Unit -> Int
in
    [./liblib.so]: String |- [ctypes]: { CDLL: String -> Any }
    [ctypes]: { CDLL: String -> Any } |- [l]: { doTwoPlusTwo: { argtypes: List Any, restype: Any } }
    [l]: { doTwoPlusTwo: { argtypes: List Any, restype: Any } }, [doTwoPlusTwo]: String |- [l]: String -> Unit -> Any
    [lib.c]: File |- [cc -c lib.c]: File -> File
    [lib.o]: File |- [cc -shared -o liblib.so lib.o]: File -> File
    [script.py]: File |- [python3 script.py]: File -> Any
    [build.sh]: File |- [liblib.so]: File
==>
Heuristic about application:
    [l]: String -> Unit -> Any $ [doTwoPlusTwo]: String === ([l] [doTwoPlusTwo]): Unit -> Any
    [l]: { doTwoPlusTwo: { argtypes: List Any, restype: Any } }, [l]: String -> Unit -> Any, [doTwoPlusTwo]: String |- ([l] [doTwoPlusTwo]): Unit -> Any
    
[doTwoPlusTwo]: Unit -> Int |- ([l] [doTwoPlusTwo]): Unit -> Any
[script.py:Python]: File |- [script.py:Shell]: File
With probability:
[liblib.so]: File |- [ctypes]: { CDLL: String -> Any } with
    [liblib.so]: File = [./liblib.so]: String 
```
```
Note: If we somehow do know how ctypes work and have intermodule analysis,
than this analysis results are obtained

_ [ctypes]: File,
_ (./liblib.so): File |-
    _ [ctypes]: { |CDLL|: String -> Any } |-
        _ [l]: { |doTwoPlusTwo|: { |argtypes|: List Any, |restype|: Any } ∪ 
        String -> Unit -> Any } |-
            _ (l['doTwoPlusTwo']): Unit -> Any
    [Python:0] script.py: File

_ [lib.c]: File |-
    _ [doTwoPlusTwo]: Unit -> Int

_ [lib.c]: File |- _ [cc -c lib.c]: File -> File,
_ [lib.o]: File |- _ [cc -shared -o liblib.so lib.o]: File -> File,
[Shell:0] script.py: File |- _ [python3 script.py]: File -> Any, |-
    _ [build.sh]: File |-
        _ [liblib.so]: File

[Python:0] 'script.py': File |- [Shell:0] 'script.py': File
```
```
Note: This example uses just implicit knowledge that is baked in the parser-translator 
(e.g we don't know what ctypes.DLL is and how brackets of l work)

_ [ctypes]: File, |-
    [Python:0] script.py: File

_ [lib.c]: File |-
    _ [doTwoPlusTwo]: Unit -> Int

_ [lib.c]: File |- _ [cc -c lib.c]: File -> File,
_ [lib.o]: File |- _ [cc -shared -o liblib.so lib.o]: File -> File,
[Shell:0] script.py: File |- _ [python3 script.py]: File -> Any, |-
    _ [build.sh]: File |-
        _ [liblib.so]: File

[Python:0] [script.py]: File |- [Shell:0] [script.py]: File

```
Linear form:
```
_ [ctypes]: File, |- [Python:0] script.py: File

_ [lib.c]: File |- _ [doTwoPlusTwo]: Unit -> Int

_ [lib.c]: File |- _ [cc -c lib.c]: File -> File,
_ [lib.o]: File |- _ [cc -shared -o liblib.so lib.o]: File -> File,
[Shell:0] script.py: File |- _ [python3 script.py]: File -> Any,
_ [python3 script.py]: File -> Any, |- _ [build.sh]: File 
_ [build.sh]: File |- _ [liblib.so]: File

[Python:0] [script.py]: File |- [Shell:0] [script.py]: File
```
