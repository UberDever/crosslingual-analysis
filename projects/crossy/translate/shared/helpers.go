package shared

func NewDeclarationConstraint(counter CounterService, decl Identifier, typ ground) Constraints {
	tau := NewVariable(counter.FreshForce(), BindingTau)
	return Constraints{
		TypeDeclKnown: []TypeDeclKnown{NewTypeDeclKnown(counter.FreshForce(), decl, tau)},
		EqualKnown:    []EqualKnown{NewEqualKnown(counter.FreshForce(), tau, typ)},
	}
}
