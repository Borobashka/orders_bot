package database

import (
	"database/sql"
	"fmt"
)

//===================================================================EMPLOYEE QUERIES==========================================================================//
func GettingEmployees() []Employee {

	results, err := Db.Query("SELECT * FROM employee")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	employees := []Employee{}
	for results.Next() {
		var empl Employee
		// for each row, scan into the Product struct
		err = results.Scan(&empl.Employee_id, &empl.Name, &empl.Creationdate, &empl.Exhausted, &empl.Role, &empl.Phone)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// append the product into products array
		employees = append(employees, empl)
	}

	return employees
}

func GettingEmployee(code string) *Employee {

	db, err := sql.Open("postgres", "postgres"+":"+"12345"+"@tcp(127.0.0.1:8080)/"+"orders_app")
	empl := &Employee{}
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}
    
	defer db.Close()
	//results, err := Db.Query("SELECT * FROM employee WHERE employee_id = ? ", code)

	sqlstm := fmt.Sprintf("SELECT * FROM employee WHERE employee_id = '%s'",
			code)

	results, err := Db.Query(sqlstm)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&empl.Employee_id, &empl.Name, &empl.Creationdate, &empl.Exhausted, &empl.Role, &empl.Phone)
		if err != nil {
			return nil
		}
	} else {
		return nil
	}

	return empl
}

func AddingEmployee(empl Employee) {

	db, err := sql.Open("postgres", "postgres"+":"+"12345"+"@tcp(127.0.0.1:8080)/"+"orders_app")

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	// curl http://localhost:8080/employee
	// --include
	// --header "Content-Type: application/json"
	// --request "POST"
	// --data '{"employee_id":12131, "name": "Alex","exhausted":true, "role": "dwd", "phone":"dwdw"}'

	sqlstm := fmt.Sprintf("INSERT INTO employee (employee_id, name, exhausted, role, phone) VALUES ('%d', '%s', '%t\n', '%s', '%s')",
		empl.Employee_id, empl.Name, empl.Exhausted, empl.Role, empl.Phone)

	insert, err := Db.Query(sqlstm)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
		//
	}

	defer insert.Close()
}

func UpEmployee(empl Employee) {

	db, err := sql.Open("postgres", "postgres"+":"+"12345"+"@tcp(127.0.0.1:8080)/"+"orders_app")

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	// curl http://localhost:8080/employee
	// --include
	// --header "Content-Type: application/json"
	// --request "UPDATE"
	// --data '{"employee_id":12131, "name": "Alex","exhausted":true, "role": "dwd", "phone":"dwdw"}'

	sqlstm := fmt.Sprintf("UPDATE employee SET name = '%s', exhausted = '%t\n', role='%s' , phone='%s' WHERE employee_id = '%d'",
		empl.Name, empl.Exhausted, empl.Role, empl.Phone, empl.Employee_id)

	insert, err := Db.Query(sqlstm)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

func DelEmployee(code string)  {

	db, err := sql.Open("postgres", "postgres"+":"+"12345"+"@tcp(127.0.0.1:8080)/"+"orders_app")

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()


	sqlstm := fmt.Sprintf("DELETE FROM employee WHERE employee_id = '%s'",
			code)

	insert, err := Db.Query(sqlstm)

		// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
		
	defer insert.Close()
}

//===================================================================DOCUMENT QUERIES==========================================================================//
func GettingDocuments() []Document {

	results, err := Db.Query("SELECT * FROM document")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	documents := []Document{}
	for results.Next() {
		var doc Document
		// for each row, scan into the Product struct
		err = results.Scan(&doc.Document_id, &doc.Year, &doc.Path, &doc.Name, &doc.Author, &doc.Creationdate, &doc.Employee_id)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// append the product into products array
		documents = append(documents, doc)
	}

	return documents
}

func GettingDocument(code string) *Document {

	db, err := sql.Open("postgres", "postgres"+":"+"12345"+"@tcp(127.0.0.1:8080)/"+"orders_app")
	doc := &Document{}
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}
    
	defer db.Close()

	sqlstm := fmt.Sprintf("SELECT * FROM document WHERE document_id = '%s'",
			code)

	results, err := Db.Query(sqlstm)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&doc.Document_id, &doc.Year, &doc.Path, &doc.Name, &doc.Author, &doc.Creationdate, &doc.Employee_id)
		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return doc
}

func AddingDocument(doc Document) {

	db, err := sql.Open("postgres", "postgres"+":"+"12345"+"@tcp(127.0.0.1:8080)/"+"orders_app")

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	// curl http://localhost:8080/employee
	// --include
	// --header "Content-Type: application/json"
	// --request "POST"
	// --data '{"employee_id":12131, "name": "Alex","exhausted":true, "role": "dwd", "phone":"dwdw"}'

	sqlstm := fmt.Sprintf("INSERT INTO document (year, path, name, author, employee_id) VALUES ('%d', '%s', '%s', '%s', '%d')",
		doc.Year, doc.Path, doc.Name, doc.Author, doc.Employee_id)

	insert, err := Db.Query(sqlstm)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
		//
	}

	defer insert.Close()
}

func UpDocument(doc Document) {

	db, err := sql.Open("postgres", "postgres"+":"+"12345"+"@tcp(127.0.0.1:8080)/"+"orders_app")

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	// curl http://localhost:8080/employee
	// --include
	// --header "Content-Type: application/json"
	// --request "UPDATE"
	// --data '{"employee_id":12131, "name": "Alex","exhausted":true, "role": "dwd", "phone":"dwdw"}'

	sqlstm := fmt.Sprintf("UPDATE document SET year = '%d', path = 's%s', name = '%s', author = '%s', employee_id = '%d' WHERE document_id = '%d'",
		doc.Year, doc.Path, doc.Name, doc.Author, doc.Employee_id, doc.Document_id)

	insert, err := Db.Query(sqlstm)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
}

func DelDocument(code string)  {

	db, err := sql.Open("postgres", "postgres"+":"+"12345"+"@tcp(127.0.0.1:8080)/"+"orders_app")

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()


	sqlstm := fmt.Sprintf("DELETE FROM document WHERE document_id = '%s'",
			code)

	insert, err := Db.Query(sqlstm)

		// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
		
	defer insert.Close()
}


func GettingMaxIdDocument() *Document {

	db, err := sql.Open("postgres", "postgres"+":"+"12345"+"@tcp(127.0.0.1:8080)/"+"orders_app")
	doc := &Document{}
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()

	sqlstm := fmt.Sprintf("SELECT MAX(document_id) FROM document")

	results, err := Db.Query(sqlstm)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&doc.Document_id)
		if err != nil {
			return nil
		}
	}

	return doc
}