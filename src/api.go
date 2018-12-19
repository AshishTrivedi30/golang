package main
import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"


)


func MysqlDatabaseConnection(w http.ResponseWriter, r *http.Request) {			
		db,err:=sql.Open("mysql","ashish:Ashish_Trivedi123@tcp(127.0.0.1:3306)/dfoundry")
		fmt.Println(db,"######",err,nil)
		if err !=nil{
		panic(err.Error())
		}
		fmt.Println("SuccessFully Connected With Mysql Database")
		respondWithJson(w, http.StatusOK,string{"result": "success"})
}
func FindMovieEndpoint( r *http.Request) {
	fmt.Fprintln(w, "not implemented yet 2!")
}

func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet 3!")
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet 4!")
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet 5!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/MysqlConnection", MysqlDatabaseConnection).Methods("GET")
	r.HandleFunc("/movies2", CreateMovieEndPoint).Methods("GET")
	r.HandleFunc("/movies3", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies4", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies5/{id}", FindMovieEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}