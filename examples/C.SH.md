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
cc -DVAR lib.c main.c -o app.exe
```
```
[C:0] VAR: Unit |-
    [C:0] lib.c: File |-
        [C:0] f: Unit -> Int, _ g: Unit -> Int

[C:0] VAR: Bot |-
    [C:0] lib.c: File |-
        _ f: Unit -> Int, _ g: Unit -> Int

[C:0] f: Unit -> Int |-
    _ main.c: File |-
        _ f: Unit -> Int |-
            _ main: Int -> Any -> Int 

_ build.sh: File |-
    _ lib.c: File, _ main.c: File |-
        _ 'cc -DVAR lib.c main.c -o app.exe': File -> File -> File |-
            [Sh:0] VAR: Unit, _ app.exe: File

[Sh:0] VAR: Unit |- [C:0] VAR: Unit

 # Даже не нужно: [C:0] f: Unit -> Int |- [C:0] f: Unit -> Int
```