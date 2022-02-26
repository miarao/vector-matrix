package vector

import (
	"errors"
)

var ErrIncompatibleDimensions = errors.New("incompatible vectors dimensions. different number of components")
