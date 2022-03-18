package vector

import (
	"errors"
)

var (
	ErrIncompatibleDimensions = errors.New("incompatible vectors dimensions. different number of components")
	ErrNot3DVector            = errors.New("vector is not 3D")
)
