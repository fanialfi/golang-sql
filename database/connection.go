package database

import (
	"database/sql"
)

const (
	driver     = "mysql"
	dataSource = "root:@tcp(127.0.0.1:3306)/db_belajar_golang"
)

// function ini digunakan untuk membuat / membuka koneksi ke database
func Connect() (*sql.DB, error) {
	// sql.Open() digunakan untuk memulai koneksi dengan database
	// didalamnya terdapat 2 parameter mandatory (wajib) => nama driver dan koneksi string
	//
	// untuk skema koneksi string setap driver bisa saja berbeda
	// disini saya menggunakan driver mysql
	// untuk skema koneksi stringnya "root:@tcp(127.0.0.1:3306)/db_belajar_golang"
	//
	// user:password@tcp(host/url:port)/db_name
	// user@tcp(host/url:port)/db_name
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
