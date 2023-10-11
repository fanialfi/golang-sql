package routing

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fanialfi/golang-sql/crud"
	"github.com/fanialfi/golang-sql/database"
	"github.com/fanialfi/golang-sql/model"
	_ "github.com/go-sql-driver/mysql"
)

func HandleUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	id := req.URL.Query().Get("id")

	switch req.Method {
	case http.MethodGet:
		data, err := queryUser(id)
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

	case http.MethodDelete:
		res.Header().Set("Content-Type", "application/json")

		data, err := crud.ExecDelete(id)
		if err != nil {
			http.Error(res, err.Error(), http.StatusNotFound)
			return
		}

		responseMessage := model.Response{
			Status: http.StatusOK,
			Msg:    data,
			Data:   nil,
		}
		responseByte, err := json.Marshal(responseMessage)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Write(responseByte)
	default:
		http.Error(res, "", http.StatusBadRequest)
		return
	}
}

func queryUser(id string) (model.Student, error) {
	var data = model.Student{}

	db, err := database.Connect()
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
