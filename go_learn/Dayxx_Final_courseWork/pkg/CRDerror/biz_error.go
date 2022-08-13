package CRDerror

import "github.com/pkg/errors"

var (
	errWrongDate = errors.New("Invalid Year, Month or Day")
)
