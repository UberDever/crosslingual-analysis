import { Constraint, UsageConstraint, DirectEdgeConstraint, AssociationConstraint, TypeDeclarationConstraint, TypeEqualConstraint, SourceFile, ResolutionConstraint, Substitution, Variable, MustResolveConstraint } from "./protocol"

const constraints: Constraint[] = [
    {
        identifier: { name: "OS", source: undefined },
        scope: { index: 0, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: { name: "OS", source: undefined },
        scope: { index: 1, source: undefined }
    } satisfies AssociationConstraint,
    {
        identifier: { name: "Filesystem", source: undefined },
        scope: { index: 1, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: { name: "Filesystem", source: undefined },
        scope: { index: 2, source: undefined }
    } satisfies AssociationConstraint,
    {
        from: { index: 2, source: undefined },
        to: { index: 1, source: undefined },
        label: "parent"
    } satisfies DirectEdgeConstraint,
    {
        identifier: {
            name: "csharp.csproj",
            source: { uri: "examples/Example 1/CSharp/CSharp.csproj", language: undefined } satisfies SourceFile
        },
        scope: { index: 2, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "csharp.csproj",
            source: { uri: "examples/Example 1/CSharp/CSharp.csproj", language: undefined } satisfies SourceFile
        },
        scope: { index: 3, source: undefined }
    } satisfies AssociationConstraint,
    {
        from: { index: 3, source: undefined },
        to: { index: 2, source: undefined },
        label: "parent"
    } satisfies DirectEdgeConstraint,
    {
        identifier: {
            name: "Example 1/VB/VB.vbproj",
            source: {
                uri: "examples/Example 1/CSharp/CSharp.csproj", range: {
                    start: { line: 4, character: 32 },
                    end: { line: 4, character: 47 }
                }, language: undefined
            } satisfies SourceFile
        },
        scope: { index: 3, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        identifier: {
            name: "Class1.vb",
            source: { uri: "Example 1/VB/Class1.vb", language: undefined } satisfies SourceFile
        },
        scope: { index: 2, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "Class1.vb",
            source: { uri: "Example 1/VB/Class1.vb", language: undefined } satisfies SourceFile
        },
        scope: { index: 4, source: undefined }
    } satisfies AssociationConstraint,
    {
        from: { index: 4, source: undefined },
        to: { index: 2, source: undefined },
        label: "parent"
    } satisfies DirectEdgeConstraint,
    {
        identifier: {
            name: "VB",
            source: {
                uri: "Example 1/VB/Class1.vb", range: {
                    start: { line: 1, character: 1 },
                    end: { line: 1, character: 1 }
                }, language: "Visual Basic"
            } satisfies SourceFile
        },
        scope: { index: 4, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "VB",
            source: {
                uri: "Example 1/VB/Class1.vb", range: {
                    start: { line: 1, character: 1 },
                    end: { line: 1, character: 1 }
                }, language: "Visual Basic"
            } satisfies SourceFile
        },
        scope: { index: 5, source: undefined }
    } satisfies AssociationConstraint,
    {
        from: { index: 5, source: undefined },
        to: { index: 4, source: undefined },
        label: "parent"
    } satisfies DirectEdgeConstraint,
    {
        identifier: {
            name: "BaseVB",
            source: {
                uri: "Example 1/VB/Class1.vb", range: {
                    start: { line: 1, character: 14 },
                    end: { line: 1, character: 20 }
                }, language: "Visual Basic"
            } satisfies SourceFile
        },
        scope: { index: 5, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "BaseVB",
            source: {
                uri: "Example 1/VB/Class1.vb", range: {
                    start: { line: 1, character: 14 },
                    end: { line: 1, character: 20 }
                }, language: "Visual Basic"
            } satisfies SourceFile
        },
        type: { index: 6, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 6, type: "tau" },
        rhs: "BaseVB"
    } satisfies TypeEqualConstraint,
    {
        declaration: {
            name: "BaseVB",
            source: {
                uri: "Example 1/VB/Class1.vb", range: {
                    start: { line: 1, character: 14 },
                    end: { line: 1, character: 20 }
                }, language: "Visual Basic"
            } satisfies SourceFile
        },
        scope: { index: 7, source: undefined }
    } satisfies AssociationConstraint,
    {
        identifier: {
            name: "field",
            source: {
                uri: "Example 1/VB/Class1.vb", range: {
                    start: { line: 2, character: 12 },
                    end: { line: 2, character: 17 }
                }, language: "Visual Basic"
            } satisfies SourceFile
        },
        scope: { index: 7, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "field",
            source: {
                uri: "Example 1/VB/Class1.vb", range: {
                    start: { line: 2, character: 12 },
                    end: { line: 2, character: 17 }
                }, language: "Visual Basic"
            } satisfies SourceFile
        },
        type: { index: 8, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 8, type: "tau" },
        rhs: "Integer"
    } satisfies TypeEqualConstraint,
    {
        identifier: {
            name: "Program.cs",
            source: { uri: "Example 1/CSharp/Program.cs", language: undefined } satisfies SourceFile
        },
        scope: { index: 2, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "Program.cs",
            source: { uri: "Example 1/CSharp/Program.cs", language: undefined } satisfies SourceFile
        },
        scope: { index: 9, source: undefined }
    } satisfies AssociationConstraint,
    {
        from: { index: 9, source: undefined },
        to: { index: 2, source: undefined },
        label: "parent"
    } satisfies DirectEdgeConstraint,
    {
        identifier: {
            name: "CSharp",
            source: {
                uri: "Example 1/CSharp/Program.cs", range: {
                    start: { line: 3, character: 11 },
                    end: { line: 3, character: 17 }
                }, language: "C#"
            } satisfies SourceFile
        },
        scope: { index: 9, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "CSharp",
            source: {
                uri: "Example 1/CSharp/Program.cs", range: {
                    start: { line: 3, character: 11 },
                    end: { line: 3, character: 17 }
                }, language: "C#"
            } satisfies SourceFile
        },
        scope: { index: 10, source: undefined }
    } satisfies AssociationConstraint,
    {
        from: { index: 10, source: undefined },
        to: { index: 9, source: undefined },
        label: "parent"
    } satisfies DirectEdgeConstraint,
    {
        identifier: {
            name: "A",
            source: {
                uri: "Example 1/CSharp/Program.cs", range: {
                    start: { line: 5, character: 11 },
                    end: { line: 5, character: 12 }
                }, language: "C#"
            } satisfies SourceFile
        },
        scope: { index: 10, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "A",
            source: {
                uri: "Example 1/CSharp/Program.cs", range: {
                    start: { line: 5, character: 11 },
                    end: { line: 5, character: 12 }
                }, language: "C#"
            } satisfies SourceFile
        },
        scope: { index: 11, source: undefined }
    } satisfies AssociationConstraint,
    {
        identifier: {
            name: "VB",
            source: {
                uri: "Example 1/CSharp/Program.cs", range: {
                    start: { line: 5, character: 15 },
                    end: { line: 5, character: 15 }
                }, language: "C#"
            } satisfies SourceFile
        },
        scope: { index: 11, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        reference: {
            name: "VB",
            source: {
                uri: "Example 1/CSharp/Program.cs", range: {
                    start: { line: 5, character: 15 },
                    end: { line: 5, character: 15 }
                }, language: "C#"
            } satisfies SourceFile
        },
        declaration: { index: 12, type: "delta" }
    } satisfies ResolutionConstraint,
    {
        declaration: { index: 12, type: "delta" },
        scope: { index: 13, type: "sigma" }
    } satisfies AssociationConstraint,
    {
        from: { index: 13, type: "sigma" },
        to: { index: 14, source: undefined },
        label: "import"
    } satisfies DirectEdgeConstraint,
    {
        identifier: {
            name: "BaseVB",
            source: {
                uri: "Example 1/CSharp/Program.cs", range: {
                    start: { line: 5, character: 15 },
                    end: { line: 5, character: 21 }
                }, language: "C#"
            } satisfies SourceFile
        },
        scope: { index: 14, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        reference: {
            name: "BaseVB",
            source: {
                uri: "Example 1/CSharp/Program.cs", range: {
                    start: { line: 5, character: 15 },
                    end: { line: 5, character: 21 }
                }, language: "C#"
            } satisfies SourceFile
        },
        declaration: { index: 15, type: "delta" }
    } satisfies ResolutionConstraint,
    {
        declaration: { index: 15, type: "delta" },
        scope: { index: 16, type: "sigma" }
    } satisfies AssociationConstraint,
    {
        from: { index: 16, type: "sigma" },
        to: { index: 11, source: undefined },
        label: "import"
    } satisfies DirectEdgeConstraint,
    {
        identifier: {
            name: "field",
            source: {
                uri: "Example 1/CSharp/Program.cs", range: {
                    start: { line: 9, character: 13 },
                    end: { line: 9, character: 18 }
                }, language: "C#"
            } satisfies SourceFile
        },
        scope: { index: 11, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "field",
            source: {
                uri: "Example 1/CSharp/Program.cs", range: {
                    start: { line: 9, character: 13 },
                    end: { line: 9, character: 18 }
                }, language: "C#"
            } satisfies SourceFile
        },
        type: { index: 17, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 17, type: "tau" },
        rhs: "Integer"
    } satisfies TypeEqualConstraint,
    {
        identifier: {
            name: "field",
            source: {
                uri: "Example 1/CSharp/Program.cs", range: {
                    start: { line: 13, character: 43 },
                    end: { line: 13, character: 52 }
                }, language: "C#"
            } satisfies SourceFile
        },
        scope: { index: 11, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "field",
            source: {
                uri: "Example 1/CSharp/Program.cs", range: {
                    start: { line: 13, character: 43 },
                    end: { line: 13, character: 52 }
                }, language: "C#"
            } satisfies SourceFile
        },
        type: { index: 18, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 18, type: "tau" },
        rhs: "Top"
    } satisfies TypeEqualConstraint,
    {
        reference: {
            name: "BaseVB",
            source: {
                uri: "Example 1/CSharp/Program.cs", range: {
                    start: { line: 5, character: 15 },
                    end: { line: 5, character: 21 }
                }, language: "C#"
            } satisfies SourceFile
        },
        scope: { index: 14, source: undefined },
    } satisfies MustResolveConstraint,
]

const substitution: Map<Variable, Substitution> = new Map<Variable, Substitution>([
    [{ index: 6, type: "tau" }, { tag: "tau", data: "BaseVB" }],
    [{ index: 8, type: "tau" }, { tag: "tau", data: "Integer" }],
    [{ index: 12, type: "delta" }, {
        tag: "delta", data: {
            name: "VB",
            source: {
                uri: "Example 1/VB/Class1.vb", range: {
                    start: { line: 1, character: 1 },
                    end: { line: 1, character: 1 }
                }, language: "Visual Basic"
            } satisfies SourceFile
        }
    }],
    [{ index: 13, type: "sigma" }, { tag: "sigma", data: { index: 5, source: undefined } }],
    [{ index: 15, type: "delta" }, {
        tag: "delta", data: {
            name: "BaseVB",
            source: {
                uri: "Example 1/VB/Class1.vb", range: {
                    start: { line: 1, character: 14 },
                    end: { line: 1, character: 20 }
                }, language: "Visual Basic"
            } satisfies SourceFile
        }
    }],
    [{ index: 16, type: "sigma" }, { tag: "sigma", data: { index: 7, source: undefined } }],
    [{ index: 17, type: "tau" }, { tag: "tau", data: "Integer" }],
    [{ index: 18, type: "tau" }, { tag: "tau", data: "Top" }],
])

console.log(JSON.stringify(constraints, null, 4))
console.log(JSON.stringify(Array.from(substitution.entries()), null, 4))
