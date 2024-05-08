package database

import (
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

	empl := &Employee{}

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

	sqlstm := fmt.Sprintf("INSERT INTO employee (employee_id, name, exhausted, role, phone) VALUES ('%d', '%s', '%t\n', '%s', '%s')",
		empl.Employee_id, empl.Name, empl.Exhausted, empl.Role, empl.Phone)

	_,err := Db.Query(sqlstm)
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
}

func UpEmployee(empl Employee) {

	sqlstm := fmt.Sprintf("UPDATE employee SET name = '%s', exhausted = '%t\n', role='%s' , phone='%s' WHERE employee_id = '%d'",
		empl.Name, empl.Exhausted, empl.Role, empl.Phone, empl.Employee_id)

	_, err := Db.Query(sqlstm)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

}

func DelEmployee(code string)  {

	sqlstm := fmt.Sprintf("DELETE FROM employee WHERE employee_id = '%s'",
			code)

	_, err := Db.Query(sqlstm)

		// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
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

func GettingDocumentsYear(year string) []Document {

	sqlstm := fmt.Sprintf("SELECT * FROM document WHERE year = '%s'", year)
	results, err := Db.Query(sqlstm)

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

func GettingMaxIdDocument(year string) *Document {

	doc := &Document{}

	sqlstm := fmt.Sprintf("SELECT MAX(document_id) FROM document WHERE year = '%s'", year)

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

func GettingDocument(code string, year string) *Document {

	doc := &Document{}

	sqlstm := fmt.Sprintf("SELECT * FROM document WHERE document_id = '%s' AND year = '%s'",
			code, year)

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

	var maxDocumentNumber int
	sqlstm := fmt.Sprintf("SELECT MAX(document_id) FROM document WHERE year = '%d'", doc.Year)

	results, err := Db.Query(sqlstm)

	if err != nil {
		fmt.Println("Err", err.Error())
	}

	if results.Next() {
		err = results.Scan(&maxDocumentNumber)
		if err != nil {
			fmt.Println("success", err.Error())
		}
	}

	// Увеличиваем номер документа на 1
	doc.Document_id = maxDocumentNumber + 1
	fmt.Println(doc)
	sqlstm = fmt.Sprintf("INSERT INTO document (document_id, year, path, name, author, employee_id) VALUES ('%d', '%d', '%s', '%s', '%s', '%d')",
		doc.Document_id, doc.Year, doc.Path, doc.Name, doc.Author, doc.Employee_id)

	_, err = Db.Query(sqlstm)
	
	fmt.Println(sqlstm)
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
		//
	}
}


func UpDocument(doc Document) {

	sqlstm := fmt.Sprintf("UPDATE document SET  path = '%s', name = '%s', author = '%s', employee_id = '%d' WHERE document_id = '%d' AND year = '%d'",
		doc.Path, doc.Name, doc.Author, doc.Employee_id, doc.Document_id, doc.Year)

	_, err := Db.Query(sqlstm)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
}

func DelDocument(code, year string)  {

	sqlstm := fmt.Sprintf("DELETE FROM document WHERE document_id = '%s' AND year = '%s'",
			code, year)

	_ , err := Db.Query(sqlstm)

		// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
		
}


