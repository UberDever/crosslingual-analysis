namespace Some

module Modules1 = 
    type SomeType = {
        FirstField: int
        SecondField: float
    }

    type Another = SomeType * int

module Modules2 = 
    open Modules1

    let getFirst a =
        let some = fst a
        let first = some.FirstField
        first


module M = 
    open Modules1
    open Modules2

    let o = { FirstField = 5; SecondField = 5.9 }
    let t = (o, 2)
    getFirst t


