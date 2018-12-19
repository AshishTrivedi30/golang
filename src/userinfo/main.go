package main
import(
	"net/http"
	"github.com/gorilla/mux"
	"../controller"
)
func main() {
                fmt.Println("MAin")
                mux := mux.NewRouter()
                // info :=Authentication{"12121"}
                // Data(info)
                // mux.HandleFunc("/userLogout",controller.userLogout)
                mux.HandleFunc("/userLogin", controller.userLogin)
                mux.HandleFunc("/userRegistration",controller.userRegistration)
                // r.HandleFunc("/movies5/{id}", FindMovieEndpoint).Methods("GET")
                http.ListenAndServe(":3000", mux)  
}  

