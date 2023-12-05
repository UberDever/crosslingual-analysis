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
```
[g]: Unit -> Int, [f]: Unit -> Int |- [main]: Unit -> Int with
    [f]: Unit -> Int |- [g]: Unit -> Int with
        [lib.c]: File |- [f]: Unit -> Int with 
            []: Any |- [lib.c]: File
    [main.c]: File |- [f]: Unit -> Int with
        []: Any |- [main.c]: File

[]: Any |- [lib.c]: File
[]: Any |- [main.c]: File
```
```prolog
export [g:lib.c]: Unit -> Int :-
    [f:lib.c]: Unit -> Int :-
        export [lib.c]: File
export [main]: Unit -> Int :-
    [f:main.c]: Unit -> Int :-
        export [main.c]: File :-
            [g:1:main.c]: Unit -> Int
    [g:2:main.c]: Unit -> Int
```
[lib.c]: File
[main.c]: File :-
    [g:1:main.c]: Unit -> Int