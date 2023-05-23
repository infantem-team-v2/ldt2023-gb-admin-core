package terrors

const (
	defaultStatusCode int = 500
)

var (
	defaultErrorMessage = externalMessage{
		Description:  "can't get error for your case",
		InternalCode: 399999,
	}

	externalMessagesMap = map[uint32]*serializedExternalMessage{}

	labels       = "job=error_logging statusCode=%d internalCode=%d externalMessage=%s internalError=%s stackTrace=%s"
	errorMessage = "Error: %s | StackTrace:\n%s\n"
)
