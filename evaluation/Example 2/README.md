
# Example 2

Python script passes data to heavy computation function, implemented in Fortran.

Makefile script is used to facilitate build

## Required environment

Also note that this example shows peculiar problem - some identifiers in one language can be case-insensitive, while in the others this is not the case

Explicit:
- Extract extensions for FFI modules for python and use them during analysis of `f2py3` command in shell
```python
>>> import _imp
>>> _imp.extension_suffixes()
['.cpython-38-x86_64-linux-gnu.so', '.abi3.so', '.so']
```
- Use `make --dry-run` to list all shell commands and extract all commands that build the project
- Use that commands to get list of all files, also they would be used as source of information from Makefile (without traceback to makefile itself tho :( )

## Extracted constraints

```lisp

(AND
    (Association (Decl OS 1) -1)
    (Declare -1 (Decl Filesystem 2))
    (Association (Decl Filesystem 2) -2)
    (Edge -2 -1 Parent)

    (Declare -2 (Decl Makefile 3))
    (Association (Decl build.sh 3) -15)
    (Edge -15 -3 makefile-module) ; This models cross-language barriers
    (Edge -3 -2 Parent)
    (Reference (Ref lib.f90 4) -3)
    (Resolves (Ref lib.f90 4) (Delta 5))
    (Typeof (Delta 5) (Tau 6))
    (Equals (Tau 6) URI)
    (Declare -3 (Decl compute_lib 7))
    (Typeof (Decl compute_lib 7) (Tau 16))
    (Equals (Tau 16) File)
    (Associated (Decl compute_lib 7) -50)
    ; This thing ties together lib.f90 and compute_lib.*
    ; Basically we say that they have the same associated scope
    (NominalEdge -50 (Ref lib.f90 4) Compile)

    (Declare -2 (Decl lib.f90 10))
    (Association (Decl lib.f90 10) -16)
    (Edge -16 -4 fortran90-module) ; This models cross-language barriers
    (Edge -5 -2 Parent)
    (Declare -5 (Decl ExpensiveComputation 11))
    (Typeof (Decl ExpensiveComputation 11) (Tau 12))
    (Equals (Tau 12) (Constructor Function Float Float))

    (Declare -2 (Decl compute.py 13))
    (Association (Decl compute.py 13) -17)
    (Edge -17 -6 python3-module) ; This models cross-language barriers
    (Edge -6 -2 Parent)
    (Reference (Ref compute_lib 14) -6)
    (Resolves (Ref compute_lib 14) (Delta 15))
    (Associated (Delta 15) (Sigma 17))
    (Edge -6 (Sigma 17) Import)
    (Typeof (Delta 15) (Tau 18))
    (Equals (Tau 18) File)
    (Reference (Ref expensivecomputation 19) -6)
    (Resolves (Ref expensivecomputation 19) (Delta 20))
    (Typeof (Delta 20) (Tau 21))
    (Equals (Tau 21) (Constructor Function Float Top))
)

```

### Scenarios

- Call Hierarchy Incoming Calls: Find the callers of `lib.f90:1:10` - it should be `compute.py:3:19`
- Call Hierarchy Outgoing Calls: Find the callee of `lib.f90:1:10` - there should be no one (afaik this could be python module itself, since it is explicitly a function)
- Prepare Call Hierarchy Request: Just the change in the state of the analyzer

## Languages

- Python
- Makefile
- Fortran

## Covered paradigms

- Declarative
- Functional (pretend that python code is functional one)

## Represented domains

- Scientific computing
- Data science
- Build systems

## Expected scenarios

- [Call Hierarchy Incoming Calls](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#callHierarchy_incomingCalls)
- [Call Hierarchy Outgoing Calls](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#callHierarchy_outgoingCalls)
- [Prepare Call Hierarchy Request](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_prepareCallHierarchy)
