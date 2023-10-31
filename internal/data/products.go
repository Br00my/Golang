package data

import (
	"fmt"
	"time"
	"errors"
	"database/sql"
	"context"
	_ "github.com/lib/pq"
	"greenworld.wow/internal/validator"
)

type ProductModel struct {
	DB *sql.DB
}

type MockProductModel struct{}
func (m MockProductModel) Insert(Product *Product) error {
	return nil
}
func (m MockProductModel) Get(id int64) (*Product, error) {
	return nil, nil
}
func (m MockProductModel) Update(Product *Product) error {
	return nil
}
func (m MockProductModel) Delete(id int64) error {
	return nil
}
func (m MockProductModel) GetAll(title string, category string, filters Filters) ([]*Product, error) {
	return nil, nil
}

type Product struct {
	ID int64 `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title string `json:"title"`
	Price int8 `json:"price"`
	Category string `json:"category"`
}

func ValidateProduct(v *validator.Validator, product *Product) {
	v.Check(product.Title != "", "title", "must be provided")
	v.Check(len(product.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(product.Category == "Reusable Water Bottles" || product.Category == "Bamboo Toothbrushes" || product.Category == "Solar Charges" || product.Category == "Eco-Friendly Cleaning Products" || product.Category == "Recycled Paper Products", "category", "is invalid")
	v.Check(product.Price > 0, "price", "must be provided")
}

func (m ProductModel) Insert(product *Product) error {
	query := `
		INSERT INTO products (title, category, price)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`

	args := []interface{}{product.Title, product.Category, product.Price}

	return m.DB.QueryRow(query, args...).Scan(&product.ID, &product.CreatedAt)
}

func (m ProductModel) Get(id int64) (*Product, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT id, created_at, title, category, price
		FROM products
		WHERE id = $1`

	var product Product
	err := m.DB.QueryRow(query, id).Scan(
		&product.ID,
		&product.CreatedAt,
		&product.Title,
		&product.Category,
		&product.Price,
	)

	if err != nil {
		switch {
			case errors.Is(err, sql.ErrNoRows):
				return nil, ErrRecordNotFound
			default:
				return nil, err
		}
	}

	return &product, nil
}

func (m ProductModel) Update(product *Product) error {
	query := `
		UPDATE products
		SET title = $1, category = $2, price = $3
		WHERE id = $4
		RETURNING $1`

	args := []interface{}{
		product.Title,
		product.Category,
		product.Price,
		product.ID,
	}

	return m.DB.QueryRow(query, args...).Scan(&product.Title)
}

func (m ProductModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
		DELETE FROM products
		WHERE id = $1`

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (m ProductModel) GetAll(title string, category string, filters Filters) ([]*Product, error) {
	query := fmt.Sprintf(`
		SELECT id, created_at, title, category, price
		FROM products
		WHERE (to_tsvector('simple', title) @@ plainto_tsquery('simple', $1) OR $1 = '') and
		(to_tsvector('simple', category) @@ plainto_tsquery('simple', $2) OR $2 = '')
		ORDER BY %s %s, id ASC
		LIMIT $3 OFFSET $4`, filters.sortColumn(), filters.sortDirection())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []interface{}{title, category, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []*Product{}

	for rows.Next() {

		var product Product

		err := rows.Scan(
			&product.ID,
			&product.CreatedAt,
			&product.Title,
			&product.Category,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, &product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
