package error

type Code string

const (
	// Common
	CodeUnknown      Code = "UNKNOWN"
	CodeInvalidInput Code = "INVALID_INPUT"
	CodeUnauthorized Code = "UNAUTHORIZED"
	CodeForbidden    Code = "FORBIDDEN"
	CodeNotFound     Code = "NOT_FOUND"
	CodeConflict     Code = "CONFLICT"
	CodeSystem       Code = "SYSTEM"
	CodeBusiness     Code = "BUSINESS"
	// external
	CodeDependencyUnavailable Code = "DEPENDENCY_UNAVAILABLE"
)
