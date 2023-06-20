package analysis

import (
	"prototype2/sexpr"
	"prototype2/util"
)

/*
module = import | export
import = {
	// or just Sexpr, this is type declarations that being imported
	types: List of Sexpr,

	// or just Sexpr, this is value for this import
	// !! WHOLE TYPE OF IMPORT IS THEN intersection(types, terms)
	terms: List of Sexpr, // zero-value: any value for all types

	// language of interest, represented by this node
	lang: string

	// ast node that bound to this import
	ast: Sexpr
}
export = {
	// or just Sexpr, this is type declarations that being exported
	types: List of Sexpr,

	// or just Sexpr, this is value of this export
	// !! WHOLE TYPE OF EXPORT IS THEN unity(types, terms)
	terms: List of Sexpr, // zero-value: null value for all types

	// language of interest, represented by this node
	lang: string

	// ast node that bound to this import
	ast: Sexpr
}

*/

type module struct {
	isImport bool
	types    Sexpr
	terms    Sexpr
	lang     string
	ast      Sexpr
}

func NewImport(types Sexpr, terms Sexpr, lang string, ast Sexpr) module {
	return module{
		isImport: true,
		types:    types,
		terms:    terms,
		lang:     lang,
		ast:      ast,
	}
}

func NewExport(types Sexpr, terms Sexpr, lang string, ast Sexpr) module {
	return module{
		isImport: false,
		types:    types,
		terms:    terms,
		lang:     lang,
		ast:      ast,
	}
}

func (n module) IsImport() bool {
	return n.isImport
}

func (n module) IsExport() bool {
	return !n.isImport
}

func Compare(import_ module, export module) bool {
	if !import_.IsImport() || !export.IsExport() {
		panic("Something went wrong")
	}
	Car := sexpr.Car

	var needs util.Set[Sexpr]
	{
		types := util.NewSet(sexpr.Equals)
		for t := Car(import_.types); !t.IsNil(); t = Car(import_.types) {
			types.Add(t)
		}
		terms := util.NewSet(sexpr.Equals)
		for t := Car(import_.terms); !t.IsNil(); t = Car(import_.terms) {
			terms.Add(t)
		}
		needs = types.Intersect(terms)
	}

	var gives util.Set[Sexpr]
	{
		types := util.NewSet(sexpr.Equals)
		for t := Car(import_.types); !t.IsNil(); t = Car(import_.types) {
			types.Add(t)
		}
		terms := util.NewSet(sexpr.Equals)
		for t := Car(import_.terms); !t.IsNil(); t = Car(import_.terms) {
			terms.Add(t)
		}
		gives = types.Unity(terms)
	}

	// check if `needs` covered by `gives`
	return !needs.Intersect(gives).IsEmpty()
}

func (n module) String() {

}
