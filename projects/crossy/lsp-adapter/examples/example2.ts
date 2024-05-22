import { Constraint, UsageConstraint, DirectEdgeConstraint, AssociationConstraint, TypeDeclarationConstraint, uinteger, TypeEqualConstraint, SourceFile, ResolutionConstraint, Variable, Substitution } from "./protocol"

// TODO: test.py should reference `lib_dir/lib.so`
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
            name: "build.sh",
            source: { uri: "Example 2/build.sh", language: undefined } satisfies SourceFile
        },
        scope: { index: 2, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "build.sh",
            source: { uri: "Example 2/build.sh", language: undefined } satisfies SourceFile
        },
        scope: { index: 3, source: undefined }
    } satisfies AssociationConstraint,
    {
        identifier: {
            name: "lib.cpp",
            source: {
                uri: "Example 2/build.sh", range: {
                    start: { line: 2, character: 13 },
                    end: { line: 2, character: 20 }
                }, language: "Shell"
            } satisfies SourceFile
        },
        scope: { index: 3, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "lib.cpp",
            source: {
                uri: "Example 2/build.sh", range: {
                    start: { line: 2, character: 13 },
                    end: { line: 2, character: 20 }
                }, language: "Shell"
            } satisfies SourceFile
        },
        type: { index: 4, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 4, type: "tau" },
        rhs: "URI"
    } satisfies TypeEqualConstraint,
    {
        reference: {
            name: "lib.cpp",
            source: {
                uri: "Example 2/build.sh", range: {
                    start: { line: 2, character: 13 },
                    end: { line: 2, character: 20 }
                }, language: "Shell"
            } satisfies SourceFile
        },
        declaration: { index: 5, type: "delta" }
    } satisfies ResolutionConstraint,
    {
        declaration: { index: 5, type: "delta" },
        scope: { index: 6, type: "sigma" }
    } satisfies AssociationConstraint,
    {
        identifier: {
            name: "lib_dir/lib.so",
            source: {
                uri: "Example 2/build.sh", range: {
                    start: { line: 9, character: 4 },
                    end: { line: 9, character: 10 }
                }, language: "Shell"
            } satisfies SourceFile
        },
        scope: { index: 3, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "lib_dir/lib.so",
            source: {
                uri: "Example 2/build.sh", range: {
                    start: { line: 9, character: 4 },
                    end: { line: 9, character: 10 }
                }, language: "Shell"
            } satisfies SourceFile
        },
        type: { index: 7, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 7, type: "tau" },
        rhs: "URI"
    } satisfies TypeEqualConstraint,
    {
        declaration: {
            name: "lib_dir/lib.so",
            source: {
                uri: "Example 2/build.sh", range: {
                    start: { line: 9, character: 4 },
                    end: { line: 9, character: 10 }
                }, language: "Shell"
            } satisfies SourceFile
        },
        scope: { index: 9, source: undefined }
    } satisfies AssociationConstraint,
    {
        from: { index: 6, type: "sigma" },
        to: { index: 9, source: undefined },
        label: "import"
    } satisfies DirectEdgeConstraint, // this models scope sameness
    {
        identifier: {
            name: "lib.cpp",
            source: { uri: "Example 2/lib.cpp", language: undefined } satisfies SourceFile
        },
        scope: { index: 2, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "lib.cpp",
            source: { uri: "Example 2/lib.cpp", language: undefined } satisfies SourceFile
        },
        scope: { index: 10, source: undefined }
    } satisfies AssociationConstraint,
    {
        identifier: {
            name: "counter_new",
            source: {
                uri: "Example 2/lib.cpp", range: {
                    start: { line: 27, character: 14 },
                    end: { line: 27, character: 25 }
                }, language: "C++"
            } satisfies SourceFile
        },
        scope: { index: 10, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "counter_new",
            source: {
                uri: "Example 2/lib.cpp", range: {
                    start: { line: 27, character: 14 },
                    end: { line: 27, character: 25 }
                }, language: "C++"
            } satisfies SourceFile
        },
        type: { index: 11, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 11, type: "tau" },
        rhs: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                "Unit",
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                }
            ]
        }
    } satisfies TypeEqualConstraint,
    {
        identifier: {
            name: "counter_free",
            source: {
                uri: "Example 2/lib.cpp", range: {
                    start: { line: 32, character: 10 },
                    end: { line: 32, character: 22 }
                }, language: "C++"
            } satisfies SourceFile
        },
        scope: { index: 10, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "counter_free",
            source: {
                uri: "Example 2/lib.cpp", range: {
                    start: { line: 32, character: 10 },
                    end: { line: 32, character: 22 }
                }, language: "C++"
            } satisfies SourceFile
        },
        type: { index: 12, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 12, type: "tau" },
        rhs: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Unit"
            ]
        }
    } satisfies TypeEqualConstraint,
    {
        identifier: {
            name: "counter_get",
            source: {
                uri: "Example 2/lib.cpp", range: {
                    start: { line: 37, character: 9 },
                    end: { line: 37, character: 20 }
                }, language: "C++"
            } satisfies SourceFile
        },
        scope: { index: 10, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "counter_get",
            source: {
                uri: "Example 2/lib.cpp", range: {
                    start: { line: 37, character: 9 },
                    end: { line: 37, character: 20 }
                }, language: "C++"
            } satisfies SourceFile
        },
        type: { index: 13, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 13, type: "tau" },
        rhs: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Integer"
            ]
        }
    } satisfies TypeEqualConstraint,
    {
        identifier: {
            name: "counter_reset",
            source: {
                uri: "Example 2/lib.cpp", range: {
                    start: { line: 42, character: 10 },
                    end: { line: 42, character: 23 }
                }, language: "C++"
            } satisfies SourceFile
        },
        scope: { index: 10, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "counter_reset",
            source: {
                uri: "Example 2/lib.cpp", range: {
                    start: { line: 42, character: 10 },
                    end: { line: 42, character: 23 }
                }, language: "C++"
            } satisfies SourceFile
        },
        type: { index: 14, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 14, type: "tau" },
        rhs: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Unit"
            ]
        }
    } satisfies TypeEqualConstraint,
    {
        identifier: {
            name: "counter_inc",
            source: {
                uri: "Example 2/lib.cpp", range: {
                    start: { line: 47, character: 10 },
                    end: { line: 47, character: 21 }
                }, language: "C++"
            } satisfies SourceFile
        },
        scope: { index: 10, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "counter_inc",
            source: {
                uri: "Example 2/lib.cpp", range: {
                    start: { line: 47, character: 10 },
                    end: { line: 47, character: 21 }
                }, language: "C++"
            } satisfies SourceFile
        },
        type: { index: 15, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 15, type: "tau" },
        rhs: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Unit"
            ]
        }
    } satisfies TypeEqualConstraint,
    {
        identifier: {
            name: "test.py",
            source: { uri: "Example 2/test.py", language: undefined } satisfies SourceFile
        },
        scope: { index: 2, source: undefined },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "test.py",
            source: { uri: "Example 2/test.py", language: undefined } satisfies SourceFile
        },
        scope: { index: 16, source: undefined }
    } satisfies AssociationConstraint,
    {
        identifier: {
            name: "counter_new",
            source: {
                uri: "Example 2/test.py", range: {
                    start: { line: 35, character: 15 },
                    end: { line: 35, character: 26 }
                }, language: "Python 3"
            } satisfies SourceFile
        },
        scope: { index: 16, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "counter_new",
            source: {
                uri: "Example 2/test.py", range: {
                    start: { line: 35, character: 15 },
                    end: { line: 35, character: 26 }
                }, language: "Python 3"
            } satisfies SourceFile
        },
        type: { index: 17, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 17, type: "tau" },
        rhs: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                "Unit",
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
            ]
        }
    } satisfies TypeEqualConstraint,

    {
        identifier: {
            name: "counter_inc",
            source: {
                uri: "Example 2/test.py", range: {
                    start: { line: 38, character: 9 },
                    end: { line: 38, character: 9 }
                }, language: "Python 3"
            } satisfies SourceFile
        },
        scope: { index: 16, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "counter_inc",
            source: {
                uri: "Example 2/test.py", range: {
                    start: { line: 38, character: 9 },
                    end: { line: 38, character: 9 }
                }, language: "Python 3"
            } satisfies SourceFile
        },
        type: { index: 18, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 18, type: "tau" },
        rhs: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Unit",
            ]
        }
    } satisfies TypeEqualConstraint,

    {
        identifier: {
            name: "counter_get",
            source: {
                uri: "Example 2/test.py", range: {
                    start: { line: 41, character: 32 },
                    end: { line: 41, character: 43 }
                }, language: "Python 3"
            } satisfies SourceFile
        },
        scope: { index: 16, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "counter_get",
            source: {
                uri: "Example 2/test.py", range: {
                    start: { line: 41, character: 32 },
                    end: { line: 41, character: 43 }
                }, language: "Python 3"
            } satisfies SourceFile
        },
        type: { index: 19, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 19, type: "tau" },
        rhs: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Integer",
            ]
        }
    } satisfies TypeEqualConstraint,

    {
        identifier: {
            name: "counter_reset",
            source: {
                uri: "Example 2/test.py", range: {
                    start: { line: 42, character: 9 },
                    end: { line: 42, character: 22 }
                }, language: "Python 3"
            } satisfies SourceFile
        },
        scope: { index: 16, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "counter_reset",
            source: {
                uri: "Example 2/test.py", range: {
                    start: { line: 42, character: 9 },
                    end: { line: 42, character: 22 }
                }, language: "Python 3"
            } satisfies SourceFile
        },
        type: { index: 20, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 20, type: "tau" },
        rhs: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Unit",
            ]
        }
    } satisfies TypeEqualConstraint,

    {
        identifier: {
            name: "counter_free",
            source: {
                uri: "Example 2/test.py", range: {
                    start: { line: 49, character: 5 },
                    end: { line: 49, character: 17 }
                }, language: "Python 3"
            } satisfies SourceFile
        },
        scope: { index: 16, source: undefined },
        usage: "reference"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "counter_free",
            source: {
                uri: "Example 2/test.py", range: {
                    start: { line: 49, character: 5 },
                    end: { line: 49, character: 17 }
                }, language: "Python 3"
            } satisfies SourceFile
        },
        type: { index: 21, type: "tau" }
    } satisfies TypeDeclarationConstraint,
    {
        lhs: { index: 21, type: "tau" },
        rhs: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Unit",
            ]
        }
    } satisfies TypeEqualConstraint,
]

