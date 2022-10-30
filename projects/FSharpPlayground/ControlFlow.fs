module ControlFlow

// Conditional expressions
open System.IO

let processFile (fileName:string) = 
    let fileExtension = Path.GetExtension(fileName)

    if fileExtension = ".fs" then
        printfn "This is a source file"
    elif fileExtension = ".fsx" then
        printfn "This is a script"
    elif fileExtension = ".fsproj" then
        printfn "This is a build configuration file"
    else printfn "Can't process file"

processFile "hello.fsx"
processFile "app.fs"
processFile "README.md"

// Exception Handling
let divide x y = 
    try
        Some(x / y)
    with 
        | :? System.DivideByZeroException -> printfn "Can't divide by zero"; None
        | ex -> printfn "Some other exception occurred";None

divide 10 2 // Some
divide 1 0 //None

// For start to end
for i=1 to 3 do
    printfn $"{i}"

// Reverse
for i=3 downto 1 do
    printfn $"{i}"

// For in
let todoList = [ "Learn F#"; "Create app"; "Profit!" ]

for todo in todoList do
    printfn $"{todo.ToUpper()}"

// Collection generation
[for todo in todoList do todo.ToUpper()]

// While loop
open System

let mutable input = ""
while (input <> "q") do
    input <- Console.ReadLine()
    printfn $"{input}"

type Address = {HouseNumber:int; StreetName: string}
type PhoneNumber = {Code:int; Number:string}

type ContactMethod = 
    | PostalMail of Address
    | Email of string
    | VoiceMail of PhoneNumber
    | SMS of PhoneNumber

let sendMessage (message:string) (contactMethod:ContactMethod) = 
    match contactMethod with 
    | PostalMail {HouseNumber=h;StreetName=s} -> printfn $"Mailing {message} to {h} {s}"
    | Email e -> printfn $"Emailing {message} to {e}"
    | VoiceMail {Code=c; Number=n} -> printfn $"Left +{c}{n} a voicemail saying {message}"

let message = "Can't talk now, learning F#!"

PostalMail {HouseNumber=1428; StreetName="Elm Street"}
|> sendMessage message

Email "suggestions@contoso.com"
|> sendMessage message

VoiceMail {Code=1;Number="5555555"}
|> sendMessage message