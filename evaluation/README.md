The following projects are used to evaluate the analyzer.

Each of the following projects has README that shortly describes the purpose of the project.

# Notation

[Paper from which notation is taken from](<../sources/A Constraint Language for Static semantic analysis based on scope graphs.pdf>)

This thing is temporary -- that is, this should be converted to typed protocol 
(probably ts one from [here](../reports/sem4/protocol.d.ts)). For now this would suffice as it is easier to write this by hand.

```lisp
; Declaration ::=
(Delta <id:int>)
(Decl <name:string> <id:int>)
; Refer ::=
(Ref <name:string> <id:int>)
; S ::=
(Sigma <id:int>)
<id:negative int> ; For the ease of reading, known scopes are marked as negative ints

; Scope-graph constraints ::=
(Reference Refer S) ; Refer is mentioned in S
(Declare S Declaration) ; Declaration is declared in S
(Edge S1 S2 Label) ; Specify a labeled link from S1 to S2
(Association Declaration S) ; Associates the S with the Declaration, used for modules and non-lexical scoping (this is a fact, used when we declare a module)
(NominalEdge S Refer Label) ; Specify that all declarations from Ref (rather from Declaration to which Refer resolves) visible in S under Label

; (Collections) N ::= 
(References S) ; references in S
(Declarations S) ; declarations in S
(VisibleRefs S) ; visible declarations in S

; Resolution constraints ::=
(Resolves Refer Declaration) ; Ref must resolve to Decl
(Associated Declaration S) ; Associated S with Declaration (this is a constraint, used when we access qualified names)
(Uniq N) ; Logical predicate that states that all declaraions are unique in N
(Subset N1 N2) ; N1 is subset of N2

; T ::=
(Tau <id:int>)
(Constructor (T, ..., T)) ; Specifies a type constructor Constructor with following types

; Typing
(Equal T1 T2) ; Specifies that types must be equal
(Typeof Declaration T) ; Specifies type of declaration
```

TODO: This buddies :) removed from [this](<sources/A Constraint Language for Static semantic analysis based on scope graphs.pdf>) qualified names. So we need to model them by ourselves. Fortunatelly, its not hard to do.

## Their explanation for notation
* <span style="color:green"> A declaration constraint s -> xD specifies that declaration xD belongs to scope s. </span>
* <span style="color:green"> A reference constraint xR -> s specifies that reference xR
belongs to scope s </span>
* <span style="color:lightblue"> A resolution constraint R |-> D specifies that a given reference
must resolve to a given declaration. Typically, the declaration
is specified as a declaration variable δ </span>
* <span style="color:lightblue"> A uniqueness constraint !N specifies that a given name collection N contains no duplicates. </span>
* <span style="color:white"> A declaration name collection D(s) is obtained by projecting
the identifiers from the set of declarations in scope s. </span>
* <span style="color:red"> A type declaration constraint D : T associates a type with a
declaration. This constraint is used in two flavors: associating
a type variable (τ) with a concrete declaration, or associating a type variable with a declaration variable </span>
* <span style="color:red"> A type equality constraint T ≡ T specifies that two types should be equal </span>
* <span style="color:green"> A direct edge constraint s1 -l-> s2 specifies a direct l-labeled
edge from scope s1 to s2.</span>
* <span style="color:green"> An association constraint xD -|> s specifies s as the associated
scope of declaration xD. </span>
* <span style="color:green"> A nominal edge constraint s -l-> xR specifies a nominal l-labeled edge from scope s to reference xR. </span>
* <span style="color:white"> A reference name collection R(s) denotes the multiset of reference identifiers of scope s. </span>
* <span style="color:white"> A visible name collection V(s) denotes the multiset of declaration identifiers that are visible from scope s (i.e., would be visible from a reference to the declared identifier in s). </span>
* <span style="color:lightblue"> A subset constraint N ⊂∼ N specifies that one name collection is included in another. </span>
* <span style="color:lightblue"> An association constraint D ~> S specifies that a given declaration has a given associated scope </span>

# Languages

Difficulty of the implementation of the analysis.
- Easy:
    * Python (used in 2, 4)
    * XML (used in 3)
    * Go (used in 1, 5)
    * JSON (used in 5)
    * HTML (used in 1)
- Medium:
    * Makefile (through shell, used in 2)
    * Shell (used in 4)
    * C# (used in 3)
    * VB (used in 3)
    * JS (used in 1)
- Hard:
    * Fortran (used in 2)
    * C++ (used in 4)

Assumed order of implementation (project numbers): 1, 5, 3, 4, 2

# Statistics

- Described paradigms: 10
- Described domains: 11
- Described scenarios: 19