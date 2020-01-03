package helpers

import "github.com/globalsign/mgo/bson"

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

func (h *CommonHelper) FillArrBson(inStructArr, outStructArr interface{}) {
	inStructArrData, err := bson.Marshal(inStructArr)
	if err != nil {
		panic(err.Error())
	}
	raw := bson.Raw{Kind: 4, Data: inStructArrData}
	raw.Unmarshal(outStructArr)
}
