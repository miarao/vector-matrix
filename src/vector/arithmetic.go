package vector

import (
	"math"
)

// implementing

func (v *Vector) Add(u *Vector) (Vector, error) {
	//cannot add vectors with different dimensions
	if err := v.matchDim(u); err != nil {
		return nil, err
	}

	sumVec := NewZeroVector(v.Dimension())
	//adding corresponding components
	for i := range *v {
		sumVec[i] = (*v)[i] + (*u)[i]
	}

	return sumVec, nil
}

// Subtract returns the difference between the subtrahend and the minuend - v.
// minuend - subtrahend = difference
func (v *Vector) Subtract(u *Vector) (Vector, error) {
	subtrahend := u.Invert()

	return v.Add(&subtrahend)
}

func (v *Vector) Distance(u *Vector) (float64, error) {
	diff, err := v.Subtract(u)
	if err != nil {
		return 0, err
	}

	res, _ := diff.Dot(&diff)

	return math.Sqrt(res), err
}

func (v *Vector) Dot(u *Vector) (float64, error) {
	if err := v.matchDim(u); err != nil {
		return 0, err
	}

	var prdct float64

	for i := range *v {
		prdct += (*v)[i] * (*u)[i]
	}

	return prdct, nil
}

func (v *Vector) Cross(u *Vector) (Vector, error) {
	if err := v.matchDim(u); err != nil {
		return nil, err
	} else if v.Dimension() != 3 {
		return nil, ErrNot3DVector
	}

	//var res []interface{}

	// calculate the cross product components
	x := (*v)[2]*(*u)[3] - (*u)[2]*(*v)[3]
	y := (*u)[1]*(*v)[3] - (*v)[1]*(*u)[3]
	z := (*v)[1]*(*u)[2] - (*u)[1]*(*v)[2]

	return New(x, y, z), nil
}
