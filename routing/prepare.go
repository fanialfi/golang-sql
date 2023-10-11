package routing

import (
	"fmt"
	"net/http"

	"github.com/fanialfi/golang-sql/database"
	_ "github.com/go-sql-driver/mysql"
)

func Prepare(res http.ResponseWriter, req *http.Request) {
	sqlPrepare()
}

// teknik prepare adalah teknik penulisan query di awal dengan kelebihan bisa di re-use digunakan berkali kali
// method ini bisa digabung / dichain dengan method Query() atau QueryRow()
func sqlPrepare() {
	db, err := database.Connect(database.Driver, database.DataSource)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// method Prepare digunakan untuk deklarasi query, object yang dikembalikan berupa *sql.Stmt
	stmt, err := db.Prepare("SELECT name, grade FROM tb_student WHERE id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// dari object stmt method QueryRow dipanggil beberapa kali dengan paameter adalah id untuk nanatinya dikirim
	// sebagai pelengkap query pada db.Prepare

	result1 := Student{}
	stmt.QueryRow("B001").Scan(&result1.Name, &result1.Grade)
	fmt.Printf("name: %s\ngrade: %d\n", result1.Name, result1.Grade)
	fmt.Printf("%#v\n", result1)
	defer stmt.Close()

	result2 := Student{}
	stmt.QueryRow("B003").Scan(&result2.Name, &result2.Grade)
	fmt.Printf("name: %s\ngrade: %d\n", result2.Name, result2.Grade)
	fmt.Printf("%#v\n", result2)

	result3 := Student{}
	stmt.QueryRow("B003").Scan(&result3.Name, &result3.Grade)
	fmt.Printf("name: %s\ngrade: %d\n", result3.Name, result3.Grade)
	fmt.Printf("%#v\n", result3)
}
