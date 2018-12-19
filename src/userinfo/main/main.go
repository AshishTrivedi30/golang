package main
import(
	"net/http"
	"github.com/gorilla/mux"
    "../controller"
    "fmt"
)
func main() {
                fmt.Println("MAin")
                mux := mux.NewRouter()
                // info :=Authentication{"12121"}
                // Data(info)
                // mux.HandleFunc("/userLogout",controller.userLogout)
                
                mux.HandleFunc("/userRegistration",controller.UserRegistration)
                mux.HandleFunc("/userLogin", controller.UserLogin)
                mux.HandleFunc("/userUpdate", controller.ForgetPassword)
                // r.HandleFunc("/movies5/{id}", FindMovieEndpoint).Methods("GET")
                http.ListenAndServe(":3000", mux)  
}  

