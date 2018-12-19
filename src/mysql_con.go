package main
import "database/sql"
import _"github.com/go-sql-driver/mysql"
import "fmt" 
func main(){
			
			db,err:=sql.Open("mysql","ashish:Ashish_Trivedi123@tcp(127.0.0.1:3306)/dfoundry")
			fmt.Println(db,"######",err,nil)
			// if err=nil{
			// 		fmt.Println("Done")
			// }
			


if err !=nil{
	panic(err.Error())
}

// // fmt.println("Go MySql Tutorial")
rows, err := db.Query("SELECT * FROM employee")	
// insert,err:=db.Query("select * from employee");
columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
values := make([]sql.RawBytes, len(columns))	
for rows.Next() {
	// get RawBytes from data
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	err = rows.Scan(scanArgs...)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Now do something with the data.
	// Here we just print each column as a string.
	var value string
	for i, col := range values {
		fmt.Print(columns[i],columns[1])
		fmt.Println(value)
		// Here we can check if the value is nil (NULL value)
		if col == nil {
			value = "NULL"
		} else {
			value = string(col)
		}
		
	}
	
}
if err !=nil{
	panic(err.Error())
}

defer db.Close()

// fmt.println("SuccessFully Connected")
}