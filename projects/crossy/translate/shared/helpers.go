package shared

func DeepCopyJSON(v any) any {
	if m, ok := v.(map[string]any); ok {
		return deepCopyMap(m)
	} else if s, ok := v.([]any); ok {
		return deepCopySlice(s)
	}
	return v
}

// DeepCopy will create a deep copy of this map. The depth of this
// copy is all inclusive. Both maps and slices will be considered when
// making the copy.
func deepCopyMap(m map[string]any) map[string]any {
	result := map[string]any{}

	for k, v := range m {
		// Handle maps
		mapvalue, isMap := v.(map[string]any)
		if isMap {
			result[k] = deepCopyMap(mapvalue)
			continue
		}

		// Handle slices
		slicevalue, isSlice := v.([]any)
		if isSlice {
			result[k] = deepCopySlice(slicevalue)
			continue
		}

		result[k] = v
	}

	return result
}

// DeepCopy will create a deep copy of this slice. The depth of this
// copy is all inclusive. Both maps and slices will be considered when
// making the copy.
func deepCopySlice(s []any) []any {
	result := []any{}

	for _, v := range s {
		// Handle maps
		mapvalue, isMap := v.(map[string]any)
		if isMap {
			result = append(result, deepCopyMap(mapvalue))
			continue
		}

		// Handle slices
		slicevalue, isSlice := v.([]any)
		if isSlice {
			result = append(result, deepCopySlice(slicevalue))
			continue
		}

		result = append(result, v)
	}

	return result
}

func NewDeclarationConstraint(counter CounterService, decl Identifier, ty Type) Constraints {
	tau := NewVariable(counter.FreshForce(), BindingTau)
	return Constraints{
		TypeDeclKnown: []TypeDeclKnown{NewTypeDeclKnown(counter.FreshForce(), decl, tau)},
		EqualKnown:    []EqualKnown{NewEqualKnown(counter.FreshForce(), tau, ty)},
	}
}
