# Example 4

Library in C++ is compiled and deployed somewhere. There are also local testing framework consisting of
python and shell files, that runs different scenarios.

Although library is written in C++, it is linked as a C one with wrappers so
Python can use it freely.

TODO: Perhaps this example has more to it in case of usecases, some of them could go here instead

## Required environment

Get all files by `cd "Example 4" && find -name *.cpp || find -name *.py || find -name *.sh`

Explicit (this can be just hardcoded, it is not used in analysis anyway):
- Find `c++` and make sure that it has `-shared` and `-o` flags
- Need to make sure that system has `python3`
Implicit:
- Shell analyzer need to:
    * Do abstract interpretation on paths (like `mkdir`, `touch`...)
    * Know what `mkdir`, `touch`, `rm` etc are

## Extracted constraints

```lisp
(AND
    (Association (Decl OS 1) -1)
    (Declare -1 (Decl Filesystem 2))
    (Association (Decl Filesystem 2) -2)
    (Edge -2 -1 Parent)

    (Declare -2 (Decl build.sh 3))
    (Association (Decl build.sh 3) -15)
    (Edge -15 -3 shell-script-module) ; This models cross-language barriers
    (Edge -3 -2 Parent)
    (Reference (Ref lib.cpp 4) -3)
    (Resolves (Ref lib.cpp 4) (Delta 5))
    (Typeof (Delta 5) (Tau 6))
    (Equals (Tau 6) URI)
    (Declare -3 (Decl lib_dir 7))
    (Association (Decl lib_dir 7) -4)
    (Declare -4 (Decl lib.so 8))
    (Typeof (Decl lib.so 8) (Tau 9))
    (Equals (Tau 9) URI)
    (Associated (Decl lib.so 8) -40)
    ; This thing ties together lib.cpp and lib.so
    ; Basically we say that lib.cpp and lib.so have the same associated scope
    (NominalEdge -40 (Ref lib.cpp 4) Compile)

    (Declare -2 (Decl lib.cpp 10))
    (Association (Decl lib.cpp 10) -16)
    (Edge -16 -4 cpp-module) ; This models cross-language barriers
    (Edge -5 -2 Parent)
    (Declare -5 (Decl counter_new 11))
    (Typeof (Decl counter_new 11) (Tau 12))
    ; Here is interesting thing, we use structural typing here, since names
    ; of types in different languages doesn't need to be the same, only the shapes
    (Equals (Tau 12) (Constructor Function Unit (Constructor Pointer (Constructor Record Int))))
    (Declare -5 (Decl counter_free 13))
    (Typeof (Decl counter_new 13) (Tau 14))
    (Equals (Tau 14) (Constructor Function (Constructor Pointer (Constructor Record Int) Unit)))
    (Declare -5 (Decl counter_get 15))
    (Typeof (Decl counter_get 15) (Tau 16))
    (Equals (Tau 16) (Constructor Function (Constructor Pointer (Constructor Record Int) Int)))
    (Declare -5 (Decl counter_reset 17))
    (Typeof (Decl counter_reset 17) (Tau 18))
    (Equals (Tau 18) (Constructor Function (Constructor Pointer (Constructor Record Int) Unit)))
    (Declare -5 (Decl counter_inc 19))
    (Typeof (Decl counter_inc 19) (Tau 20))
    (Equals (Tau 20) (Constructor Function (Constructor Pointer (Constructor Record Int) Unit)))

    (Declare -2 (Decl test.py 21))
    (Association (Decl test.py 21) -17)
    (Edge -17 -6 python3-module) ; This models cross-language barriers
    (Edge -6 -2 Parent)
    (Reference (Ref "~/dev/mag/crosslingual-analysis/evaluation/Example 4/lib_dir/lib.so" 22) -6)
    (Resolves (Ref "~/dev/mag/crosslingual-analysis/evaluation/Example 4/lib_dir/lib.so" 22) (Delta 23))
    (Associated (Delta 23) (Sigma 24))
    ; Here I use only one scope (-7), but on every `lib.counter_...` there must be new scope
    (Edge -7 (Sigma 24) Import) ; -7 imports (Sigma 24)
    (Reference (Ref counter_new 25) -7)
    (Resolves (Ref counter_new 25) (Delta 26))
    (Typeof (Delta 26) (Tau 27))
    (Equals (Tau 27) (Constructor Function (Constructor Pointer (Constructor Record Int) Unit)))
    (Reference (Ref counter_get 28) -7)
    (Resolves (Ref counter_get 28) (Delta 29))
    (Typeof (Delta 29) (Tau 30))
    (Equals (Tau 30) (Constructor Function (Constructor Pointer (Constructor Record Int) Int)))
    (Reference (Ref counter_reset 31) -7)
    (Resolves (Ref counter_reset 31) (Delta 32))
    (Typeof (Delta 32) (Tau 33))
    (Equals (Tau 33) (Constructor Function (Constructor Pointer (Constructor Record Int) Unit)))
    (Reference (Ref counter_inc 34) -7)
    (Resolves (Ref counter_inc 34) (Delta 35))
    (Typeof (Delta 35) (Tau 36))
    (Equals (Tau 36) (Constructor Function (Constructor Pointer (Constructor Record Int) Unit)))
    (Reference (Ref counter_free 37) -7)
    (Resolves (Ref counter_free 37) (Delta 38))
    (Typeof (Delta 38) (Tau 39))
    (Equals (Tau 39) (Constructor Function (Constructor Pointer (Constructor Record Int) Unit)))
)
```

### Scenarios

- Completion: Autocomplete on `lib.counter_` in `test.py` somewhere
- Completion resolve: same as the above
- Signature help: Show signature for `test.py:45:9`

## Languages

- C++
- Python 3
- Shell

## Covered paradigms

- Procedural
- Shell

## Represented domains

- Embedded
- Testing
- DevOps

## Expected scenarios

- [Completion](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_completion)
- [Completion resolve](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionItem_resolve)
- [Signature help](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_signatureHelp)
