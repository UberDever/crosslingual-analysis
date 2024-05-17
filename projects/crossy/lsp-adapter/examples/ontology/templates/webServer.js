function declare_WebServer(args) {
    need(args, isCounter, isScope, isIdentifier)
    let [counter, scope, id] = args
    let app_json_scope = {
        "index": counter.fresh(),
        "name": "_"
    }
    let server_scope = counter.fresh()
    let get_request_scope = counter.fresh()
    let template = {
        "Usage": [
            {
                "id": counter.fresh(),
                "identifier": id,
                "usage": "declaration",
                "scope": scope,
            },
            {
                "id": counter.fresh(),
                "identifier": {
                    "name": "GET",
                    "path": "",
                    "start": 0,
                    "length": 0
                },
                "usage": "declaration",
                "scope": {
                    "index": server_scope,
                    "name": "_"
                }
            },
            {
                "id": counter.fresh(),
                "identifier": {
                    "name": "application/json",
                    "path": "",
                    "start": 0,
                    "length": 0
                },
                "usage": "declaration",
                "scope": {
                    "index": get_request_scope,
                    "name": "_"
                }
            }
        ],
        "AssociationKnown": [
            {
                "id": counter.fresh(),
                "declaration": id,
                "scope": {
                    "index": server_scope,
                    "name": "_"
                }
            },
            {
                "id": counter.fresh(),
                "declaration": {
                    "name": "GET",
                    "path": "",
                    "start": 0,
                    "length": 0
                },
                "scope": {
                    "index": get_request_scope,
                    "name": "_"
                }
            },
            {
                "id": counter.fresh(),
                "declaration": {
                    "name": "application/json",
                    "path": "",
                    "start": 0,
                    "length": 0
                },
                "scope": app_json_scope,
            }
        ]
    }
    return [template, app_json_scope]
}

function reference_WebServer(args) {
    need(args, isCounter, isScope, isIdentifier)
    let [counter, scope, id] = args
    let aux_scope = counter.fresh()
    let server_resolution_scope = {
        "index": aux_scope,
        "name": "_"
    }
    let template = {
        "Usage": [
            {
                "id": counter.fresh(),
                "identifier": id,
                "usage": "reference",
                "scope": scope
            }
        ],
        "DirectEdge": [
            {
                "id": counter.fresh(),
                "lhs": {
                    "index": aux_scope,
                    "name": "_"
                },
                "rhs": scope,
                "label": "parent"
            }
        ],
        "NominalEdge": [
            {
                "id": counter.fresh(),
                "scope": server_resolution_scope,
                "reference": id,
                "label": "import"
            }
        ]
    }
    return [template, server_resolution_scope]
}