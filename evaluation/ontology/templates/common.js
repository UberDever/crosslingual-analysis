function need(args, ...fs) {
    let zip = rows => rows[0].map((_, c) => rows.map(row => row[c]))
    let argsLenIsSameAsTypeChecks = _ => fs.length == args.length
    ensure('just', argsLenIsSameAsTypeChecks)
    for (let [f, arg] of zip([fs, args])) {
        ensure(arg, f)
    }
}

function isCounter(x) {
    return true
}

function isScope(x) {
    expected = ['name', 'index']
    allKeys = expected.every(v => Object.keys(x).includes(v))
    scope = x.name == '_'
    return allKeys && scope
}

function isIdentifier(x) {
    expected = ['name', 'path', 'start', 'length']
    return Object.keys(x).every(v => expected.includes(v))
}

function nonNull(x) {
    return x != null
}

function ensure(x, f) {
    if (!f(x)) {
        throw `Expected to ${JSON.stringify(x)} pass the ${f.name} check`
    }
}