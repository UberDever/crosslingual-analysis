function constructors() {
    return [
        {
            "name": "Function",
            "variance": [
                "-",
                "+"
            ]
        }
    ]
}

function NewType(ctor, args) {
    let ctors = constructors()
    let c = ctors.find(e => e.name == ctor)
    if (c == undefined) {
        throw `Cannot find constructor ${ctor}`
    }

    if (c.variance.length != args.length) {
        throw `Constructor ${ctor} expected ${c.variance.length} args, but got ${args.length}`
    }

    return {
        ...c,
        "args": args
    }
}