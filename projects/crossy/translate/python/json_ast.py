import ast
import sys
import json
import ast2json
        
if len(sys.argv) < 2:
    print(f'USAGE: {sys.argv[0]} <source code>')
    sys.exit(1)

source = sys.argv[1]
root = ast.parse(source)
s = json.dumps(ast2json.ast2json(root), indent=4)
print(s)