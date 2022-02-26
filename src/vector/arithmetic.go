package vector

// implementing

func Equals(v Vector, u Vector) bool {
	if v.Dimension() != v.Dimension() {
		return false
	}

	for i := range v {
		if v[i] != u[i] {
			return false
		}
	}

	return true
}

func Add(v Vector, u Vector) (Vector, error) {
	//cannot add vectors with different dimensions
	if err := MatchingDimensions(v, u); err != nil {
		return nil, err
	}

	sumVec := make(Vector, v.Dimension())
	//adding corresponding components
	for i := range v {
		sumVec[i] = v[i] + u[i]
	}

	return sumVec, nil
}

// Subtract returns the difference between the subtrahend and the minuend.
// minuend - subtrahend = difference
func Subtract(minuend Vector, subtrahend Vector) (Vector, error) {
	subtrahend = subtrahend.Invert()

	return Add(minuend, subtrahend)
}
