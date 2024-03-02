# Example 5

Golang service connects to a local database and processes data to serve clients with information about the weather.

For a sake of simplicity database is simple file, but can also be a mongodb or some database with a connection string.

## Required environment

Get all files by `cd "Example 5" && find -name *.go || find -name *.json`

Implicit (encoded in the translators):
- http.HandleFunc modifies creates new URIs on the basis of <host> (declaration)
- http.ListenAndServe with the nil parameter sets <host> to localhost
- os.ReadFile reads file that is specified by filepath (reference)

## Extracted constraints

[Notation is defined in top-level readme](../README.md)

```lisp
(AND
    (Associated (Decl OS 1) -1)
    (Declare -1 (Decl Filesystem 2))
    (Associated (Decl Filesystem 2) -2)
    (Edge -2 -1 Parent)

    (Declare -2 (Decl server.go 3))
    (Associated (Decl server.go 3) -15) ; -3 is "http.server"
    (Edge -15 -3 golang-module) ; This models cross-language barriers
    (Edge -3 -2 Parent)
    (Declare -3 (Decl http://localhost:8080/ 4))

    (Reference (Ref "~/dev/mag/crosslingual-analysis/evaluation/Example 5/weather.json" 5) -3)
    (Resolves (Ref "~/dev/mag/crosslingual-analysis/evaluation/Example 5/weather.json" 5) (Delta 6))

    (Declare -2 (Decl weather.json 6))
    (Associated (Decl weather.json 6) -16)
    (Edge -16 -4 json-module) ; This models cross-language barriers
    (Edge -4 -2 Parent)
    ; Here must go all json structure in form of the tree with only lists being typed
    ; But we have only necessary stuff here
    (Edge -5 -4 Parent) ; Root object (this is a questionable way to encode this)
    (Declare -5 (Decl main 7))
    (Associated (Decl main 7) -6)
    (Declare -6 (Decl temp 8))
    (Typeof (Decl temp 8) (Tau 9))
    (Equals (Tau 9) Float)
)
```

### Scenarios

- Pull diagnostics: First run pull diagnostics as is, then rename `temp` in `weather.json:16:9`
to `temperature` and run diagnostics again. Should expect squigglies or some stuff in
 `server.go:32:8` on json access
- Document Link Request/Resolve: Resolve `weather.json` line in `server.go:25:29` to the actual file
(yeah vscode already can do this, but who cares) 

## Languages

- Golang
- JSON

## Covered paradigms

- Datum (structured data)
- Procedural

## Represented domains

- Client-server
- Database management

## Expected scenarios

- [Pull diagnostics](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_pullDiagnostics)
- [Document Link Request](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_documentLink)
- [Document Link Resolve Request](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#documentLink_resolve)