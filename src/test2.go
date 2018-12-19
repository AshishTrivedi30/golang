package main

import (
	"./lib"
	
)
func main() {
	lib.Demo()
	// fmt.Println(lib.Number)
	// fmt.Println(lib.Pi)
	// s :=lib.A{X:10}
	// fmt.Println(s.X) 
	
}

// package main

// import (

//         "./cfg"
//         "fmt"
//         // "reflect"
// 	// "log"
//         // "gopkg.in/mgo.v2"
//         // "gopkg.in/mgo.v2/bson"
//         // "./lib"
        
// )


// func main() {
//         session :=cfg.mongo("localhost:27017")
//         fmt.Println(session) 
//         //("server1.example.com,server2.example.com")
//         // if err != nil {
//         //         panic(err)
//         // }
//         // defer session.Close()
//         // fmt.Println(reflect.TypeOf(session))
//         // fmt.Println(reflect.TypeOf(err))
        
//         // Optional. Switch the session to a monotonic behavior.
        


//         // c := session.DB("test").C("people")
//         // err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},
// 	//                &Person{"Cla", "+55 53 8402 8510"})
//         // if err != nil {
//         //         log.Fatal(err)
//         // }

//         // result := Person{}
//         // err = c.Find(bson.M{"name": "Ale"}).One(&result)
//         // if err != nil {
//         //         log.Fatal(err)
//         // }

//         // fmt.Println("Phone:", result.Phone)
// }
