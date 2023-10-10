package database

import (
	"database/sql"
)

const (
	Driver     = "mysql"
	DataSource = "root:@tcp(127.0.0.1:3306)/db_belajar_golang"
)

// function ini digunakan untuk membuat / membuka koneksi ke database
func Connect(driverName, dataSourceName string) (*sql.DB, error) {
	// sql.Open() digunakan untuk memulai koneksi dengan database
	// didalamnya terdapat 2 parameter mandatory (wajib) => nama driver dan koneksi string
	//
	// untuk skema koneksi string setap driver bisa saja berbeda
	// disini saya menggunakan driver mysql
	// untuk skema koneksi stringnya "root:@tcp(127.0.0.1:3306)/db_belajar_golang"
	//
	// user:password@tcp(host/url:port)/db_name
	// user@tcp(host/url:port)/db_name
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
