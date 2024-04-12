function declare_WebServer(args) {
    let a = {
        "Usage": [
            {
                "id": 1,
                "identifier": {
                    "argument": -2
                },
                "usage": "declaration",
                "scope": {
                    "argument": -1
                }
            },
            {
                "id": 3,
                "identifier": {
                    "name": "GET",
                    "path": "",
                    "start": 0,
                    "length": 0
                },
                "usage": "declaration",
                "scope": {
                    "index": 0,
                    "name": "_"
                }
            },
            {
                "id": 5,
                "identifier": {
                    "name": "application/json",
                    "path": "",
                    "start": 0,
                    "length": 0
                },
                "usage": "declaration",
                "scope": {
                    "index": 1,
                    "name": "_"
                }
            }
        ],
        "AssociationKnown": [
            {
                "id": 2,
                "declaration": {
                    "argument": -2
                },
                "scope": {
                    "index": 0,
                    "name": "_"
                }
            },
            {
                "id": 4,
                "declaration": {
                    "name": "GET",
                    "path": "",
                    "start": 0,
                    "length": 0
                },
                "scope": {
                    "index": 1,
                    "name": "_"
                }
            },
            {
                "id": 6,
                "declaration": {
                    "name": "application/json",
                    "path": "",
                    "start": 0,
                    "length": 0
                },
                "scope": {
                    "index": 2,
                    "name": "_"
                }
            }
        ]
    }
    return a
}