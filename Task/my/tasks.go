package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/lib/pq"
// )

// //student data
// type student struct {
// 	fName string
// 	lName string
// }

// //main function
// func main() {

// 	connStr := "user=postgres dbname=test password=1234 host=localhost sslmode=disable"

// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		fmt.Print(err)
// 	}

// 	//sql to insert student information
// 	statement := "INSERT INTO student(fname, lname) VALUES($1, $2)"
// 	//prepare statement for sql
// 	stmt, err := db.Prepare(statement)
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	defer stmt.Close()

// 	eName := student{}

// 	for i := 0; i < 3; i++ {
// 		fmt.Print("First Name: ")
// 		fmt.Scanf("%s", &eName.fName)

// 		fmt.Print("Last Name: ")
// 		fmt.Scanf("%s", &eName.lName)

// 		stmt.QueryRow(eName.fName, eName.lName)
// 	}

// 	rows, err := db.Query("Select fname, lname from student")
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	defer rows.Close()

// 	fmt.Println("---------------------------------------------------------------------")

// 	for rows.Next() {

// 		var fname string
// 		var lname string
// 		err := rows.Scan(&fname, &lname)
// 		if err != nil {
// 			fmt.Print(err)
// 		}

// 		fmt.Printf("%s %s\n", fname, lname)
// 	}
// }
