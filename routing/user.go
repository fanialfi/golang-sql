package routing

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fanialfi/golang-sql/database"
	_ "github.com/go-sql-driver/mysql"
)

func HandleUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	id := req.URL.Query().Get("id")

	if req.Method == http.MethodGet {
		data, err := queryUser(database.Driver, database.DataSource, id)
		if err != nil {
			fmt.Println(err.Error())

			message := fmt.Sprintf("data dengan id %s tidak ada di database\n", id)
			http.Error(res, message, http.StatusNotFound)
			return
		}

		dataByte, err := json.Marshal(data)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Write(dataByte)
	} else {
		http.Error(res, "", http.StatusBadRequest)
		return
	}
}

func queryUser(driverName, dataSourceName, id string) (Student, error) {
	var data = Student{}

	db, err := database.Connect(driverName, dataSourceName)
	if err != nil {
		return data, err
	}
	defer db.Close()

	// untuk query yang menghasilkan 1 baris record saja maka bisa menggunakan method QueryRow()
	// lalu bisa di chain dengan method Scan()
	err = db.
		QueryRow("SELECT * FROM tb_student WHERE id = ?", id).
		Scan(&data.Id, &data.Name, &data.Age, &data.Grade)

	if err != nil {
		return data, err
	}

	return data, err
}
