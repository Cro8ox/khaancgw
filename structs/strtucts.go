package structs

type (
	Response struct {
		StatusCode int
		Body       ResponseBody
	}

	ResponseBody struct {
		Message string      `json:"message"`
		Body    interface{} `json:"body"`
	}

	Sorter map[string]string

	PaginationInput struct {
		Limit  int    `json:"limit"`
		Page   int    `json:"page"`
		Sorter Sorter `json:"sorter"`
	}

	CursorInput struct {
		Limit      int  `form:"limit" json:"limit"`
		PreviousID uint `form:"previous_id" json:"previous_id"`
	}

	PaginationResponse struct {
		Total int         `json:"total"`
		Items interface{} `json:"items"`
	}

	CursorResponse struct {
		HasNext bool        `json:"has_next"`
		Items   interface{} `json:"items"`
	}

	SuccessResponse struct {
		Success bool `json:"success"`
	}
)
