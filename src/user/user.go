package user
import "composition"
import (
	"github.com/jinzhu/gorm"
	"user/entity"
)

type GormConnectionProvider interface {
	Gorm() *gorm.DB
}

type User struct {
}

func(this *User) UpdateSchema() error {
	var connectionProviders []GormConnectionProvider
	err := composition.GetInstance().Import(&connectionProviders)
	if (err != nil) {
		return err
	}
	for _, connectionProvider := range(connectionProviders) {
		db := connectionProvider.(GormConnectionProvider).Gorm()
		db.AutoMigrate(entity.User{})
	}
	return nil
}

func NewUserModule() (*User, error) {
	module := new(User)
	return module, nil
}

func init() {
	composition.GetInstance().Export(NewUserModule)
	composition.GetInstance().Export(func() (*GetUserService, error) {
		return new(GetUserService), nil
	})
}