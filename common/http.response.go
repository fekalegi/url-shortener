package common

type BaseResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Errors  interface{} `json:"errors,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type ResponseMeta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

func ErrorResponse(err string) BaseResponse {
	return BaseResponse{
		Message: err,
		Status:  false,
	}
}

func BadRequestResponse(errs interface{}) BaseResponse {
	return BaseResponse{
		Message: "Invalid body or query param",
		Status:  false,
		Errors:  errs,
	}
}

func SuccessResponseWithData(data interface{}, msg string) BaseResponse {
	return BaseResponse{
		Data:    data,
		Message: msg,
		Status:  true,
	}
}

func SuccessResponseNoData(msg string) BaseResponse {
	return BaseResponse{
		Message: msg,
		Status:  true,
	}
}

func SuccessPaginationResponse(data interface{}, msg string, meta ResponseMeta) BaseResponse {
	return BaseResponse{
		Data:    data,
		Message: msg,
		Status:  true,
		Errors:  nil,
		Meta:    meta,
	}
}
