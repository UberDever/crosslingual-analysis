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
```
Note: If we somehow do know how ctypes work and have intermodule analysis,
than this analysis results are obtained

_ (ctypes): File,
_ (./liblib.so): File |-
    _ (ctypes): { |CDLL|: String -> Any } |-
        _ (l): { |doTwoPlusTwo|: { |argtypes|: List Any, |restype|: Any } ∪ 
        String -> Unit -> Any } |-
            _ (l['doTwoPlusTwo']): Unit -> Any
    [Python:0] script.py: File

_ (lib.c): File |-
    _ (doTwoPlusTwo): Unit -> Int

_ (lib.c): File |- _ (cc -c lib.c): File -> File,
_ (lib.o): File |- _ (cc -shared -o liblib.so lib.o): File -> File,
[Shell:0] script.py: File |- _ (python3 script.py): File -> Any, |-
    _ (build.sh): File |-
        _ (liblib.so): File

[Python:0] 'script.py': File |- [Shell:0] 'script.py': File
```
```
Note: This example uses just implicit knowledge that is baked in the parser-translator 
(e.g we don't know what ctypes.CDD is and how brackets of l work)

_ (ctypes): File, |-
    [Python:0] script.py: File

_ (lib.c): File |-
    _ (doTwoPlusTwo): Unit -> Int

_ (lib.c): File |- _ (cc -c lib.c): File -> File,
_ (lib.o): File |- _ (cc -shared -o liblib.so lib.o): File -> File,
[Shell:0] script.py: File |- _ (python3 script.py): File -> Any, |-
    _ (build.sh): File |-
        _ (liblib.so): File

[Python:0] (script.py): File |- [Shell:0] (script.py): File

```
