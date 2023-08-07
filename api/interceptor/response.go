package interceptor

const (
	DEFAULT_ERROR_MSG         = "Something went wrong - (default)"
	DEFAULT_HTTP_ERROR_CODE   = 500
	DEFAULT_HTTP_SUCCESS_CODE = 200
)

type Response interface {
	Data() interface{}
	Success() bool
	Error() interface{}
	ErrorMessage() string
}

type response struct {
	ResData         interface{} `json:"data"`
	ResSuccess      bool        `json:"success"`
	ResError        interface{} `json:"error"`
	ResErrorMessage string      `json:"message"`
}

func CreateResponse(success bool, data interface{}, err interface{}, errorMessage string) Response {
	if !success && errorMessage == "" {
		errorMessage = DEFAULT_ERROR_MSG
	}
	return &response{
		ResData:         data,
		ResSuccess:      success,
		ResError:        err,
		ResErrorMessage: errorMessage,
	}
}

func (r *response) Data() interface{} {
	return r.ResData
}

func (r *response) Success() bool {
	return r.ResSuccess
}

func (r *response) Error() interface{} {
	return r.ResError
}

func (r *response) ErrorMessage() string {
	return r.ResErrorMessage
}