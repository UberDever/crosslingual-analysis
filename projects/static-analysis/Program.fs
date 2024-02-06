// For more information see https://aka.ms/fsharp-console-apps

let upper = List.map 

let uppered = "Hello from F#"
               |> String.map (System.Char.ToUpper)
printfn "%s" uppered
