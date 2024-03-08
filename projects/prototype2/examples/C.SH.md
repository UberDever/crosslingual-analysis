```c
// lib.c
#ifdef VAR
int f() {
    return 1;
}
#else
int g() {
    return 2;
}
#endif

// main.c
int f();

int main(int argc, char** argv) {
    return f();
}
```
```sh
# build.sh
cc -DVAR lib.c main.c -o app.exe
```
Results:
```yaml
Fragments:
    [VAR]: Bot
    [VAR]: Unit
    [f:lib.c]: Unit -> Int
    [g]: Unit -> Int
    [f:main.c]: Unit -> Int
    [main]: Int -> List String -> Int
    [lib.c]: File
    [main.c]: File
    [build.sh]: File
    [cc -DVAR lib.c main.c -o app.exe]: File -> File -> File
    [VAR:build.sh]: Unit

Scope:
    [VAR:build.sh]: Unit
    [f:main.c]: Unit -> Int
    [main]: Int -> List String -> Int
    [lib.c]: File
    [main.c]: File
    [build.sh]: File
        [cc -DVAR lib.c main.c -o app.exe]: File -> File -> File
        --
        [VAR]: Bot
        [g]: Unit -> Int
        --
        [VAR]: Unit
        [f:lib.c]: Unit -> Int

Links:
    [VAR:build.sh]: Unit :-
        [build.sh]: File
    [cc -DVAR lib.c main.c -o app.exe]: File -> File -> File :-
        [build.sh]: File
    [main]: Int -> List String -> Int :-
        [f:main.c]: Unit -> Int :-
            [main.c]: File
    [g]: Unit -> Int :-
        [VAR]: Bot :-
            [lib.c]: File
    [f:lib.c]: Unit -> Int :-
        [VAR]: Unit :-
            [lib.c]: File
    [VAR]: Unit :- 
        [VAR:build.sh]: Unit
    [f:main.c]: Unit -> Int :- 
        [f:lib.c]: Unit -> Int
```