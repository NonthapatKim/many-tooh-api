package response

import "github.com/NonthapatKim/many-tooth-api/internal/core/constant"

type SuccessResponse struct {
	Code    string                  `json:"code" default:"0000"`
	Message constant.SuccessMessage `json:"message"`
}

var (
	SuccessResponse_Created = SuccessResponse{
		Code:    constant.SuccessCode,
		Message: constant.SuccessMessage_Created,
	}

	SuccessResponse_Updated = SuccessResponse{
		Code:    constant.SuccessCode,
		Message: constant.SuccessMessage_Updated,
	}

	SuccessResponse_Deleted = SuccessResponse{
		Code:    constant.SuccessCode,
		Message: constant.SuccessMessage_Deleted,
	}
)
