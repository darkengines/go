package service
import "composition"
import "github.com/ahmetalpbalkan/go-linq"

type QueryAdapter interface {
	Handle(interface{}) (error)
	CanHandle(interface{}) bool
}

func GetQueryAdapter(context interface{}) (queryAdapter QueryAdapter, err error) {
	container := composition.GetInstance()
	var queryAdapters []QueryAdapter
	err = container.Import(&queryAdapters)

	var abstractQueryAdapter interface{}
	abstractQueryAdapter, err = linq.From(queryAdapters).Single(func(queryAdapter linq.T) (bool, error) {
		return queryAdapter.(QueryAdapter).CanHandle(context), nil
	})
	queryAdapter = abstractQueryAdapter.(QueryAdapter) 
	return
}