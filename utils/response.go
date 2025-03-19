package utils

/*
omitempty = ini akan tidak menampilkan jika field nya tidak di ambil
*/
type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func BuildResponseSucces(message string, data any) Response {
	res := Response{
		Status:  true,
		Message: message,
		Data:    data,
	}

	return res
}

func BuildResponseFailed(message string) Response {
	res := Response{
		Status:  false,
		Message: message,
	}

	return res
}
