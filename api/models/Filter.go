package models

import (
	"database/sql"
)

type Filter struct {
	Id          int64	`json:"id"`
	Name        string	`json:"name"`
	Status		int		`json:"status"`
	Order		int		`json:"order"`
}

var filter_id int64 = 0

func (f *Filter) GetFilters(db *sql.DB) (*[]Filter, error) {
	defer db.Close()
	var filters []Filter
    rows, err := db.Query("SELECT * FROM filters")
    if err != nil {
        return &filters, err
    }
	defer rows.Close()
	
	for rows.Next() {
		err := rows.Scan(&f.Id, &f.Name, &f.Status , &f.Order)
		if err != nil {
			continue
		}
		filters = append(filters, *f)
	}
	return &filters, nil
}


func (f *Filter) GetFilterById(db *sql.DB, filter_id int64) (error) {
	defer db.Close()
	row := db.QueryRow("SELECT * FROM filters WHERE id=$1", filter_id)
	err := row.Scan(&f.Id, &f.Name, &f.Status, &f.Order)
    return err
}

func (f *Filter) AddFilter(db *sql.DB) (int64, error) {
	defer db.Close()
	
	result, err := db.Exec("INSERT INTO filters (name, status, `order`) values ($1,$2,$3)", f.Name, f.Status, f.Order)
	if err != nil {
        return filter_id, err
    }
	filter_id, err = result.LastInsertId()
	return filter_id, err
}

func (f *Filter) EditFilter(db *sql.DB) (int64, error) {
	defer db.Close()
	_, err := db.Exec("UPDATE filters SET name = $1, status = $2, `order` = $3 WHERE id = $4", f.Name, f.Status, f.Order, f.Id)
	return f.Id, err
}

func (f *Filter) DeleteImage(db *sql.DB) (int64, error) {
	defer db.Close()
	_, err := db.Exec("DELETE FROM filters WHERE id = $1", f.Id)
	_, err = db.Exec("DELETE FROM image_to_filter WHERE filter_id = $1", f.Id)
	return f.Id, err
}

func (f *Filter) GetFiltersByImageId(db *sql.DB, image_id int64) (*[]Filter, error) {
	defer db.Close()
	var Filters []Filter
	rows, err := db.Query("SELECT filters.id, filters.name, filters.status, filters.`order` FROM filters LEFT JOIN image_to_filter ON (filters.id = image_to_filter.filter_id) WHERE  image_to_filter.image_id=$1", image_id)
	defer rows.Close()
	if err != nil {
        return &Filters, err
    }
	for rows.Next() {
		err := rows.Scan(&f.Id, &f.Name, &f.Status , &f.Order)
		if err != nil {
			continue
		}
		Filters = append(Filters, *f)
	}
	return &Filters, nil
}