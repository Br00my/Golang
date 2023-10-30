package data
import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict = errors.New("edit conflict")
)

type Models struct {
	Products interface {
		Insert(product *Product) error
		Get(id int64) (*Product, error)
		Update(product *Product) error
		Delete(id int64) error
		GetAll(title string, category string, filters Filters) ([]*Product, error)
	}
}

func NewModels(db *sql.DB) Models {
	return Models{
		Products: ProductModel{DB: db},
	}
}

func NewMockModels() Models {
	return Models{
		Products: MockProductModel{},
	}
}
