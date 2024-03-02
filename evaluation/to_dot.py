#!/usr/local/bin/python3

import argparse
import sexpdata


class Traverser:
    ignore_unmatched = False

    def __init__(self, ignore_unmatched):
        self.ignore_unmatched = ignore_unmatched

    def _AND(self, args):
        graph = ""
        graph += 'digraph G {\n'
        for arg in args:
            graph += '    ' + self.traverse_tree(arg)
        graph += '\n}'
        return graph

    def _Edge(self, args):
        match args:
            case [scope1, scope2, label]:
                s1 = self.traverse_tree(scope1)
                s2 = self.traverse_tree(scope2)
                return f'"{s1}" -> "{s2}" [label="{label}"]\n'
            case _: self._raise(args)

    def _Decl(self, args):
        match args:
            case [name, id]:
                return f'{name}:{id}'
            case _: self._raise(args)

    def _Ref(self, args):
        match args:
            case [name, id]:
                return f'{name}:{id}'
            case _: self._raise(args)

    def _Delta(self, args):
        match args:
            case [id]:
                return f'δ {id}'
            case _: self._raise(args)

    def _Tau(self, args):
        match args:
            case [id]:
                return f'τ {id}'
            case _: self._raise(args)

    def _Sigma(self, args):
        match args:
            case [id]:
                return f'ς {id}'
            case _: self._raise(args)

    def _Typeof(self, args):
        match args:
            case [decl, tau]:
                d = self.traverse_tree(decl)
                t = self.traverse_tree(tau)
                return f'/* {d} :: {t} */\n'
            case _: self._raise(args)

    def _Equals(self, args):
        match args:
            case [tau, typ]:
                t1 = self.traverse_tree(tau)
                t2 = self.traverse_tree(typ)
                return f'/* {t1} == {t2} */\n'
            case _: self._raise(args)

    def _Constructor(self, args):
        if isinstance(args, list):
            arguments = ", ".join(self.traverse_tree(arg) for arg in args[1:])
            return f'{args[0]}<{arguments}>'
        else:
            return args

    def _Resolves(self, args):
        match args:
            case [ref, decl]:
                r = self.traverse_tree(ref)
                d = self.traverse_tree(decl)
                return f'"{r}" -> "{d}" [style=dotted]\n'
            case _: self._raise(args)

    def _Reference(self, args):
        match args:
            case [ref, scope]:
                r = self.traverse_tree(ref)
                s = self.traverse_tree(scope)
                return f'"{r}" -> "{s}"\n'
            case _: self._raise(args)

    def _Declare(self, args):
        match args:
            case [scope, decl]:
                s = self.traverse_tree(scope)
                d = self.traverse_tree(decl)
                return f'"{s}" -> "{d}"\n'
            case _: self._raise(args)

    def _Association(self, args):
        match args:
            case [decl, scope]:
                s = self.traverse_tree(scope)
                d = self.traverse_tree(decl)
                return f'"{d}" -> "{s}" [arrowhead=onormal]\n'
            case _: self._raise(args)

    def _Associated(self, args):
        match args:
            case [decl, scope]:
                s = self.traverse_tree(scope)
                d = self.traverse_tree(decl)
                return f'"{d}" -> "{s}" [arrowhead=onormal] #assoc\n'
            case _: self._raise(args)

    def _NominalEdge(self, args):
        match args:
            case [scope, ref, label]:
                r = self.traverse_tree(ref)
                s = self.traverse_tree(scope)
                return f'"{s}" -> "{r}" [label="{label}"]\n'
            case _: self._raise(args)

    def traverse_tree(self, tree):
        import inspect
        methods = dict(inspect.getmembers(
            Traverser, predicate=inspect.isfunction))
        match tree:
            case [sexpdata.Symbol(s), *args]:
                return methods['_' + s](self, args)
            case s:
                if isinstance(s, int):
                    name = chr(65 + 1 - s) if s < 0 else "???"
                    return f"Scope {name}"
                elif isinstance(s, str):
                    return s
                self._raise(tree)

    def _raise(self, tree):
        if not self.ignore_unmatched:
            raise Exception("Unreachable " + str(tree))


parser = argparse.ArgumentParser(usage="""
Convert sexpr notation from the README in project
files to dot notation for easy visualisation
""")

parser.add_argument(
    '-p', '--path', help='Path to sexpr notation', required=True)
parser.add_argument(
    '-o', '--output', help='Path to resulting dot notation', required=True)

args = parser.parse_args()

with open(args.path, 'r') as file:
    text = file.read()
    tree = sexpdata.loads(text)

out = Traverser(ignore_unmatched=False).traverse_tree(tree)

with open(args.output, 'w+') as file:
    file.write(out)
