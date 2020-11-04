package myerrors

import "errors"

// var ErrDetailsExist = errors.New("DETAILS_EXIST")
// var ErrNIL = errors.New("NIL")
// var ErrNoRecordsDeleted = errors.New("no records were deleted")

var ErrOrderNotValidated = errors.New("order validation failed")

// var ErrNotDeliverableArea = errors.New("we dont deliver here yet")
// var ErrRestaurantClosed = errors.New("RESTAURANT CLOSED")

// // Error codes

// var ErrBasicCode = 1
// var ErrNoRecordsCode = 0

// var ErrOrderNotValidatedCode = 2

// var FileUploadErrorCode = 3

type MyErrors struct {
	ErrBasicCode             int
	ErrNoRecordsCode         int
	ErrOrderNotValidatedCode int
	FileUploadErrorCode      int
	ErrDetailsExist          error
	ErrNIL                   error
	ErrNoRecordsDeleted      error
	ErrOrderNotValidated     error
	ErrNotDeliverableArea    error
	ErrRestaurantClosed      error
}

func New() MyErrors {
	e := MyErrors{
		ErrBasicCode:             1,
		ErrNoRecordsCode:         0,
		ErrOrderNotValidatedCode: 2,
		FileUploadErrorCode:      3,
		ErrDetailsExist:          errors.New("DETAILS_EXIST"),
		ErrNIL:                   errors.New("NIL"),
		ErrNoRecordsDeleted:      errors.New("no records were deleted"),
		ErrOrderNotValidated:     errors.New("order validation failed"),
		ErrNotDeliverableArea:    errors.New("we dont deliver here yet"),
		ErrRestaurantClosed:      errors.New("RESTAURANT CLOSED"),
	}

	return e
}
