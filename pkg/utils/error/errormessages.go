package error

// ErrorMessages is a map of error codes to error messages.
//

// ToDo: expand this map to include more error codes.
var errorCodeToUserMessages = map[errorCode]string{
	NoCode:       "No code for this error. Used for testing only.",
	ErrorPackage: "Occurred in the error package",
	BashPackage:  "Occurred in the bash package",
}
