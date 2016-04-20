package json

import "encoding/json"
import "communication/http"
import "io"
import serviceModule "service"

func GetParameter(context *http.HttpContext) (parameter serviceModule.Query, err error) {
	request := context.Request
	contentLength := request.ContentLength
	content := make([]byte, contentLength)
	_, err = request.Body.Read(content)
	if err != nil && err != io.EOF {
		return
	}
	var rawQuery JsonQuery
	err = json.Unmarshal(content, &rawQuery)
	if err != nil {
		return
	}

	serviceProvider := new(serviceModule.ServiceProvider)
	service, err := serviceProvider.GetService(rawQuery.ServiceName)
	if err != nil {
		return
	}

	subParameter := service.Parameter()
	err = json.Unmarshal(rawQuery.RawParameter, &subParameter)
	if err != nil {
		return
	}

	parameter = serviceModule.Query{Service: service, Parameter: subParameter}
	return
}

func Serialize(item interface{}) ([]byte, error) {
	return json.Marshal(item)
}

