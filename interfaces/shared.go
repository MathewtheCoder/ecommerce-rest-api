package interfaces

func CreateDetailedErrorDto(key string, err error) map[string]interface{} {
	return map[string]interface{}{
		"success": false,
		"message": err.Error(),
		"errors":  err,
	}
}
