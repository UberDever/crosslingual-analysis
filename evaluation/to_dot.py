#!/usr/local/bin/python3

import argparse
import sexpdata
import json


class SexprTraverser:
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
            SexprTraverser, predicate=inspect.isfunction))
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


class JsonTraverser:
    ignore_unmatched = False

    def __init__(self, ignore_unmatched):
        self.ignore_unmatched = ignore_unmatched

    def traverse_tree(self, constraints):
        import inspect
        methods = dict(inspect.getmembers(
            JsonTraverser, predicate=inspect.isfunction))
        tab = ''

        def constraints_to_str(name: str):
            particular = constraints[name[0].upper() + name[1:]]
            if not particular:
                return ''
            return tab + ''.join(
                methods['_' + name](self, u) for u in particular)
        constraints_names = [
            'usage',
            'resolution',
            'uniqueness',
            'typeDeclKnown',
            'typeDeclUnknown',
            'directEdge',
            'associationKnown',
            'nominalEdge',
            'subset',
            'equalKnown',
            'equalUnknown',
            'associationUnknown',
        ]

        graph = ""
        graph += 'digraph G {\n'
        for t in constraints_names:
            graph += constraints_to_str(t)
        graph += '}'
        return graph

    def _identifier(self, id):
        return f"'{id['name']}' at {id['path']}:{id['start']}:{id['length']}"

    def _type(self, t):
        match t:
            case {"tag": "ground", "name": name}:
                return name
            case {"tag": "application",
                  "app": {"constructor": {"name": name}, "args": [*args]}}:
                return f'{name}({", ".join(self._type(t) for t in args)})'

            case _: self._raise(t)

    def _variable(self, v):
        match v:
            case {"index": id, "name": name}:
                match name:
                    case "delta":
                        return f'δ {id}'
                    case "sigma":
                        return f'ς {id}'
                    case "tau":
                        return f'τ {id}'
                    case "_":
                        return str(id)
            case _:
                self._raise(v)

    def _names_collection(self, names):
        match names:
            case {"collection": collection, "scope": scope}:
                c = str(collection[0]).upper()
                s = self._variable(scope)
                return f'{c}({s})'
            case _: self._raise(args)

    def _usage(self, args):
        match args:
            case {"id": _, "identifier": identifier,
                  "usage": usage, "scope": scope}:
                i = self._identifier(identifier)
                s = self._variable(scope)
                if usage == 'declaration':
                    return f'"{s}" -> "{i}"\n'
                else:
                    return f'"{i}" -> "{s}"\n'
            case _: self._raise(args)

    def _resolution(self, args):
        match args:
            case {"id": _, "reference": identifier,
                  "declaration": decl}:
                i = self._identifier(identifier)
                d = self._variable(decl)
                return f'"{i}" -> "{d}" [style=dotted]\n'
            case _: self._raise(args)

    def _uniqueness(self, args):
        match args:
            case {"id": _, "names": names}:
                n = self._names_collection(names)
                return f'/* !{n} */\n'
            case _: self._raise(args)

    def _typeDeclKnown(self, args):
        match args:
            case {"id": _, "declaration": decl, "variable": tau}:
                d = self._identifier(decl)
                t = self._variable(tau)
                return f'/* {d} :: {t} */\n'
            case _: self._raise(args)

    def _typeDeclUnknown(self, args):
        match args:
            case {"id": _, "declaration": decl, "variable": tau}:
                d = self._variable(decl)
                t = self._variable(tau)
                return f'/* {d} :: {t} */\n'
            case _: self._raise(args)

    def _directEdge(self, args):
        match args:
            case {"id": _, "lhs": lhs, "rhs": rhs, "label": label}:
                s1 = self._variable(lhs)
                s2 = self._variable(rhs)
                return f'"{s1}" -> "{s2}" [label="{label}"]\n'
            case _: self._raise(args)

    def _associationKnown(self, args):
        match args:
            case {"id": _, "declaration": decl, "scope": scope}:
                d = self._identifier(decl)
                s = self._variable(scope)
                return f'"{d}" -> "{s}" [arrowhead=onormal]\n'
            case _: self._raise(args)

    def _associationUnknown(self, args):
        match args:
            case {"id": _, "declaration": decl, "scope": scope}:
                d = self._variable(decl)
                s = self._variable(scope)
                return f'"{d}" -> "{s}" [arrowhead=onormal]\n'
            case _: self._raise(args)

    def _nominalEdge(self, args):
        match args:
            case {"id": _, "reference": ref, "scope": scope, "label": label}:
                r = self._identifier(ref)
                s = self._variable(scope)
                return f'"{r}" -> "{s}" [label="{label}"]\n'
            case _: self._raise(args)

    def _subset(self, args):
        match args:
            case {"id": _, "lhs": lhs, "rhs": rhs}:
                ll = self._names_collection(lhs)
                rr = self._names_collection(rhs)
                return f'/* {ll} ⊆ {rr} */\n'
            case _: self._raise(args)

    def _equalKnown(self, args):
        match args:
            case {"id": _, "t1": lhs, "t2": rhs}:
                ll = self._variable(lhs)
                rr = self._type(rhs)
                return f'/* {ll} == {rr} */\n'
            case _: self._raise(args)

    def _equalUnknown(self, args):
        match args:
            case {"id": _, "t1": lhs, "t2": rhs}:
                ll = self._variable(lhs)
                rr = self._variable(rhs)
                return f'/* {ll} == {rr} */\n'
            case _: self._raise(args)

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
parser.add_argument(
    '-t', '--traverser', help='Use the following traverser type',
    choices=['sexpr', 'json'], default='json')

args = parser.parse_args()

match args.traverser:
    case "sexpr":
        with open(args.path, 'r') as file:
            text = file.read()
            tree = sexpdata.loads(text)
        out = SexprTraverser(ignore_unmatched=False).traverse_tree(tree)
    case "json":
        with open(args.path, 'r') as file:
            tree = json.load(file)
        out = JsonTraverser(ignore_unmatched=False).traverse_tree(tree)
    case _:
        raise Exception("Unreachable")


with open(args.output, 'w+') as file:
    file.write(out)
