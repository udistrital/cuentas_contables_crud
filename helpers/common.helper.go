package helpers

type CommonHelper struct{}

func (h *CommonHelper) DefaultResponse(code int, err error, info interface{}) map[string]interface{} {
	response := make(map[string]interface{})

	response["Code"] = code
	response["Message"] = nil
	response["Body"] = info

	if err != nil {
		response["Message"] = err.Error()
		response["Type"] = "error"
		response["Code"] = 500
	}

	return response
}
