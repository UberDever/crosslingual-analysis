
# Example 1

Simple fullstack app with bidirectional client-server communication.
Golang serves the page, JS fetches state and posts update request to server on button press.
Server serves update request, updates state. JS fetches updated state.

## Required environment

Implicit (encoded in the translators):
- http.HandleFunc modifies creates new URIs on the basis of <host>
- http.ListenAndServe with the nil parameter sets <host> to localhost

## Extracted constraints

[Notation is defined in top-level readme](../README.md)

```lisp
(AND
    (Associate (Decl OS 1) -1)
    (Declare -1 (Decl Filesystem 2))
    (Associate (Decl Filesystem 2) -2)
    (Edge -2 -1 Parent)

    (Declare -2 (Decl server.go 3))
    (Associate (Decl server.go 3) -3) ; -3 is "http.server"
    (Edge -3 -2 Parent)
    (Declare -3 (Decl http://localhost:3333/item 4))
    (Associate (Decl http://localhost:3333/item 4) -4)
    (Declare -4 (Decl GET 5))
    (Associate (Decl GET 5) -8)
    (Declare -8 (Decl Response 18))
    (Associate (Decl Response 18) -9)
    (Declare -10 (Decl Json 19))
    (Associate (Decl Json 19) -11)
    (Declare -11 (Decl count 30))
    (Typeof (Decl count 30) (Tau 31))
    (Equals (Tau 31) Int)

    (Declare -4 (Decl POST 6))
    (Typeof (Decl POST 6) (Tau 32))
    (Equals (Tau 32) Top)

    (Declare -2 (Decl index.html 7))
    (Associate (Decl index.html) -5) ; -5 is "http.client"
    (Edge -5 -2 Parent)
    (Reference (Ref http://localhost:3333/item 8) -5)
    (Resolves (Ref http://localhost:3333/item 8) (Delta 9))
    (Associated (Delta 9) (Sigma 10))
    (Edge -6 (Sigma 10) Import) ; -6 imports (Sigma 10)
    (Reference (Ref POST 11) -6)
    (Resolves (Ref POST 11) (Delta 12))
    (Typeof (Ref POST 11) (Tau 33))
    (Equals (Tau 33) Top)

    (Reference (Ref http://localhost:3333/item 13) -5)
    (Resolves (Ref http://localhost:3333/item 13) (Delta 14))
    (Associated (Delta 14) (Sigma 15))
    (Edge -7 (Sigma 15) Import) ; -7 imports (Sigma 15)
    (Reference (Ref GET 16) -7)
    (Resolves (Ref GET 16) (Delta 17))
    (Associated (Delta 17) (Sigma 20))
    (Edge -12 (Sigma 20) Import)
    (Reference (Ref Response 21) -12)
    (Resolves (Ref Response 21) (Delta 22))
    (Associated (Delta 22) (Sigma 23))
    (Edge -13 (Sigma 23) Import)
    (Reference (Ref Json 24) -13)
    (Resolves (Ref Json 24) (Delta 25))
    (Associated (Delta 25) (Sigma 26))
    (Edge -14 (Sigma 26) Import)
    (Reference (Ref count 27) -14)
    (Resolves (Ref count 27) (Delta 28))
    (Typeof (Ref count 27) (Tau 29))
    (Equals (Tau 29) Top)
)
```

### Scenarios

- Rename: index.html:26:108 and server.go:10:22 rename in both `count` to `counter`
- Find references: find places where http://localhost:3333/item is referenced (index.html:15:20, index.html:20:20, server.go:20:10, server.go:37:10)
- Goto Declaration: find declaraton of http://localhost:3333/item (server.go:37:10)
- Goto Definition: find def of http://localhost:3333/item (server.go:20:10)

## Languages

- HTML
- JS (in HTML)
- Golang https://github.com/robertkrimen/otto

## Covered paradigms

- Multi-paradigm (JS)
- Markup

## Represented domains

- Web

## Expected scenarios

- [Rename](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_rename)
- [Find References](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_references)
- [Goto Declaration](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_declaration)
- [Goto Definition](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_definition)