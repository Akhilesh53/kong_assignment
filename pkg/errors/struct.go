package errors

type Error struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	StatusCode   int    `json:"status_code"`
}

func GetDefaultError() Error {
	return Error{}
}
