package mycode

import "github.com/gogf/gf/v2/errors/gcode"

var (
	CodeNil                       = gcode.New(-1, "", nil)                             // No error code specified.
	CodeOK                        = gcode.New(0, "OK", nil)                            // It is OK.
	CodeInternalError             = gcode.New(50, "Internal Error", nil)               // An error occurred internally.
	CodeValidationFailed          = gcode.New(51, "Validation Failed", nil)            // Data validation failed.
	CodeDbOperationError          = gcode.New(52, "Database Operation Error", nil)     // Database operation error.
	CodeInvalidParameter          = gcode.New(53, "Invalid Parameter", nil)            // The given parameter for current operation is invalid.
	CodeMissingParameter          = gcode.New(54, "Missing Parameter", nil)            // Parameter for current operation is missing.
	CodeInvalidOperation          = gcode.New(55, "Invalid Operation", nil)            // The function cannot be used like this.
	CodeInvalidConfiguration      = gcode.New(56, "Invalid Configuration", nil)        // The configuration is invalid for current operation.
	CodeMissingConfiguration      = gcode.New(57, "Missing Configuration", nil)        // The configuration is missing for current operation.
	CodeNotImplemented            = gcode.New(58, "Not Implemented", nil)              // The operation is not implemented yet.
	CodeNotSupported              = gcode.New(59, "Not Supported", nil)                // The operation is not supported yet.
	CodeOperationFailed           = gcode.New(60, "Operation Failed", nil)             // I tried, but I cannot give you what you want.
	CodeNotAuthorized             = gcode.New(61, "Not Authorized", nil)               // Not Authorized.
	CodeSecurityReason            = gcode.New(62, "Security Reason", nil)              // Security Reason.
	CodeServerBusy                = gcode.New(63, "Server Is Busy", nil)               // Server is busy, please try again later.
	CodeUnknown                   = gcode.New(64, "Unknown Error", nil)                // Unknown error.
	CodeNotFound                  = gcode.New(65, "Not Found", nil)                    // Resource does not exist.
	CodeInvalidRequest            = gcode.New(66, "Invalid Request", nil)              // Invalid request.
	CodeNecessaryPackageNotImport = gcode.New(67, "Necessary Package Not Import", nil) // It needs necessary package import.
	CodeBusinessValidationFailed  = gcode.New(300, "Business Validation Failed", nil)  // Business validation failed.
)

func GetErrorCode(err interface{}) gcode.Code {
	if err == nil {
		return nil
	}

	switch c := err.(type) {
	case gcode.Code:
		return c
	default:
		return CodeNil
	}

}
