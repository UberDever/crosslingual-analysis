package analysis

import (
	"fmt"
	"prototype2/sexpr"
	"reflect"
	"sort"
)

/*
One import can be linked with one and only one export

Ontology:
	Type lattice
	Possible language pairs

Type lattice:
|-Any
|-{Nominal types} | Unit
|-None,
where Nominal types = Int | String | {Identifier}
					  ^ all this types considered unit aliases,
					  in a sense that they inhabited by only one value
					  and equal to themselves only
					  Names are handy to store analysis results only

Possible language pairs:
	List of Connection, where Connection = (lhs, rhs, semantic)

module = {
	imports: List import
	exports: List export
	lang: string
	priority: int
}

statement = import | export
statement = {
	type: Sexpr // A | A -> A | A * A | A + A
	value: Sexpr
	source: Sexpr // Presumably AST Node
}

import = {
	statement
}

export = {
	statement
}

intraLinks = List (statement, statement)
interLinks = List (import, export)
*/

type module struct {
	path       string
	imports    []import_
	exports    []export
	lang       string
	priority   int
	intralinks []struct {
		from statement
		to   statement
	}
}

type statement struct {
	T, V   Sexpr
	Source Sexpr
}

func (s statement) String() string {
	return s.T.String() + s.V.String() + s.Source.String()
}

type import_ struct {
	statement
}

type export struct {
	statement
}

func NewImport(T, V, Source Sexpr) import_ {
	return import_{statement{T: T, V: V, Source: Source}}
}
func NewExport(T, V, Source Sexpr) export {
	return export{statement{T: T, V: V, Source: Source}}
}

func Function(types ...any) Sexpr {
	ts := sexpr.S(types...)
	return sexpr.Cons("->", ts)
}
func Product(types ...any) Sexpr {
	ts := sexpr.S(types...)
	return sexpr.Cons("x", ts)
}
func Sum(types ...any) Sexpr {
	ts := sexpr.S(types...)
	return sexpr.Cons("+", ts)
}

func CompareTypes(lhs Sexpr, rhs Sexpr) bool {
	l, r := lhs.String(), rhs.String()
	_, _ = l, r
	if lhs.String() == "'Any' " || rhs.String() == "'Any' " {
		return true
	}
	if lhs.String() == "'None' " || rhs.String() == "'None' " {
		return false
	}
	if lhs.IsAtom() || rhs.IsAtom() {
		if reflect.TypeOf(lhs) != reflect.TypeOf(rhs) {
			return false
		}
		return lhs.Data == rhs.Data
	}

	return CompareTypes(sexpr.Car(lhs), sexpr.Car(rhs)) &&
		CompareTypes(sexpr.Cdr(lhs), sexpr.Cdr(rhs))
}

func CompareValues(lhs Sexpr, rhs Sexpr) bool {
	return lhs.String() == rhs.String()
}

type Interlink struct {
	from       import_
	fromModule *module
	to         export
	toModule   *module
}

func (l Interlink) String() string {
	semantic := "?"
	for _, link := range ontology {
		if link.lhs == l.fromModule.lang && link.rhs == l.toModule.lang {
			semantic = link.semantic
		}
	}

	imp := fmt.Sprintf("%s: %s ",
		sexpr.MinifySexpr(l.from.T.StringReadable()),
		sexpr.MinifySexpr(l.from.V.StringReadable()),
	)
	exp := fmt.Sprintf("%s: %s ",
		sexpr.MinifySexpr(l.to.T.StringReadable()),
		sexpr.MinifySexpr(l.to.V.StringReadable()),
	)
	return "Semantic: " + semantic + "\n" +
		"import in " + l.fromModule.path + "\n" +
		imp + "\nis satisfied by \n" +
		"export from " + l.toModule.path + "\n" +
		exp + "\n"
}

func Link(modules []module) []Interlink {
	links := []Interlink{}
	sort.Slice(modules, func(i, j int) bool {
		return modules[i].priority > modules[j].priority
	})
	for i := range modules {
		for j := i + 1; j < len(modules); j++ {
			lhs := modules[i]
			rhs := modules[j]
			if langsCompatible(lhs.lang, rhs.lang) {
				for _, imp := range lhs.imports {
					for _, exp := range rhs.exports {
						typesEqual := CompareTypes(imp.T, exp.T)
						valuesEqual := CompareValues(imp.V, exp.V)
						if typesEqual && valuesEqual {
							wasLinked := false
							for _, link := range links {
								if link.from == imp {
									wasLinked = true
									break
								}
							}
							if !wasLinked {
								links = append(links, Interlink{imp, &lhs, exp, &rhs})
							}
							break
						}
					}
				}
				for _, imp := range rhs.imports {
					for _, exp := range lhs.exports {
						typesEqual := CompareTypes(imp.T, exp.T)
						valuesEqual := CompareValues(imp.V, exp.V)
						if typesEqual && valuesEqual {
							wasLinked := false
							for _, link := range links {
								if link.from == imp {
									wasLinked = true
									break
								}
							}
							if !wasLinked {
								links = append(links, Interlink{imp, &rhs, exp, &lhs})
							}
							break
						}
					}
				}
			}
		}
	}
	return links
}

// Types and Type lattice are coded implicitly in type functions
var ontology = []struct {
	lhs, rhs, semantic string
}{
	{lhs: "C#", rhs: "JS", semantic: "do Http responce"},
	{lhs: "JS", rhs: "C#", semantic: "do Http request"},

	{lhs: "Sh", rhs: "Python", semantic: "lookup file in directory"},
	{lhs: "Sh", rhs: "C", semantic: "lookup file in directory"},
	{lhs: "Python", rhs: "Sh", semantic: "use file produced by shell command"},
	{lhs: "Python", rhs: "C", semantic: "call C function"},
	{lhs: "C", rhs: "Sh", semantic: "use file produced by shell command"},
	{lhs: "C", rhs: "Python", semantic: "export function for FFI call"},

	// NOTE: This semantic is provided by intermodule analysis
	// it is not interesting for me
	// {lhs: "C", rhs: "C", semantic: ""},
}

func langsCompatible(lhs, rhs string) bool {
	for _, link := range ontology {
		if link.lhs == lhs && link.rhs == rhs {
			return true
		}
	}
	return false
}
