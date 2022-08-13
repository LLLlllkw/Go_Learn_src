package CRDerror

import "github.com/pkg/errors"

var (
	errInvalidDate   = errors.New("Invalid data")
	errInsertFailure = errors.New("Fail to insert")
	errQueryFailure  = errors.New("Fail to query")
	errDeleteFailure = errors.New("Fail to delete")
)
