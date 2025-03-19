package dto

type (
	PagginationResponse struct {
		Page             int `json:"page"`
		Size             int `json:"size"`
		TotalPage        int `json:"totalPage"`
		TotalData        int `json:"totalData"`
		TotalDataCurrent int `json:"totalDataCurrent"`
	}

	PagginationRequest struct {
		Filter string
		Offset int
		Limit  int
	}
)
