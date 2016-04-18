package databasee
import "composition"
import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/sqlite"

type Database struct {
	gorm *gorm.DB
}

func(this *Database) Gorm() *gorm.DB {
	return this.gorm
}

func NewDatabase() (*Database, error) {
	database := new(Database)
	db, err := gorm.Open("sqlite3", "sqlite3.db")
	if (err != nil) {
		return nil, err
	}
	database.gorm = db
	return database, nil
}

func init() {
	composition.GetInstance().Export(NewDatabase)
}