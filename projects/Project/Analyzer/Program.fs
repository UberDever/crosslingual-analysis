module Analyzer

open Hime.Redist
open System

// let traverse (root: ASTNode) =
//     match root.Children with
//     | :? System.Collections.IEnumerable as children ->
//      Seq.iter (fun e -> Console.WriteLine e.ToString) children

[<EntryPoint>]
let main argv =
    let exp = "(2 * 3) * 4 - 2"
    let result = Project.Frontend.GetResult(exp)

    match result.Errors.Count with
    | 0 -> ()
    | _ -> Seq.iter (fun (e: ParseError) -> Console.WriteLine e.ToString) result.Errors

    0
