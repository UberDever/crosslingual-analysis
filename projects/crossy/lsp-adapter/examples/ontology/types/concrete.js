function basicTypes() {
    return [
        "Numeric",
        "Integer",
        "Bool",
        "String",
        "Unit",
        "Top",
    ]
}

function ConcreteType(t) {
    if (!basicTypes().includes(t)) {
        throw `Cannot find ${t} type`
    }
    return t
}

function subtypes() {
    let types = basicTypes()
    let top_subtypes = types.map(el => [el, "Top"])
        .filter(el => { let [fst, snd] = el; return !(fst == "Top" && snd == "Top") })
    return [
        ["Integer", "Numeric"],
        ...top_subtypes
    ]
}

function SubtypesOf(type) {
    return subtypes.map(el => el[1] == type)
}

function SupertypesOf(type) {
    return subtypes.map(el => el[0] == type)
}