package localerror

type CodePair struct {
	HttpStatusCode   int
	ServiceErrorCode string
}

type StatusInfo struct {
	HttpStatusCode      int
	ServiceErrorCode    string
	ServiceErrorMessage string
}

func NewCodePair(httpStatusCode int, serviceErrorCode string) CodePair {
	return CodePair{
		HttpStatusCode:   httpStatusCode,
		ServiceErrorCode: serviceErrorCode,
	}
}

func NewStatusInfo(httpStatusCode int, serviceErrorCode string, serviceErrorMessage string) StatusInfo {
	return StatusInfo{
		HttpStatusCode:      httpStatusCode,
		ServiceErrorCode:    serviceErrorCode,
		ServiceErrorMessage: serviceErrorMessage,
	}
}
