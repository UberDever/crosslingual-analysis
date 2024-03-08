`lib.c`
```c
static int f() {return 42;}
int g() {return f() + 5}
```
`main.c`
```c
static int f() {return 84;}
int g();
int main() {return f() + g(); }
```
Results:
```yaml
Fragments:
    [lib.c]: File
    [main.c]: File
    [g:lib.c]: Unit -> Int
    [f:lib.c]: Unit -> Int
    [main]: Unit -> Int
    [f:main.c]: Unit -> Int
    [g:1:main.c]: Unit -> Int
    [g:2:main.c]: Unit -> Int

Scope:
    # filesystem
    [g:lib.c]: Unit -> Int
    [main]: Unit -> Int
    [main.c]: File
    [lib.c]: File
        -- # lib.c
        [f:lib.c]: Unit -> Int
        -- # main.c
        [f:main.c]: Unit -> Int

Links:
    [g:lib.c]: Unit -> Int :-
        [f:lib.c]: Unit -> Int :-
            [lib.c]: File
    [main]: Unit -> Int :-
        [f:main.c]: Unit -> Int :-
            [main.c]: File :-
                [g:1:main.c]: Unit -> Int
        [g:2:main.c]: Unit -> Int
```