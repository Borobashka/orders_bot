package database

// import "time"

type Employee struct {
	Employee_id  int    `json:"employee_id"`
	Name         string `json:"name"`
	Creationdate string `json:"crteationdate"`
	Exhausted    bool    `json:"exhausted"`
	Role         string `json:"role"`
	Phone        string `json:"phone"`
}

// type Document struct {
// 	document_id 	int
// 	year 			int
// 	path 			string
// 	name 			string
// 	author 			string
// 	creationdate 	time.Time
// 	employee_id 	int
// }