package service

import "composition"
import "github.com/ahmetalpbalkan/go-linq"

type ServiceProvider struct {
}

func (this *ServiceProvider) services() (services []Service, err error) {
	exportProvider := composition.GetInstance()
	var abstractServices []Service
	err = exportProvider.Import(&abstractServices)
	services = make([]Service, cap(abstractServices))
	for index, abstractService := range abstractServices {
		services[index] = abstractService.(Service)
	}
	return services, err
}

func (this *ServiceProvider) GetService(name string) (service Service, err error) {
	services, servicesError := this.services()
	if (servicesError != nil) {
		return nil, err
	}
	var abstractService interface{}
	abstractService, err = linq.From(services).Single(func(service linq.T) (bool, error) {
		return service.(Service).Name() == name, nil
	})
	service = abstractService.(Service) 
	return
}