package shared

type Unrecognized struct {
	source source
	Text   string `json:"text"`
	//TODO: DirectTo string
}

func NewUnrecognized(path string, start uint, length uint, text string) Unrecognized {
	return Unrecognized{
		source: source{
			Path: path,
			locationRange: locationRange{
				Start:  start,
				Length: length,
			},
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
