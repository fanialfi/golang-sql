package model

// skema struct ini disiapkan sama seperti skema pada table tb_student yang ada di database
// yang digunakan sebagai tipe data penampung hasil query
type Student struct {
	Id    string `json:"id"`
	Name  string `json:"nama"`
	Age   int    `json:"umur"`
	Grade int    `json:"kelas"`
}
