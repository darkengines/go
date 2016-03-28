package database
import "github.com/darkengines/composition"
import "github.com/jinzhu/gorm"
import "reflect"

type Database struct {
	gorm gorm.DB
}

func(this *Database) Gorm() *gorm.DB {
	return this.gorm
}

func(this *Database) Initialize() error {
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	if (err != nil) {
		return err
	}
	this.gorm = db
	return nil
}

func init() {
	composition.Export(reflect.TypeOf((*Database{})(nil)).Elem())
}