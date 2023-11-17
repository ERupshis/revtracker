package errors

import (
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
)

// DatabaseErrorsToRetry errors to retry request to database.
var DatabaseErrorsToRetry = []error{
	errors.New(pgerrcode.UniqueViolation),
	errors.New(pgerrcode.ConnectionException),
	errors.New(pgerrcode.ConnectionDoesNotExist),
	errors.New(pgerrcode.ConnectionFailure),
	errors.New(pgerrcode.SQLClientUnableToEstablishSQLConnection),
	errors.New(pgerrcode.SQLServerRejectedEstablishmentOfSQLConnection),
	errors.New(pgerrcode.TransactionResolutionUnknown),
	errors.New(pgerrcode.ProtocolViolation),
}

var ErrRecordNotFound = fmt.Errorf("record wasn't found in db")
