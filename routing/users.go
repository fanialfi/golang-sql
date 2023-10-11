package routing

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fanialfi/golang-sql/database"
	"github.com/fanialfi/golang-sql/model"

	// driver di import namun perlu ditambahkan _
	// karena meskipun di butuhkan package "database/sql", namun kita tidak langsung berinteraksi dengan driver tersebut
	_ "github.com/go-sql-driver/mysql"
)

func HandleUsers(res http.ResponseWriter, req *http.Request) {
	data, err := queryUsers()
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(result)
}

// function ini digunakan untuk melakukan query ke database di DBMS
func queryUsers() ([]model.Student, error) {
	db, err := database.Connect()
	if err != nil {
		fmt.Printf("line 39 %s", err.Error())
		return nil, err
	}
	defer db.Close()

	// db.Query() digunakan untuk eksekusi sql query
	// parameter kedua dari function ini adalah variadic, jadi bisa di isi atau tidak
	// jika pada argument pertama di ada tanda ?, maka parameter kedua harus diisi sebanyak tanda ? ada pada argument pertama
	// nantinya tanda ? akan ter replace oleh argument setelahnya
	// teknik penulisan query ini dianjurkan untuk mencegah sql injection
	rows, err := db.Query("SELECT * FROM tb_student")
	if err != nil {
		fmt.Printf("line 46 %s", err.Error())
		return nil, err
	}

	// function db.Query() menghasilkan instance bertipe *sql.Rows
	// yang juga perlu do close saat sudah tidak digunakan
	defer rows.Close()

	// digunakan untuk menampung hasil query
	var result []model.Student

	// perulangan dilakukan sebanyak berapa record yang berhasil di query
	// perulangan dengan kondisi acuan rows.Next() ini dilakukan sebanyak jumplah total record yang ada
	for rows.Next() {
		each := model.Student{}

		// setelah di iterasi tiap tiap rows,
		// method rows.Scan() digunakan untuk mengambil nilai record yang sedang diiterasi dan kemudian disimpan kedalam variabel poiter
		// variabel yang digunakan untuk menyimpan nilai dari field field record di tuliskan sebagai parameter variadic
		// sesuai dengan field yang di select pada query,
		// jika select query menggunakan tanda asterisk (*), maka urutan penulisan parameter-nya sama dengan struktur yang ada di database
		//
		// query
		// select id, name, grade, ...
		//
		// Scan
		// &each.id, &each.name, &each.grade
		err = rows.Scan(&each.Id, &each.Name, &each.Age, &each.Grade)
		if err != nil {
			fmt.Printf("line 57 %s", err.Error())
			return nil, err
		}

		// data record yang didapatkan kemudian di append ke slice result
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("line 65 %s", err.Error())
		return nil, err
	}

	return result, nil
}
