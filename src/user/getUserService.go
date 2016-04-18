package user

import "composition"
import "user/entity"

type GetUserService struct {
}

func(this *GetUserService) Parameter() interface{} {
	return new(uint)
}

func(this *GetUserService) Process(id interface{}) (interface{}, error) {
	var connectionProviders []GormConnectionProvider
	err := composition.GetInstance().
	Import(&connectionProviders)
	if (err != nil) {
		return nil, err
	}
	result := new(entity.User)
	filter := entity.User{}
	filter.ID = *(id.(*uint))
	connectionProviders[0].(GormConnectionProvider).Gorm().Find(result, filter)
	return result, nil
}

func(this *GetUserService) Name() string {
	return "GetUserService"
}

func init() {
	composition.GetInstance().Export(func() (*GetUserService, error) {
		return new(GetUserService), nil
	})
}