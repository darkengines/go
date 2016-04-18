package json
import "service"
import "encoding/json"
import "communication/http"
import "composition"
import "io"

type HttpRequestQueryAdapter struct {

}

func(this *HttpRequestQueryAdapter) Handle(context interface{}) (error) {
	httpContext := context.(http.HttpContext)
	request := httpContext.Request
	response := httpContext.Response
	
	contentLength := request.ContentLength
	content := make([]byte, contentLength)
	_, err := request.Body.Read(content)
	if (err != nil && err != io.EOF) {
		return err
	}
	var rawQuery JsonQuery 
	err = json.Unmarshal(content, &rawQuery)
	if (err != nil) {
		return err
	}
	
	serviceProvider := new(service.ServiceProvider)
	service, err := serviceProvider.GetService(rawQuery.ServiceName)
	if (err != nil) {
		return err
	}
	
	parameter := service.Parameter()
	err = json.Unmarshal(rawQuery.RawParameter, &parameter)
	if (err != nil) {
		return err
	}
	
	result, err := service.Process(parameter)
	if (err != nil) {
		return err
	}
	
	bytes, err := json.Marshal(result)
	if (err != nil) {
		return err
	}
	
	_, err = response.Write(bytes)
	if (err != nil) {
		return err
	}
	
	return nil
}

func(this *HttpRequestQueryAdapter) CanHandle(data interface{}) (bool) {
	request, ok := data.(http.HttpContext)
	return ok && request.Request.Header.Get("content-type") == "application/json"
}

func init() {
	container := composition.GetInstance()
	container.Export(func() (*HttpRequestQueryAdapter, error) {
		return new(HttpRequestQueryAdapter), nil
	})
}