package main

import "fmt"

func AcceptAndAdd() {
	db := GetDb()
	defer db.Close()
	sql:="INSERT INTO CUSTOMERS (NAME,CITY,EMAIL) VALUES (?,?,?)" //First round to db;Server will keep a compiled copy of this
	stmt,err:=db.Prepare(sql)
	checkForError(err)
	defer stmt.Close()
	for {
		fmt.Println("Enter customer Details:")
		name:=Input("Name:")
		city:=Input("City:")
		email:=Input("Email:")
		result,err:=stmt.Exec(name,city,email)//subsequent roundtrips to db; server executes the compiled sql with these values
		checkForError(err)
		newId,err:=result.LastInsertId()
		checkForError(err)
		fmt.Printf("New Customer data added successfully with id %v\n",newId)
		choice:=Input("Want yo add another? (yes/no):")
		if choice=="no"{
			break

		}
	}
}