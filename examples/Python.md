```py
import os
import sys
import re

def main():
    if re.match('pattern', sys.argv[0]):
        os.exit[-1]
def foo(): pass
```
Results:
```yaml
Fragments:
    [os]: File | { exit: Int -> Any }
    [sys]: File | { argv: Int -> Any }
    [re]: File | { match: String -> Any -> Any }
    [main.py]: File
    [main]: List String -> Int
    [foo]: Unit -> Unit

Scope:
    [os]: File | { exit: Int -> Any }
    [sys]: File | { argv: Int -> Any }
    [re]: File | { match: String -> Any -> Any }
    [main.py]: File
    [main]: List String -> Int
    [foo]: Unit -> Unit

Links:
    [main]: List String -> Int, [foo]: Unit -> Unit :-
        [os]: File | { exit: Int -> Any }
        [sys]: File | { argv: Int -> Any }
        [re]: File | { match: String -> Any -> Any }
        [main.py]: File :-
            [os]: File | { exit: Int -> Any }
            [sys]: File | { argv: Int -> Any }
            [re]: File | { match: String -> Any -> Any }
```