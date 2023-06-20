package analysis

import (
	"prototype2/sexpr"
	"prototype2/util"
)

/*
judgment = import | export
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
	terms: List of Sexpr, // zero-value: undefined value for all types

	// language of interest, represented by this node
	lang: string

	// ast node that bound to this import
	ast: Sexpr
}

*/

type judgment struct {
	isImport bool
	types    Sexpr
	terms    Sexpr
	lang     string
	ast      Sexpr
}

func NewImport(types Sexpr, terms Sexpr, lang string, ast Sexpr) judgment {
	return judgment{
		isImport: true,
		types:    types,
		terms:    terms,
		lang:     lang,
		ast:      ast,
	}
}

func NewExport(types Sexpr, terms Sexpr, lang string, ast Sexpr) judgment {
	return judgment{
		isImport: false,
		types:    types,
		terms:    terms,
		lang:     lang,
		ast:      ast,
	}
}

func (n judgment) IsImport() bool {
	return n.isImport
}

func (n judgment) IsExport() bool {
	return !n.isImport
}

func (n judgment) TypesAndTerms() util.Set[Sexpr] {
	Car := sexpr.Car
	Cdr := sexpr.Cdr

	s := util.NewSet(sexpr.Equals)
	l := n.types
	for !l.IsNil() {
		s.Add(Car(l))
		l = Cdr(l)
	}
	l = n.terms
	for !l.IsNil() {
		s.Add(Car(l))
		l = Cdr(l)
	}
	return s
}

func CanLink(import_ judgment, export judgment) bool {
	if !import_.IsImport() || !export.IsExport() {
		panic("Something went wrong")
	}

	needs := import_.TypesAndTerms()
	gives := export.TypesAndTerms()

	return !gives.Contains(needs)
}

func (n judgment) String() {

}
