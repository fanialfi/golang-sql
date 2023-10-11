package crud

import (
	"fmt"

	"github.com/fanialfi/golang-sql/database"
	"github.com/fanialfi/golang-sql/model"
	_ "github.com/go-sql-driver/mysql"
)

func ExecInsert(data model.Student) (string, error) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer db.Close()

	// untuk operasi insert, update, delete direkomendasikan eksekusi perintah perintah tersebut pakek db.Exec
	// teknik Prepare juga bisa digunakan pada metode ini
	_, err = db.Exec("INSERT INTO tb_student VALUES(?, ?, ?, ?)", data.Id, data.Name, data.Age, data.Grade)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return "Insert Data Success", nil
}
