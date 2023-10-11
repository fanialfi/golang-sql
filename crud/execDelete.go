package crud

import (
	"fmt"

	"github.com/fanialfi/golang-sql/database"
	_ "github.com/go-sql-driver/mysql"
)

func ExecDelete(id string) (string, error) {
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM tb_student WHERE id = ?", id)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return "", fmt.Errorf("data dengan id %s tidak ada di database", id)
	}

	return "delete data success", nil
}
