//packege vector implements basic

package vector

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	UnitElement         = 1.0
	NegativeUnitElement = -1.0
)

type Vector []float64

type Operations interface {
	Scale(scalar float64)
	Copy()
	Dimension()
	Invert()
}

func NewZeroVector(len int) Vector {
	return make(Vector, len)
}

func New(val []interface{}) (Vector, error) {
	v := make([]float64, len(val))

	for i, value := range val {
		switch value.(type) {

		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64:
			v[i] = value.(float64)
			break

		case string:
			f, err := strconv.ParseFloat(value.(string), 64)
			if err != nil {
				return nil, fmt.Errorf("string is not a numeric type. err: %w", err)
			}

			v[i] = f
			break

		default:
			return nil, errors.New("invalid type. Must be either numeric type or string contains numeric")
		}
	}

	return v, nil
}

func (v Vector) Copy() Vector {
	vec := make(Vector, len(v))
	copy(vec, v)
	return vec
}

func (v Vector) Scale(scalar float64) Vector {
	vec := make(Vector, len(v))

	for i, val := range v {
		vec[i] = val * scalar
	}

	return vec
}

func (v Vector) Dimension() int {
	return len(v)
}

func MatchingDimensions(v Vector, u Vector) error {
	if v.Dimension() != u.Dimension() {
		return ErrIncompatibleDimensions
	}

	return nil
}

func (v Vector) Invert() Vector {
	vec := make(Vector, len(v))

	for i := range v {
		v[i] = NegativeUnitElement * v[i]
	}

	return vec
}
