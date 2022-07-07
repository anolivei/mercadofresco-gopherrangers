package locality

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	GetAll(ctx context.Context) ([]Locality, error)
	GetById(ctx context.Context, id string) (Locality, error)
	Create(ctx context.Context, id string, localityName, provinceName, countryName string) (Locality, error)
	ReportSellers(ctx context.Context, id string) (ReportSeller, error)
}

type mariaDBRepository struct {
	db *sql.DB
}

func NewMariaDBRepository(db *sql.DB) Repository {
	return &mariaDBRepository{db: db}
}

func (m mariaDBRepository) ReportSellers(ctx context.Context, id string) (ReportSeller, error) {
	var reportSeller ReportSeller

	rows, err := m.db.QueryContext(ctx, "SELECT l.id, l.locality_name, COUNT(seller.id) FROM localities l LEFT JOIN seller ON l.id  = seller.locality_id")

	if err != nil {
		return ReportSeller{}, err
	}

	defer rows.Close()

	for rows.Next() {

		rows.Scan(&reportSeller.LocalityID, &reportSeller.LocalityName, &reportSeller.SellersCount)

		if err != nil {
			return ReportSeller{}, err
		}
	}
	return reportSeller, nil
}

func (m mariaDBRepository) Create(ctx context.Context, id string, localityName, provinceName, countryName string) (Locality, error) {

	locality := Locality{Id: id, LocalityName: localityName, ProvinceName: provinceName, CountryName: countryName}

	queryInsert := "INSERT INTO localities (id, locality_name, province_name, country_name) VALUES (?,?,?,?)"

	_, err := m.db.ExecContext(ctx, queryInsert, &locality.Id, &locality.LocalityName, &locality.ProvinceName, &locality.CountryName)

	if err != nil {
		return Locality{}, err
	}

	return locality, nil

}

func (m mariaDBRepository) GetAll(ctx context.Context) ([]Locality, error) {
	var localityList []Locality

	rows, err := m.db.QueryContext(ctx, "SELECT * FROM localities")

	if err != nil {
		return localityList, err
	}

	defer rows.Close()

	for rows.Next() {
		var locality Locality

		err = rows.Scan(&locality.Id, &locality.LocalityName, &locality.ProvinceName, &locality.CountryName)

		if err != nil {
			return localityList, err
		}

		localityList = append(localityList, locality)
	}

	return localityList, err
}

func (m mariaDBRepository) GetById(ctx context.Context, id string) (Locality, error) {
	var locality Locality

	rows, err := m.db.QueryContext(ctx, "SELECT * FROM localities WHERE id = ?", id)

	if err != nil {
		return locality, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&locality.Id, &locality.LocalityName, &locality.ProvinceName, &locality.CountryName)

		if err != nil {
			return locality, err
		}

		return locality, nil
	}

	if err := rows.Err(); err != nil {
		return locality, err
	}

	return locality, fmt.Errorf("id does not exists")
}
