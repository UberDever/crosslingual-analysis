# Example 3

Application written in C# (this one is CLI, but it can also be desktop/server) uses
classes written in Visual Basic to extend it's functionality. All project is managed by solution (although it is not
relevant to the analysis).

## Required environment

Analyze `Example 3/CSharpAndVB.sln` first to get list of all project files, then
extract relevant (for the sake of Example) files: `CSharp.csproj`, `Program.cs`, `Class1.vb`.

Implicit:
- All `using` declarations can be omitted if fully qualified names used instead (we model this like that aswell)

## Extracted constraints

```lisp
(AND
    (Association (Decl OS 1) -1)
    (Declare -1 (Decl Filesystem 2))
    (Association (Decl Filesystem 2) -2)
    (Edge -2 -1 Parent)

    (Declare -2 (Decl CSharp.csproj 3))
    (Association (Decl server.go 3) -15)
    (Edge -15 -3 csharp-project-module) ; This models cross-language barriers
    (Edge -3 -2 Parent)
    (Reference (Ref "~/dev/mag/crosslingual-analysis/evaluation/Example 3/VB/VB.vbproj" 5) -3)

    (Declare -2 (Decl Class1.vb 6))
    (Association (Decl Class1.vb 6) -16)
    (Edge -16 -4 visual-basic-module) ; This models cross-language barriers
    (Edge -4 -2 Parent)
    (Declare -4 (Decl VB 13))
    (Association (Decl VB 13) -9)
    (Edge -5 -4 Parent)
    (Declare -9 (Decl BaseVB 7))
    (Association (Decl BaseVB 7) -5)
    (Edge -9 -5 Parent)
    (Declare -5 (Decl field 7))
    (Typeof (Decl field 7) (Tau 8))
    (Equals (Tau 8) Int)

    (Declare -2 (Decl Program.cs 9))
    (Association (Decl Program.cs 9) -17)
    (Edge -17 -6 csharp-module) ; This models cross-language barriers
    (Edge -6 -2 Parent)
    (Declare -6 (Decl CSharp 10))
    (Association (Decl CSharp 10) -7)
    (Declare -7 (Decl A 11))
    (Association (Decl A 11) -8)
    (Reference (Ref VB 12) -8) 
    (Resolves (Ref VB 12) (Delta 14))
    (Associated (Delta 14) (Sigma 15))
    (Edge -10 (Sigma 15) Import) ; -10 imports (Sigma 15)
    (Reference (Ref BaseVB 16) -10)
    (Resolves (Ref BaseVB 16) (Delta 17))
    (Associated (Delta 17) (Sigma 18))
    (Edge -8 (Sigma 18) Import) ; inheritance 'extends', we import VB.BaseVB
)
```

### Scenarios

- Goto Type Definition: Use `Program.cs:5:15` to navigate to `Class1.vb:2:18`
- Prepare Type Hierarchy: Nothing to show to the user, just change state of an analyzer
- Type Hierarchy Supertypes: Use `Program.cs:5:11` to show `Class1.vb:2:18`
- Type Hierarchy Subtypes: Use `Class1.vb:2:18` to show `Program.cs:5:11`

## Languages

- C#
- Visual Basic
- XML (csproj, vbproj)

## Covered paradigms

- OOP
- Structured (XML)

## Represented domains

- Enterprise
- Fintech

## Expected scenarios

- [Goto Type Definition](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_typeDefinition)
- [Prepare Type Hierarchy](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_prepareTypeHierarchy)
- [Type Hierarchy Supertypes](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#typeHierarchy_supertypes)
- [Type Hierarchy Subtypes](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#typeHierarchy_subtypes)


[Maybe not needed [- [Document symbols](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_documentSymbol)]]