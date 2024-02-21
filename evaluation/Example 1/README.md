
# Example 1

Simple fullstack app with bidirectional client-server communication.
Golang serves the page, JS fetches state and posts update request to server on button press.
Server serves update request, updates state. JS fetches updated state.

## Required environment

TODO

## Extracted data

```dot

# http.server and http.client are just names for scopes, they
# doesn't matter as far as they linked correctly
# http.server is the scope that golang server creates
# http.client is the scope that javascript client operates in
# object/opaque is when the thing doesn't have an identifier (or i haven't found it)
digraph G {
    "server.go" [shape=box]
    "index.html" [shape=box]
    "http://localhost:3333/item [1]" [shape=box]
    "GET http://localhost:3333/item [2]" [shape=box]
    "POST http://localhost:3333/item [3]" [shape=box]
    "http://localhost:3333/item [4]" [shape=box]
    "GET http://localhost:3333/item [5]" [shape=box]
    "POST http://localhost:3333/item [6]" [shape=box]
    OS -> Filesystem
    Filesystem -> "server.go" [arrowhead=onormal]
    "http.server" -> OS [label=P]
    "server.go" -> "http.server" [arrowhead=onormal]
    "http.server" -> "http://localhost:3333/item [1]"
    "http.server" -> "GET http://localhost:3333/item [2]"
    "http.server" -> "POST http://localhost:3333/item [3]"
    Filesystem -> "index.html" [arrowhead=onormal]
    "index.html" -> "...html stuff... (could be module aswell)" [arrowhead=onormal]
    "index.html" -> "http.client"
    "http.client" -> OS [label=P]
    "http://localhost:3333/item [4]" -> "http.client"
    "GET http://localhost:3333/item [5]" -> "http.client"
    "POST http://localhost:3333/item [6]" -> "http.client"
    "http.client" -> "http://localhost:3333/item [4]" [label=I, arrowhead=onormal]
    "http.client" -> "GET http://localhost:3333/item [5]" [label=I, arrowhead=onormal]
    "http.client" -> "POST http://localhost:3333/item [6]" [label=I, arrowhead=onormal]

    "GET http://localhost:3333/item [2]" -> 1 [arrowhead=onormal]
    1 -> "count [8]"
}

```

### Rename


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