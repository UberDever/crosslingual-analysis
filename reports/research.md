# Questions

# Insights

Как я понимаю, задача семантического анализа мультиязыкового кода делится на:
1. Парсинг кода
    - Парсинг AST
    - Использование другого представления
1. Извлечение информации
    - Извлечение признаков по известной семантической модели
    - Извлечение по принципу ad hoc
    - Сохранение в едином формате
    - Постобработка (связывание)
1. Использование информации
    - Дамп в машино-читаемый формат
    - Интеграция в виде плагина

# Sources

[Nemerle](http://nemerle.org/About) CLR managed language with extensive preprocessing stage, like lisp macros

[Razor](https://learn.microsoft.com/en-us/aspnet/core/mvc/views/razor?view=aspnetcore-6.0) is a markup syntax for embedding .NET based code into webpages

[Hime](https://github.com/cenotelie/hime) is the parser generator is a parser generator that targets the .Net platform, Java and Rust

[Multilingual Source Code Analysis: A Systematic Literature Review](https://ieeexplore.ieee.org/abstract/document/7953501)

[Static Code Analysis of Multilanguage Software Systems](https://arxiv.org/pdf/1906.00815.pdf)

[Lightweight Multilingual Software Analysis](https://arxiv.org/abs/1808.01210)

[Lightweight Call-Graph Construction for Multilingual Software Analysis](https://arxiv.org/abs/1808.01213)

[MLSA](https://github.com/MultilingualStaticAnalysis/MLSA)

[The Effectiveness of Supervised Machine Learning Algorithms in Predicting Software Refactoring](https://arxiv.org/abs/2001.03338)

[Machine Learning for Software refactoring](https://github.com/refactoring-ai/predicting-refactoring-ml)

[PEG Grammars](https://en.wikipedia.org/wiki/Parsing_expression_grammar)

[Cross-Language Code Search using Static and Dynamic Analyses](https://dl.acm.org/doi/pdf/10.1145/3468264.3468538)

[Lightweight Multi-Language Syntax Transformation
with Parser Parser Combinators](https://dl.acm.org/doi/pdf/10.1145/3314221.3314589)

[Semantic reasoning about the sea of nodes / 2.2.2023](https://www.researchgate.net/publication/323333737_Semantic_reasoning_about_the_sea_of_nodes)

[See-of-nodes / 2.2.2023](https://darksi.de/d.sea-of-nodes/)

[Deep Learning for Source Code Modeling and Generation: Models, Applications and Challenges](https://arxiv.org/pdf/2002.05442.pdf)

[An empirical analysis of the utilization of multiple programming languages in open source projects](https://dl.acm.org/doi/abs/10.1145/2745802.2745805)

[On multi-language software development, cross-language links and accompanying tools: a survey of professional software developers](https://link.springer.com/article/10.1186/s40411-017-0035-z)

[POLYCRUISE: A Cross-Language Dynamic Information Flow Analysis](https://chapering.github.io/pubs/sec22.pdf)

[Cross-Language Support Mechanisms Significantly Aid Software Development](https://link.springer.com/chapter/10.1007/978-3-642-33666-9_12#citeas)

[Mulang](https://mumuki.github.io/mulang/)

[Typing in lambda calculus, Lambda cube](https://en.m.wikipedia.org/wiki/Lambda_cube)

[Pragmatic evidence of cross-language link detection: A systematic literature review](https://www.sciencedirect.com/science/article/abs/pii/S0164121223002200)

When it comes to semantic information representation in the context of programming languages, there are several main kinds of models that are commonly used. These models aim to capture the meaning and structure of programs to enable various analysis and processing tasks. Here are some of the prominent models:

1. Abstract Syntax Tree (AST): An AST represents the hierarchical structure of a program by abstracting away the low-level details. It captures the syntax and the relationships between different components of the program, such as statements, expressions, and declarations. ASTs are widely used in compilers, static analyzers, and program transformations.
1. Control Flow Graph (CFG): A CFG represents the flow of control within a program. It models the different paths that the program can take during its execution. Nodes in the CFG represent basic blocks of code, and edges represent the control flow between these blocks. CFGs are used in program analysis and optimization, such as data-flow analysis and loop analysis.
1. Data Flow Graph (DFG): A DFG models the flow of data within a program. It captures the dependencies between different variables and expressions. Nodes in the DFG represent computations, and edges represent the flow of data between these computations. DFGs are useful for various analyses, such as reaching definitions analysis and data dependence analysis.
1. Type Hierarchy: In statically typed programming languages, type hierarchies represent the relationships between different types in the language. They capture concepts like inheritance, polymorphism, and type compatibility. Type hierarchies are crucial for type checking, type inference, and object-oriented analysis.
1. Semantic Networks: Semantic networks are graphical representations that capture the meaning and relationships between different entities in a program. They consist of nodes representing entities (e.g., classes, methods, variables) and edges representing relationships (e.g., inheritance, method invocations). Semantic networks can be used for various program understanding and analysis tasks.
1. Knowledge Graphs: Knowledge graphs provide a structured representation of information by capturing entities, their attributes, and relationships in a graph format. In the context of programming languages, knowledge graphs can be used to represent libraries, frameworks, APIs, and their relationships. They enable program analysis, recommendation systems, and other intelligent tooling.
1. Program Dependence Graph (PDG): A PDG represents the dependencies between program statements based on data and control flow. It captures both data dependencies (how data flows between statements) and control dependencies (how control decisions affect the execution). PDGs are used in program slicing, program comprehension, and optimization.
1. Call Graph: A call graph represents the calling relationships between functions or methods in a program. It shows how functions invoke other functions, allowing for the analysis of function interactions, program execution paths, and identifying entry points. Call graphs are utilized in various analyses like reaching definitions, interprocedural analysis, and software maintenance tasks.
1. Abstract Interpretation: Abstract interpretation is a framework for semantic analysis that involves constructing abstract models of program behavior. It aims to approximate the possible runtime states and behaviors of a program. Abstract interpretation can be used for program verification, bug detection, and program optimization.
1. Ontologies: Ontologies are formal representations of knowledge that define concepts, their properties, and relationships within a domain. In the context of programming languages, ontologies can capture language-specific concepts, libraries, frameworks, and their relationships. They provide a structured and shared understanding of a programming domain, facilitating program analysis, interoperability, and knowledge-based tooling.
1. Program Slice: A program slice represents a subset of a program that focuses on a particular computation or variable of interest. It captures the dependencies and statements that affect the computation or variable. Program slicing is useful for debugging, program comprehension, and reducing the complexity of analysis tasks.
1. Intermediate Representations (IR): Intermediate representations are language-independent representations of programs that are closer to machine-executable code than the high-level source code. They abstract away language-specific details and provide a common representation for program analysis and transformations. Examples of IRs include LLVM IR, Java bytecode, and Microsoft Intermediate Language (MSIL).

Interesting: AST, DFG, Call graph, Abstract interpretation, ontologies
Also: [Graphlist](https://blog.scitools.com/graphlist/)

Use-cases types:

1. Language Interoperability: use case is to ensure proper interoperability between different language components
    - Type Compatibility: Code analysis tools can analyze the type systems of different languages and check for type compatibility when passing data or invoking functions across language boundaries
    - Language-Specific Constructs: Different languages have their own idioms, data structures, and syntax. 
    - API Conformance: In a multi-language codebase, it's common to have components that expose APIs to interact with other language modules.
1. Dependency Analysis: Codebases with multiple languages often have dependencies between different components written in different languages
    - Cross-Language Dependency Tracking: Code analysis tools can examine the codebase and identify dependencies between different language modules
    - Missing Dependency Detection: When introducing changes to a multi-language codebase, it's essential to ensure that all necessary dependencies are correctly included
    - Mismatched Dependency Versions: In a multi-language codebase, different components may rely on specific versions of libraries or frameworks
1. Security Analysis: Analyzing code that combines multiple languages is crucial for detecting security vulnerabilities
1. Cross-Language Documentation Generation: Documentation is essential for understanding and maintaining codebases
1. Code Migration and Porting: When migrating or porting code from one language to another, code analysis plays a crucial role
1. Metrics
