package controller
// import "github.com/dchest/passwordreset"

import (
        "../model"
        // "../cfg"
        "fmt"
        "net/http"
	// "log"
        // "gopkg.in/mgo.v2"
        // "gopkg.in/mgo.v2/bson"
        "encoding/json"
        "github.com/dgrijalva/jwt-go"
        // "api.jwt.auth/services"
        // "api.jwt.auth/services/models"
        // "../info"
        
        "golang.org/x/crypto/bcrypt"

)
// var HOST,PORT string
// var MongoDBConnection *mgo.Collection =cfg.Mongo("localhost:27017")

// type User interface{
//         // userRegistration() error
//         userLogin()
// }  
type Authentication struct{
        User_id int
        Password  string 
        // Email string      
}

type UserInfo struct {
	Name ,Password ,Role ,Email string 
	Phone ,User_id int 
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func UserRegistration(w http.ResponseWriter, r *http.Request) {
        Info := UserInfo{}
        err :=json.NewDecoder(r.Body).Decode(&Info)
        if err !=nil {
                panic(err)                
        } 
        Hash_Password, _ :=HashPassword(Info.Password)
        Info.Password = Hash_Password
        err = model.InsertUserData(Info)
        if err !=nil {
                panic(err)                
        } 
        // fmt.Println(result)
        w.WriteHeader(200)
        w.Header().Set("Content-Type","application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Successfully Registerd"))              
}


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
func UserLogin(w http.ResponseWriter, r *http.Request) {
        UserData := Authentication{}
        err :=json.NewDecoder(r.Body).Decode(&UserData)
        if err !=nil {
                panic(err)                
        }
        // Hash_Password, _ :=HashPassword(UserData.Password)
        result :=model.LoginUser(UserData.User_id)
        User_Test := CheckPasswordHash(UserData.Password,result)
        var Jwt_User string
        if User_Test ==true{
                Jwt_User ,err =GenerateJwt()
                if err !=nil {
                        panic(err)                
                }
                w.WriteHeader(200)
                w.Header().Set("Content-Type","application/json")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte(Jwt_User))  
        } else{
                w.WriteHeader(404)
                w.Header().Set("Content-Type","application/json") 
                w.Write([]byte("You Are not Authentic User"))    
        }
}

func ForgetPassword(w http.ResponseWriter, r *http.Request)  {
        Info := UserInfo{}
        err :=json.NewDecoder(r.Body).Decode(&Info) 
        if err !=nil {
                panic(err)                
        }
        err =model.UpdateUserPassword(Info.User_id,Info.Password)
        if err !=nil {
                w.WriteHeader(404)
                w.Header().Set("Content-Type","application/json") 
                w.Write([]byte("You Are not Authentic User"))            
        }else{
                w.WriteHeader(200)
                w.Header().Set("Content-Type","application/json")
                w.WriteHeader(http.StatusOK)
                w.Write([]byte("SuccessFully Password Updated"))
        }  
}








// func userLogout(w http.ResponseWriter, r *http.Request)  {
//         err := services.Logout(r)
//         w.Header().Set("Content-Type", "application/json")
//         if err != nil {
//         w.WriteHeader(http.StatusInternalServerError)
//         } else {
//         w.WriteHeader(http.StatusOK)
//         }        
// }







