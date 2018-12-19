package config

import (
          "fmt"
	     // "log"
	  "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "./lib"	
)

func Mongo(host string) *mgo.Collection {
        // var Connection *mgo.Collection =cfg.Mongo("localhost:27017")
        Connection ,err := mgo.Dial(host)
        if err != nil {
               return nil
        }	
        MongoDBConnection := Connection.DB("test").C("info")
        fmt.Println(MongoDBConnection) 
        fmt.Println(err)
        fmt.Println("DOnneeennene")       
        return MongoDBConnection
}

// func main() {
//      session :=Mongo("localhost:27017")
//      fmt.Println(session) 
//   } 
// func hello()  {
     
// }
