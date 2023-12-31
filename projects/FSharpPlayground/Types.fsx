module Types

(** CREATE TUPLES**)

let point1 = (1.0, 2.0)

let gitHubStars = ("dotnet/fsharp", 2800)

(**GET DATA**)
fst gitHubStars // Get first item
snd gitHubStars // Get second item

let projectName, stars = gitHubStars

printfn $"{projectName}: {stars}"

let projectName2, _ = gitHubStars // Stars is ignored

(**DEFINE RECORD**)
type GitHubProjectEnum = { ProjectName: string; Stars: int; Status: bool }

(**CREATE RECORD**)
let fsharp1 =
    { ProjectName = "dotnet/fsharp"
      Stars = 2800 
      Status = false}

(**ACCESS RECORD DATA**)
printfn $"{fsharp1.ProjectName} has {fsharp1.Stars} stars"

(**UPDATE RECORD DATA**)
let updatedFSharp = { fsharp1 with Stars = 2801; ProjectName = "some" }

printfn $"{updatedFSharp.ProjectName} now has {updatedFSharp.Stars} stars"

(**RECORD MEMBERS**)
type GitHubProjectWithMember =
    { ProjectName: string
      Stars: int }
    member this.GetUrl () =
        $"https://github.com/{this.ProjectName}"

    member _.GetCodeOwner () = "dotnet"

let fsharpProj =
    { ProjectName = "dotnet/fsharp"
      Stars = 2800 }

fsharpProj.GetUrl()
fsharpProj.GetCodeOwner()

type StringGitHubProject =
    { ProjectName: string
      Stars: int
      State: string }

let fakeProject =
    { ProjectName = "Amazing Project"
      Stars = 0
      State = "asdfwerq" //No enforcement
    }

(**DEFINE DISCRIMINATED UNION*)
type ProjectState =
    | Archived
    | Active of {| Maintainer: string |}

type GitHubProject =
    { ProjectName: string
      Stars: int
      State: ProjectState }

(**USE DISCRIMINATED UNION*)
// let (anotherFakeProject: GitHubProject) =
//     { ProjectName = "Other amazing project"
//       Stars = 0
//       State = "asdfasdf" }

let corefxlab =
    { ProjectName = "corefxlab"
      Stars = 1500
      State = Archived }

let fsharp =
    { ProjectName = "dotnet/fsharp"
      Stars = 2800
      State = Active {| Maintainer = "F# Team" |} }