const substitution: Map<Variable, Substitution> = new Map<Variable, Substitution>([
    [{ index: 4, type: "tau" }, { tag: "tau", data: "URI" }],
    [{ index: 7, type: "tau" }, { tag: "tau", data: "URI" }],
    [{ index: 11, type: "tau" }, {
        tag: "tau", data: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                "Unit",
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                }
            ]
        }
    }],
    [{ index: 12, type: "tau" }, {
        tag: "tau", data: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Unit",
            ]
        }
    }],
    [{ index: 13, type: "tau" }, {
        tag: "tau", data: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Integer",
            ]
        }
    }],
    [{ index: 14, type: "tau" }, {
        tag: "tau", data: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Unit",
            ]
        }
    }],
    [{ index: 15, type: "tau" }, {
        tag: "tau", data: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Unit",
            ]
        }
    }],
    [{ index: 17, type: "tau" }, {
        tag: "tau", data: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                "Unit",
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                }
            ]
        }
    }],
    [{ index: 18, type: "tau" }, {
        tag: "tau", data: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Unit",
            ]
        }
    }],
    [{ index: 19, type: "tau" }, {
        tag: "tau", data: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Integer",
            ]
        }
    }],
    [{ index: 20, type: "tau" }, {
        tag: "tau", data: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Unit",
            ]
        }
    }],
    [{ index: 21, type: "tau" }, {
        tag: "tau", data: {
            name: "Function",
            variance: [
                "-",
                "+"
            ],
            args: [
                {
                    name: "Pointer",
                    variance: [
                        "+"
                    ],
                    args: [
                        "Counter",
                    ]
                },
                "Unit",
            ]
        }
    }],
    [{ index: 6, type: "sigma" }, { tag: "sigma", data: { index: 10, source: undefined } }],
    [{ index: 5, type: "delta" }, {
        tag: "delta", data: {
            name: "lib.cpp",
            source: { uri: "Example 2/lib.cpp", language: undefined } satisfies SourceFile
        }
    }],
])

console.log(JSON.stringify(constraints, null, 4))
console.log(JSON.stringify(Array.from(substitution.entries()), null, 4))
