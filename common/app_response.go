package common

type successResp struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successResp {
	return &successResp{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *successResp {
	return NewSuccessResponse(data, nil, nil)
}
