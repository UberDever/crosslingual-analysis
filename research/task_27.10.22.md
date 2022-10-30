# Sources

[Nemerle](http://nemerle.org/About) CLR managed language with extensive preprocessing stage, like lisp macros

[Razor](https://learn.microsoft.com/en-us/aspnet/core/mvc/views/razor?view=aspnetcore-6.0) is a markup syntax for embedding .NET based code into webpages

[Hime](https://github.com/cenotelie/hime) is the parser generator is a parser generator that targets the .Net platform, Java and Rust

# Tasks

- [ ] Setup C# development cycle with wsl
- [ ] Make simple generic C# console app
- [ ] Install Hime 
- [ ] Take a grasp on library
- [ ] Try to implement "Find all references" analysis on simple console app
- [ ] Make simple C# ASP.Net app with Razor pages
- [ ] Try to somehow analyze .cshtml page
- [ ] @Difficult Try to analyze web app with AJAX call to dotnet method

# Results
1. From channel [MichaelRyanClarkson](https://youtube.com/c/MichaelRyanClarkson) learned about Hindley-Millner type inference algorithm, turn out it is pretty simple
1. Watched [Nemerle talk](https://www.youtube.com/watch?v=HSPivYkQ2t4) from CLRium
    - Nemerle is multiparadigm CLR based language
    - Main feature is macros and metaprogramming, based on syntax
    - Macros have 3 phases of substitution (@Insight this enables interdependence and mutual recursion)
    - Macros just functions, that executed in compile time
1. Found F# parsing [toolchain](https://en.wikibooks.org/wiki/F_Sharp_Programming/Lexing_and_Parsing), need to investigate further
1. Completed F# microsoft [microcourse](https://www.youtube.com/c/dotNET/videos)
