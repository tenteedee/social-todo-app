package common

type SuccessResponse struct {
	Data   any `json:"data"`
	Paging any `json:"paging"`
	Filter any `json:"filter"`
}

func NewSuccessResponse(data any, paging any, filter any) *SuccessResponse {
	return &SuccessResponse{
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}

func SimpleSuccessResponse(data any) *SuccessResponse {
	return &SuccessResponse{
		Data:   data,
		Paging: nil,
		Filter: nil,
	}
}
