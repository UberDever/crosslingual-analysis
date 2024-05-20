import { Constraint, UsageConstraint, DirectEdgeConstraint, AssociationConstraint, TypeDeclarationConstraint, uinteger, TypeEqualConstraint, SourceFile, ResolutionConstraint, Variable, Substitution } from "./protocol"

const constraints: Constraint[] = [
    {
        identifier: { name: "OS", source: undefined },
        scope: { index: 0, type: "scope" },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: { name: "OS", source: undefined },
        scope: { index: 1, type: "scope" }
    } satisfies AssociationConstraint,
    {
        identifier: { name: "Filesystem", source: undefined },
        scope: { index: 1, type: "scope" },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: { name: "Filesystem", source: undefined },
        scope: { index: 2, type: "scope" }
    } satisfies AssociationConstraint,
    {
        from: { index: 2, type: "scope" },
        to: { index: 1, type: "scope" },
        label: "parent"
    } satisfies DirectEdgeConstraint,

    {
        identifier: {
            name: "build.sh",
            source: { uri: "Example 2/build.sh", language: undefined } satisfies SourceFile
        },
        scope: { index: 2, type: "scope" },
        usage: "declaration"
    } satisfies UsageConstraint,
    {
        declaration: {
            name: "build.sh",
            source: { uri: "Example 2/build.sh", language: undefined } satisfies SourceFile
        },
        scope: { index: 3, type: "scope" }
    } satisfies AssociationConstraint,
]

const substitution: Map<Variable, Substitution> = new Map<Variable, Substitution>([])

console.log(JSON.stringify(constraints, null, 4))
console.log(JSON.stringify(Array.from(substitution.entries()), null, 4))
