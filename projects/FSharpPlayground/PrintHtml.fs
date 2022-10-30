module PrintHtml

open FSharp.Data

let html = Http.RequestString("https://google.com/")

printfn $"{html}"
