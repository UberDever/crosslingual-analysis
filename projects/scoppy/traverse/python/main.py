import ast
import sys

def walk(tree, f):
    f(tree)
    match tree:
        case ast.AST():
            for field in tree._fields:
                walk(getattr(tree, field), f)
        case list():
            for e in tree:
                walk(e, f)
        

if len(sys.argv) < 2:
    print("Please provide source code to traverse")
    exit(-1)

source = sys.argv[1]
root = ast.parse(source)
walk(root, lambda e: print(e))