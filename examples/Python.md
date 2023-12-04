```py
import os
import sys
import re

def main():
    if re.match('pattern', sys.argv[0]):
        os.exit(-1)
def foo(): pass
```
```
Note: Behold! Refinement types!

_ (os): File, _ (sys): File, _ (re): File, |-
    _ main.py: File |-
        _ (re): { match: String -> Any -> Any }, _ (sys): { argv: Int -> Any }, _ (os): { exit: Int -> Any }, |-
            _ (main): List String -> Int,
        _ (foo): Unit -> Unit,
```
Linear form:
```
_ (os): File, _ (sys): File, _ (re): File, |- _ main.py: File
_ main.py: File |- _ (re): { match: String -> Any -> Any }, _ (sys): { argv: Int -> Any }, _ (os): { exit: Int -> Any }, |-
_ (re): { match: String -> Any -> Any }, _ (sys): { argv: Int -> Any }, _ (os): { exit: Int -> Any }, |- _ (main): List String -> Int, _ (foo): Unit -> Unit,
```
Logic form:
```
(): Any |- (os): File, (sys): File, (re): File
(main.py): File |- (re): { match: String -> Any -> Any }, (sys): { argv: Int -> Any }, (os): { exit: Int -> Any }
in
    (os): File, (sys): File, (re): File |- (main.py): File
    (re): { match: String -> Any -> Any }, (sys): { argv: Int -> Any }, (os): { exit: Int -> Any } |- (main): List String -> Int
    (re): { match: String -> Any -> Any }, (sys): { argv: Int -> Any }, (os): { exit: Int -> Any } |- (foo): Unit -> Unit
```