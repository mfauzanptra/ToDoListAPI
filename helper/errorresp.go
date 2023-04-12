package helper

func ErrResp(status, message string) interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": message,
	}
}
