package responses

import (
	"net/http"
)

type CommonResponse struct {
	Status  int    `json:"status"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func ResponseNoContent() CommonResponse {
	result := CommonResponse{
		Status:  http.StatusNoContent,
		Message: MESSAGE_NO_CONTENT,
		Data:    nil,
	}

	return result
}

func ResponseCreated(data any) CommonResponse {
	result := CommonResponse{
		Status:  http.StatusCreated,
		Message: MESSAGE_CREATED,
		Data:    data,
	}

	return result
}

func ResponseSuccess(data any) CommonResponse {
	result := CommonResponse{
		Status:  http.StatusOK,
		Message: MESSAGE_SUCCESS,
		Data:    data,
	}

	return result
}

// //////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////
// //////////////////////////////////////////////////////////////
func ForbiddenResponse(data any) CommonResponse {
	result := CommonResponse{
		Status:  http.StatusForbidden,
		Message: "forbidden",
		Data:    data,
	}
	return result
}
