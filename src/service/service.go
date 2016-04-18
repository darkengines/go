package service

type Service interface {
	Parameter() interface{}
	Process(interface{}) (interface{}, error)
	Name() string
}