package main

import (
	"fmt"
	"strconv"
)

func GetOneCust() {
	id, err := strconv.Atoi(Input("Enter ID to search:"))
	checkForError(err)
	db:=GetDb()
	defer db.Close()

	sql:="SELECT NAME, CITY, EMAIL FROM CUSTOMERS WHERE ID=?"
	stmt,err:=db.Prepare(sql)
	checkForError(err)
	defer stmt.Close()

	row:=stmt.QueryRow(id)
	var name,city,email string
	if err:=row.Scan(&name,&city,&email);err!=nil{
		fmt.Println("No data found for id",id)
	}else{
		fmt.Printf("Name:%s\n",name)
		fmt.Printf("City:%s\n",city)
		fmt.Printf("email:%s\n",email)
		fmt.Println()
	}
}