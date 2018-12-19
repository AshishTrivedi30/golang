package model

import (
		"gopkg.in/mgo.v2"
		// "../info"
		"fmt"
        // "net/http"
        // "github.com/gorilla/mux"
		//  "log"
		"userinfo/config"
        // "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
		// "encoding/json"
        "golang.org/x/crypto/bcrypt"

        // "github.com/dgrijalva/jwt-go"
)

var MongoDBConnection *mgo.Collection =config.Mongo("localhost:27017")

type UserInfo struct {
	Name ,Password ,Role ,Email string 
	Phone ,User_id int 
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}


func InsertUserData(info interface{}) error {
		err :=MongoDBConnection.Insert(info)
		// fmt.Println(result)
		if err !=nil {
				return err   
		}       
		return err     
}

func LoginUser(id int) string {
		result :=UserInfo{}
		fmt.Println(id)      
        MongoDBConnection.Find(bson.M{"user_id" :id}).One(&result)
        fmt.Println(result.Password)
		return result.Password
}

func UpdateUserPassword(id int ,password string) error {
	result :=UserInfo{}
	// result2 :=UserInfo{}
	err :=MongoDBConnection.Find(bson.M{"user_id" :id}).One(&result)
	if err !=nil {
		return err 
	}         
	Hash_Password, _ :=HashPassword(password)
	// fmt.Println(info)
	// fmt.Println(result2)
	MongoDBConnection.Update(bson.M{"password" : result.Password} ,
	bson.M{"$set": bson.M{"password":Hash_Password}})

	// MongoDBConnection.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"Name": "new Name"}}// if err !=nil {
	// 	fmt.Println(err)  
	// }	
	fmt.Println("Success")
	return err
}