package database

type Employee struct {
	Employee_id  int    `json:"employee_id"`
	Name         string `json:"name"`
	Creationdate string `json:"crteationdate"`
	Exhausted    bool   `json:"exhausted"`
	Role         string `json:"role"`
	Phone        string `json:"phone"`
}

type Document struct {
	Document_id 	int 	`json:"document_id"`
	Year 			int 	`json:"year"`
	Path 			string  `json:"path"`
	Name 			string  `json:"name"`
	Author 			string  `json:"author"`
	Creationdate 	string  `json:"creationdate"`
	Employee_id 	int 	`json:"employee_id"`
}