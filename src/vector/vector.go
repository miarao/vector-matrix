package vector

import (
	"errors"
	"fmt"
	"strconv"
)

type Vector []float64

func New(val []interface{}) (Vector, error) {
	vec := make([]float64, len(val))

	for i, v := range val {
		switch v.(type) {

		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32:
			vec[i]  = v.(float64)
			break

		case string:
			f, err := strconv.ParseFloat(v.(string), 64)
			if err != nil {
				return nil, fmt.Errorf("string is not a numeric type. err: %w", err)
			}

			vec[i] = f
			break

		default:
			return nil, errors.New("invalid type. Must be either numeric type or string contains numeric")
		}
	}

	return vec, nil
}

func NewFromFloat64Slice(values []float64) Vector {
	v := make(Vector, len(values))
	copy(v, values)
	return v
}

func (vec Vector) Copy()  (Vector, error){
	v := make(Vector, len(vec))
	copy(v, vec)
	return v, nil
}
