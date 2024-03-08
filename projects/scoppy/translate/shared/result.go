package shared

type Constraint interface {
	variantConstraint()
}

type decl struct {
	Name string `json:"name"`
}

func (d decl) variantConstraint() {}

func NewDecl(name string) decl {
	return decl{
		Name: name,
	}
}

type locationRange struct {
	Start  uint `json:"start"`
	Length uint `json:"length"`
}

type Unrecognized struct {
	Path string `json:"path"`
	locationRange
	Text string `json:"text"`
	// TODO: DirectTo string
}

func NewUnrecognized(path string, start uint, length uint, text string) Unrecognized {
	return Unrecognized{
		Path: path,
		locationRange: locationRange{
			Start:  start,
			Length: length,
		},
		Text: text,
	}
}

type result struct {
	Id           uint           `json:"id"`
	Constraints  []Constraint   `json:"constraints"`
	Unrecognized []Unrecognized `json:"unrecognized"`
}

func NewResult(id uint, constraints []Constraint, unrecognized []Unrecognized) result {
	return result{
		Id:           id,
		Constraints:  constraints,
		Unrecognized: unrecognized,
	}
}
