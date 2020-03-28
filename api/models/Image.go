package models

import (
	"database/sql"
)

type Image struct {
	Id          int64			`json:"id"`
	Name        string			`json:"name"`
	Status		int				`json:"status"`
	Order		int				`json:"order"`
	Filters		*[]Filter		`json:"filters"`
}
var image_id int64 = 0
var err error = nil

func (i *Image) GetImages(db *sql.DB) (*[]Image, error) {
	defer db.Close()
	var images []Image
    rows, err := db.Query("SELECT  id, name, status, `order` FROM images")
    if err != nil {
        return &images, err
    }
	defer rows.Close()
	
	for rows.Next() {
		err := rows.Scan(&i.Id, &i.Name, &i.Status , &i.Order)
		if err != nil {
			continue
		}
		images = append(images, *i)
	}
	return &images, nil
}


func (i *Image) GetImageById(db *sql.DB, image_id int64) (error) {
	defer db.Close()
	row := db.QueryRow("SELECT  id, name, status, `order` FROM images WHERE id=$1  ORDER BY created_at DESC", image_id)
	err := row.Scan(&i.Id, &i.Name, &i.Status, &i.Order)
	if err != nil {
        return err
    }

	filter := Filter{}
	filters, err := filter.GetFiltersByImageId(db, image_id)
	if err != nil {
        return err
    }

	i.Filters = filters
	return nil
}

func (i *Image) AddImage(db *sql.DB) (int64, error) {
	defer db.Close()
	
	result, err := db.Exec("INSERT INTO images (name, status, `order`, created_at, updated_at) values ($1, $2, $3, DateTime('now'), DateTime('now'))", i.Name, i.Status, i.Order)
	if err != nil {
        return image_id, err
    }
	image_id, err = result.LastInsertId()
	if err != nil {
        return image_id, err
	}
	if i.Filters != nil {
		err = addImageToFilter(image_id, i.Filters, db)
	}
	return image_id, err
}

func (i *Image) EditImage(db *sql.DB) (int64, error) {
	defer db.Close()
	_, err := db.Exec("UPDATE images SET name = $1, status = $2, `order` = $3, updated_at = DateTime('now') WHERE id = $4", i.Name, i.Status, i.Order, i.Id)
	if err != nil {
        return i.Id, err
	}
	if i.Filters != nil {
		err = addImageToFilter(i.Id, i.Filters, db)
	}
	return i.Id, err
}

func (i *Image) PatchImage(db *sql.DB) (int64, error) {
	defer db.Close()
	_, err = db.Exec("UPDATE images SET status = $1, updated_at = DateTime('now') WHERE id = $2", i.Status, i.Id)
	return i.Id, err
}

func (i *Image) DeleteImage(db *sql.DB) (int64, error) {
	defer db.Close()
	_, err := db.Exec("DELETE FROM images WHERE id = $1", i.Id)
	_, err = db.Exec("DELETE FROM image_to_filter WHERE image_id = $1", i.Id)
	return i.Id, err
}
func addImageToFilter(image_id int64, f *[]Filter, db *sql.DB) error{
	_, err := db.Exec("DELETE FROM image_to_filter WHERE image_id = $1", image_id)
	for _, val := range *f {
		_, err = db.Exec("INSERT INTO image_to_filter (image_id, filter_id) values ($1,$2)", image_id, val.Id)
	}
	return err
}