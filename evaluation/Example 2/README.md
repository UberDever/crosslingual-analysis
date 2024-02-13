
# Example 3

Python script passes data to heavy computation function, implemented in fortran.

Makefile script is used to facilitate build

## Required environment

This can be used to extract extension for generated 'python module'
```python
>>> import _imp
>>> _imp.extension_suffixes()
['.cpython-38-x86_64-linux-gnu.so', '.abi3.so', '.so']
```

## Extracted data

TODO

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
