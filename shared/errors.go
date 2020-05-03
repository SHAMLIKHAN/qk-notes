package shared

const (
	// DecodeErrorCode : DecodeError
	DecodeErrorCode = 1
	// ValidationErrorCode : ValidationError
	ValidationErrorCode = 2
	// DatabaseErrorCode : DatabaseError
	DatabaseErrorCode = 3
	// InputDataErrorCode : InputDataError
	InputDataErrorCode = 4
	// JWTErrorCode : JWTError
	JWTErrorCode = 5

	// DecodeError : Decode Error
	DecodeError = "Error while decoding JSON data"
	// ValidationError : Validation Error
	ValidationError = "Error while validating input data"
	// DatabaseError : Database Error
	DatabaseError = "Error while querying the DB"
	// InputDataError : Input Data Error
	InputDataError = "Error while handling input data"
	// JWTError : JWT Error
	JWTError = "Error while generating JWT token"
)
