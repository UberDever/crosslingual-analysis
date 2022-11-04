module Analyzer

open Hime.Redist
open System

let rec traverse (f: string -> unit) (node: ASTNode) =
    f (node.ToString())
    Seq.iter (traverse f) node.Children

[<EntryPoint>]
let main argv =
    let exp = "(2 * 3) * 4 - 2"
    let result = Project.Frontend.GetResult(exp)

    let errors = Seq.map (fun e -> e.ToString()) result.Errors
    Seq.iter (printfn "%s") errors

    traverse (printfn "%s") result.Root

    0
