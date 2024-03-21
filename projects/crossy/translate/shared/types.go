package shared

type variance string

const (
	VarianceCovariant     variance = "+"
	VarianceContravariant variance = "-"
	VarianceInvariant     variance = "="
)

type constructor struct {
	name string
	argsVariance []variance
}

// type application struct {
// 	lhs *
// }

type context struct {
	ground []string
	constructors []constructor
	subtypes []string
}

