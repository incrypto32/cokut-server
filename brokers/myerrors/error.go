package myerrors

import "errors"

var ErrDetailsExist = errors.New("DETAILS_EXIST")
var ErrNIL = errors.New("NIL")
var ErrNoRecordsDeleted = errors.New("no records were deleted")
var ErrOrderNotValidated = errors.New("order validation failed")

// Error codes

var ErrBasicCode = 1
var ErrNoRecordsCode = 0
var ErrOrderNotValidatedCode = 2
var FileUploadErrorCode = 3

type MyErrors struct {
	ErrBasicCode             int
	ErrNoRecordsCode         int
	ErrOrderNotValidatedCode int
	FileUploadErrorCode      int

	ErrDetailsExist      error
	ErrNIL               error
	ErrNoRecordsDeleted  error
	ErrOrderNotValidated error
}
