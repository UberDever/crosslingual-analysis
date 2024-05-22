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
            name: "server.go",
            source: { uri: "Example 3/backend/server.go", language: undefined } satisfies SourceFile
        },
        scope: { index: 2, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "server.go",
            source: { uri: "Example 3/backend/server.go", language: undefined } satisfies SourceFile
        },
        scope: { index: 3, source: undefined }
    } satisfies AssociationConstraint,
    {
        identifier: {
            name: "localhost:3333/item",
            source: {
                uri: "Example 3/backend/server.go", range: {
                    start: { line: 20, character: 22 },
                    end: { line: 20, character: 27 }
                }, language: "Golang"
            } satisfies SourceFile
        },
        scope: { index: 3, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "localhost:3333/item",
            source: {
                uri: "Example 3/backend/server.go", range: {
                    start: { line: 20, character: 22 },
                    end: { line: 20, character: 27 }
                }, language: "Golang"
            } satisfies SourceFile
        },
        scope: { index: 4, source: undefined }
    } satisfies AssociationConstraint,
    {
        identifier: {
            name: "GET",
            source: {
                uri: "Example 3/backend/server.go", range: {
                    start: { line: 26, character: 15 },
                    end: { line: 26, character: 18 }
                }, language: "Golang"
            } satisfies SourceFile
        },
        scope: { index: 4, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "GET",
            source: {
                uri: "Example 3/backend/server.go", range: {
                    start: { line: 26, character: 15 },
                    end: { line: 26, character: 18 }
                }, language: "Golang"
            } satisfies SourceFile
        },
        scope: { index: 5, source: undefined }
    } satisfies AssociationConstraint,
    {
        identifier: {
            name: "application/json",
            source: {
                uri: "Example 3/backend/server.go", range: {
                    start: { line: 31, character: 21 },
                    end: { line: 31, character: 25 }
                }, language: "Golang"
            } satisfies SourceFile
        },
        scope: { index: 5, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "application/json",
            source: {
                uri: "Example 3/backend/server.go", range: {
                    start: { line: 31, character: 21 },
                    end: { line: 31, character: 25 }
                }, language: "Golang"
            } satisfies SourceFile
        },
        scope: { index: 6, source: undefined }
    } satisfies AssociationConstraint,
    {
        identifier: {
            name: "count",
            source: {
                uri: "Example 3/backend/server.go", range: {
                    start: { line: 10, character: 22 },
                    end: { line: 10, character: 27 }
                }, language: "Golang"
            } satisfies SourceFile
        },
        scope: { index: 6, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "count",
            source: {
                uri: "Example 3/backend/server.go", range: {
                    start: { line: 10, character: 22 },
                    end: { line: 10, character: 27 }
                }, language: "Golang"
            } satisfies SourceFile
        },
        type: { index: 7, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 7, type: "tau" },
        rhs: "Integer"
    } satisfies TypeEqualConstraint,
    {
        identifier: {
            name: "POST",
            source: {
                uri: "Example 3/backend/server.go", range: {
                    start: { line: 24, character: 15 },
                    end: { line: 24, character: 19 }
                }, language: "Golang"
            } satisfies SourceFile
        },
        scope: { index: 4, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "POST",
            source: {
                uri: "Example 3/backend/server.go", range: {
                    start: { line: 24, character: 15 },
                    end: { line: 24, character: 19 }
                }, language: "Golang"
            } satisfies SourceFile
        },
        scope: { index: 8, source: undefined }
    } satisfies AssociationConstraint,

    {
        identifier: {
            name: "script.js",
            source: { uri: "Example 3/frontend/script.js", language: undefined } satisfies SourceFile
        },
        scope: { index: 2, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "script.js",
            source: { uri: "Example 3/frontend/script.js", language: undefined } satisfies SourceFile
        },
        scope: { index: 9, source: undefined }
    } satisfies AssociationConstraint,
    {
        identifier: {
            name: "http://localhost:3333/item",
            source: {
                uri: "Example 3/frontend/script.js", range: {
                    start: { line: 3, character: 12 },
                    end: { line: 3, character: 38 }
                }, language: "JavaScript"
            } satisfies SourceFile
        },
        scope: { index: 9, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        reference: {
            name: "http://localhost:3333/item",
            source: {
                uri: "Example 3/frontend/script.js", range: {
                    start: { line: 3, character: 12 },
                    end: { line: 3, character: 38 }
                }, language: "JavaScript"
            } satisfies SourceFile
        },
        declaration: { index: 10, type: "delta" }
    } satisfies ResolutionConstraint,
    {
        declaration: { index: 10, type: "delta" },
        scope: { index: 11, type: "sigma" }
    } satisfies AssociationConstraint,
    {
        from: { index: 11, type: "sigma" },
        to: { index: 12, source: undefined },
        label: "import"
    } satisfies DirectEdgeConstraint,
    {
        identifier: {
            name: "POST",
            source: {
                uri: "Example 3/frontend/script.js", range: {
                    start: { line: 4, character: 18 },
                    end: { line: 4, character: 22 }
                }, language: "JavaScript"
            } satisfies SourceFile
        },
        scope: { index: 12, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,

    {
        identifier: {
            name: "http://localhost:3333/item",
            source: {
                uri: "Example 3/frontend/script.js", range: {
                    start: { line: 8, character: 12 },
                    end: { line: 8, character: 38 }
                }, language: "JavaScript"
            } satisfies SourceFile
        },
        scope: { index: 9, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        reference: {
            name: "http://localhost:3333/item",
            source: {
                uri: "Example 3/frontend/script.js", range: {
                    start: { line: 8, character: 12 },
                    end: { line: 8, character: 38 }
                }, language: "JavaScript"
            } satisfies SourceFile
        },
        declaration: { index: 13, type: "delta" }
    } satisfies ResolutionConstraint,
    {
        declaration: { index: 13, type: "delta" },
        scope: { index: 14, type: "sigma" }
    } satisfies AssociationConstraint,
    {
        from: { index: 14, type: "sigma" },
        to: { index: 15, source: undefined },
        label: "import"
    } satisfies DirectEdgeConstraint,

    {
        identifier: {
            name: "GET",
            source: {
                uri: "Example 3/frontend/script.js", range: {
                    start: { line: 9, character: 27 },
                    end: { line: 12, character: 10 }
                }, language: "JavaScript"
            } satisfies SourceFile
        },
        scope: { index: 15, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        reference: {
            name: "GET",
            source: {
                uri: "Example 3/frontend/script.js", range: {
                    start: { line: 9, character: 27 },
                    end: { line: 12, character: 10 }
                }, language: "JavaScript"
            } satisfies SourceFile
        },
        declaration: { index: 16, type: "delta" }
    } satisfies ResolutionConstraint,
    {
        declaration: { index: 16, type: "delta" },
        scope: { index: 17, type: "sigma" }
    } satisfies AssociationConstraint,
    {
        from: { index: 17, type: "sigma" },
        to: { index: 18, source: undefined },
        label: "import"
    } satisfies DirectEdgeConstraint,

    {
        identifier: {
            name: "application/json",
            source: {
                uri: "Example 3/frontend/script.js", range: {
                    start: { line: 11, character: 13 },
                    end: { line: 11, character: 35 }
                }, language: "JavaScript"
            } satisfies SourceFile
        },
        scope: { index: 18, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        reference: {
            name: "application/json",
            source: {
                uri: "Example 3/frontend/script.js", range: {
                    start: { line: 11, character: 13 },
                    end: { line: 11, character: 35 }
                }, language: "JavaScript"
            } satisfies SourceFile
        },
        declaration: { index: 19, type: "delta" }
    } satisfies ResolutionConstraint,
    {
        declaration: { index: 19, type: "delta" },
        scope: { index: 20, type: "sigma" }
    } satisfies AssociationConstraint,
    {
        from: { index: 20, type: "sigma" },
        to: { index: 21, source: undefined },
        label: "import"
    } satisfies DirectEdgeConstraint,
    {
        identifier: {
            name: "count",
            source: {
                uri: "Example 3/frontend/script.js", range: {
                    start: { line: 14, character: 100 },
                    end: { line: 14, character: 105 }
                }, language: "JavaScript"
            } satisfies SourceFile
        },
        scope: { index: 21, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "count",
            source: {
                uri: "Example 3/frontend/script.js", range: {
                    start: { line: 14, character: 100 },
                    end: { line: 14, character: 105 }
                }, language: "JavaScript"
            } satisfies SourceFile
        },
        type: { index: 22, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 22, type: "tau" },
        rhs: "Top"
    } satisfies TypeEqualConstraint,
]

const substitution: Map<Variable, Substitution> = new Map<Variable, Substitution>([
    [{ index: 7, type: "tau" }, { tag: "tau", data: "Integer" }],
    [{ index: 10, type: "delta" }, {
        tag: "delta", data: {
            name: "localhost:3333/item",
            source: {
                uri: "Example 3/backend/server.go", range: {
                    start: { line: 20, character: 22 },
                    end: { line: 20, character: 27 }
                }, language: "Golang"
            } satisfies SourceFile
        }
    }],
    [{ index: 11, type: "sigma" }, { tag: "sigma", data: { index: 4, source: undefined } }],
    [{ index: 13, type: "delta" }, {
        tag: "delta", data: {
            name: "localhost:3333/item",
            source: {
                uri: "Example 3/backend/server.go", range: {
                    start: { line: 20, character: 22 },
                    end: { line: 20, character: 27 }
                }, language: "Golang"
            } satisfies SourceFile
        }
    }],
    [{ index: 14, type: "sigma" }, { tag: "sigma", data: { index: 4, source: undefined } }],
    [{ index: 16, type: "delta" }, {
        tag: "delta", data: {
            name: "GET",
            source: {
                uri: "Example 3/backend/server.go", range: {
                    start: { line: 26, character: 15 },
                    end: { line: 26, character: 18 }
                }, language: "Golang"
            } satisfies SourceFile
        }
    }],
    [{ index: 17, type: "sigma" }, { tag: "sigma", data: { index: 5, source: undefined } }],
    [{ index: 19, type: "delta" }, {
        tag: "delta", data: {
            name: "application/json",
            source: {
                uri: "Example 3/backend/server.go", range: {
                    start: { line: 31, character: 21 },
                    end: { line: 31, character: 25 }
                }, language: "Golang"
            } satisfies SourceFile
        }
    }],
    [{ index: 20, type: "sigma" }, { tag: "sigma", data: { index: 6, source: undefined } }],
    [{ index: 22, type: "tau" }, { tag: "tau", data: "Top" }],
])


console.log(JSON.stringify(constraints, null, 4))
console.log(JSON.stringify(Array.from(substitution.entries()), null, 4))