package user
import "reflect"
import "github.com/darkengines/composition"
import (
	"github.com/jinzhu/gorm"
	"github.com/darkengines/user/entity"
)

type GormConnectionProvider interface {
	Gorm() gorm.DB
}

type user struct {
}

func(this * user) UpdateSchema() {
	connectionProviders := composition.GetExportedValues(reflect.TypeOf((*GormConnectionProvider)(nil)).Elem())
	for _, connectionProvider := range(connectionProviders) {
		db := connectionProvider.(GormConnectionProvider).Gorm()
		db.AutoMigrate(entity.User{})
	}
}

func init() {
	composition.Export(reflect.TypeOf((*user)(nil)).Elem())
}