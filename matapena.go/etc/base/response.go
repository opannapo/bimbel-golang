package base

// JSONRes ...
type JSONRes struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

// JSONError ...
type JSONError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
