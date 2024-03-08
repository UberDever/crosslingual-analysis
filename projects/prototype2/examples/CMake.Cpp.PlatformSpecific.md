
`CMakeLists.txt`
```cmake
cmake_minimum_required(VERSION 3.16)

set(PROJECT_NAME platform_specific)
project(${PROJECT_NAME} CXX)

string(TOLOWER ${CMAKE_SYSTEM_NAME} SYSTEM)

set(COMMON_SRC
    ${CMAKE_CURRENT_SOURCE_DIR}/common/main.cpp)

set(PLATFORM_SRC
    ${CMAKE_CURRENT_SOURCE_DIR}/${SYSTEM}/add.cpp)

set(SRC_FILES ${COMMON_SRC} ${PLATFORM_SRC})

add_executable(${PROJECT_NAME} ${SRC_FILES})
```
`linux/add.cpp`
```cpp
int add(int a, int b) { return a + b; }
```
`windows/add.cpp`
```cpp
int add(int a, int b) { return a - b; }
```
`common/main.cpp`
```cpp
int add(int a, int b);
int main() { return add(1, 2); }
```
`build.sh`
```sh
build_windows() {
    mkdir -p build
    cmake -b build -DCMAKE_SYSTEM_NAME=Windows ./CMakeLists.txt
}

build_linux() {
    mkdir -p build
    cmake -b build -DCMAKE_SYSTEM_NAME=Linux ./CMakeLists.txt
}

```
Results:
```yaml
Fragments:
    - { term: "PROJECT_NAME", type: "String" }
    - { term: "platform_specific", type: "String" }
    - { term: "${PROJECT_NAME}", type: "String" }
    - { term: "${CMAKE_SYSTEM_NAME}", type: "String" }
    - { term: "SYSTEM", type: "String" }


export [add]: (Int, Int) -> Int :- 
    export [linux/add.cpp]: File
export [add]: (Int, Int) -> Int :- 
    export [windows/add.cpp]: File
export [main]: Unit -> Int :- [add]: (Int, Int) -> Int :-
    export [main.cpp]: File

[mkdir -p build]: Unit -> Directory & {name: 'build'} | Unit :-
    export [build_windows]: Unit -> Int :-
        export [build.sh]: File
export [CMAKE_SYSTEM_NAME]: String & 'Windows' :-
    [cmake -p build -DCMAKE_SYSTEM_NAME=Windows ./CMakeLists.txt]: Any -> Any :-
        export [build_windows]: Unit -> Int :-
            export [build.sh]: File
        [CMakeLists.txt]: File
[mkdir -p build]: Unit -> Directory & {name: 'build'} | Unit :-
    export [build_linux]: Unit -> Int :-
        export [build.sh]: File
export [CMAKE_SYSTEM_NAME]: String & 'Linux' :-
    [cmake -p build -DCMAKE_SYSTEM_NAME=Linux ./CMakeLists.txt]: Any -> Any :-
        export [build_linux]: Unit -> Int :-
            export [build.sh]: File
        [CMakeLists.txt]: File

[CMAKE_CURRENT_SOURCE_DIR]: String :-
    [set(COMMON_SRC ${CMAKE_CURRENT_SOURCE_DIR}/common/main.cpp)]: Unit 

[SRC_FILES]: List String :-
    [PLATFORM_SRC]: List String :-
        :-
            [${CMAKE_CURRENT_SOURCE_DIR}/${SYSTEM}/*.cpp]: List File :-
                [SYSTEM]: String :- 
                    [string(TOLOWER ${CMAKE_SYSTEM_NAME} SYSTEM)]: Unit :-
                        export [CMakeLists.txt]: File
                        [CMAKE_SYSTEM_NAME]: String
    [COMMON_SRC]: List String :-
        [file(GLOB COMMON_SRC ${CMAKE_CURRENT_SOURCE_DIR}/common/*.cpp)]: Unit :-
            export [CMakeLists.txt]: File
            [${CMAKE_CURRENT_SOURCE_DIR}/common/*.cpp]: List File
```

