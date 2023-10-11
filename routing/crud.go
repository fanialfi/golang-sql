package routing

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fanialfi/golang-sql/crud"
	"github.com/fanialfi/golang-sql/model"
)

func Crud(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		res.Header().Set("Content-Type", "application/json")
		data := model.Student{}

		defer req.Body.Close()

		err := json.NewDecoder(req.Body).Decode(&data)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		msg, err := crud.ExecInsert(data)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		responseMessage := model.Response{
			Status: http.StatusOK,
			Msg:    msg,
			Data:   data,
		}

		responseByte, err := json.Marshal(responseMessage)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Write(responseByte)
		fmt.Printf("%#v\n", data)
	default:
		http.Error(res, "bad request", http.StatusBadRequest)
		return
	}
}
