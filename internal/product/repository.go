package products

import (
	"database/sql"
	"fmt"
)

const (
	GETALL  = "SELECT * FROM products"
	GETBYID = "SELECT * FROM products WHERE id=?"
	STORE   = `INSERT INTO products (product_code, description,
				width, height, length, net_weight, expiration_rate,
				recommended_freezing_temperature, freezing_rate,
				product_type_id, seller_id)
				VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	UPDATE = `UPDATE products SET
				product_code=?, description=?, width=?, height=?,
				length=?, net_weight=?, expiration_rate=?,
				recommended_freezing_temperature=?, freezing_rate=?,
				product_type_id=?, seller_id=?
				WHERE id=?`
	DELETE  = "DELETE FROM products WHERE id=?"
	LAST_ID = "SELECT MAX(id) as last_id FROM products"
	PRODUCT_CODE = `SELECT product_code FROM products
					WHERE id != ? and product_code = ?`
)

type Product struct {
	ID                             int     `json:"id"`
	ProductCode                    string  `json:"product_code" validate:"required"`
	Description                    string  `json:"description" validate:"required"`
	Width                          float64 `json:"width" validate:"required,gt=0"`
	Height                         float64 `json:"height" validate:"required,gt=0"`
	Length                         float64 `json:"length" validate:"required,gt=0"`
	NetWeight                      float64 `json:"net_weight" validate:"required,gt=0"`
	ExpirationRate                 string  `json:"expiration_rate" validate:"required"`
	RecommendedFreezingTemperature float64 `json:"recommended_freezing_temperature" validate:"required,gt=0"`
	FreezingRate                   float64 `json:"freezing_rate" validate:"required,gt=0"`
	ProductTypeId                  int     `json:"product_type_id" validate:"required,gt=0"`
	SellerId                       int     `json:"seller_id"`
}

type Repository interface {
	Store(prod Product) (Product, error)
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Update(prod Product, id int) (Product, error)
	Delete(id int) error
	CheckProductCode(id int, productCode string) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Store(prod Product) (Product, error) {
	stmt, err := r.db.Prepare(STORE)
	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(&prod.ProductCode, &prod.Description,
		&prod.Width, &prod.Height, &prod.Length, &prod.NetWeight,
		&prod.ExpirationRate, &prod.RecommendedFreezingTemperature,
		&prod.FreezingRate, &prod.ProductTypeId, &prod.SellerId)
	if err != nil {
		return Product{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return Product{}, fmt.Errorf("fail to save")
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return Product{}, err
	}
	prod.ID = int(lastId)
	return prod, nil
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	rows, err := r.db.Query(GETALL)
	if err != nil {
		return ps, err
	}
	defer rows.Close()
	for rows.Next() {
		var prod Product
		err := rows.Scan(&prod.ID, &prod.ProductCode, &prod.Description,
			&prod.Width, &prod.Height, &prod.Length, &prod.NetWeight,
			&prod.ExpirationRate, &prod.RecommendedFreezingTemperature,
			&prod.FreezingRate, &prod.ProductTypeId, &prod.SellerId)
		if err != nil {
			return ps, err
		}
		ps = append(ps, prod)
	}
	return ps, nil
}

func (r *repository) GetById(id int) (Product, error) {
	var prod Product
	stmt, err := r.db.Prepare(GETBYID)
	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&prod.ID, &prod.ProductCode, &prod.Description,
		&prod.Width, &prod.Height, &prod.Length, &prod.NetWeight,
		&prod.ExpirationRate, &prod.RecommendedFreezingTemperature,
		&prod.FreezingRate, &prod.ProductTypeId, &prod.SellerId)
	if err != nil {
		return Product{}, fmt.Errorf("product %d not found", id)
	}
	return prod, nil
}

func (r *repository) Update(prod Product, id int) (Product, error) {
	stmt, err := r.db.Prepare(UPDATE)
	if err != nil {
		return Product{}, err
	}
	defer stmt.Close()
	olderProduct, _ := r.GetById(id)
	result, err := stmt.Exec(&prod.ProductCode, &prod.Description,
		&prod.Width, &prod.Height, &prod.Length, &prod.NetWeight,
		&prod.ExpirationRate, &prod.RecommendedFreezingTemperature,
		&prod.FreezingRate, &prod.ProductTypeId, &prod.SellerId, id)
	if err != nil {
		return Product{}, err
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Println(olderProduct)
	fmt.Println(prod)
	if rowsAffected == 0 && prod != olderProduct {
		return Product{}, fmt.Errorf("product %d not found", id)
	}
	return prod, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare(DELETE)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("product %d not found", id)
	}
	return nil
}

func (r *repository) CheckProductCode(id int, productCode string) bool {
	stmt, err := r.db.Prepare(PRODUCT_CODE)
	if err != nil {
		return false
	}
	defer stmt.Close()
	err = stmt.QueryRow(id, productCode).Scan(&productCode)
	return err != nil
}
