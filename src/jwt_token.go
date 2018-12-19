package main
import (
        // "../cfg"
        "fmt"
        // "net/http"
        // "github.com/gorilla/mux"
	    // "log"
        // "gopkg.in/mgo.v2"
        // "gopkg.in/mgo.v2/bson"
        // "encoding/json"
        "github.com/dgrijalva/jwt-go"

)
// var Connection *mgo.Session =cfg.Mongo("localhost:27017")
var MyJwtKey = []byte("MySecreatKey")
func GenerateJwt() (string,error)  {
        Token := jwt.New(jwt.SigningMethodHS256)
		// claims := token.SignedString(MyJwtKey)
		tokenString ,err :=Token.SignedString(MyJwtKey)
        if err !=nil {
			fmt.Errorf("Something Went Wrong : %s", err.Error())
			return "",err
		}
        return tokenString ,nil
}
func main()  {
	tokenString ,err :=GenerateJwt()
	if err !=nil {
		fmt.Errorf("Something Went Wrong")
	}
	fmt.Println(tokenString)

}