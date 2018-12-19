package main
import (
    "fmt")
type I interface {
    method1() 
}
type T struct{
    Name string
}
func mesure(i I)  {
    i.method1()
}
func (t T) method1() {
    fmt.Println(t)
}


// func (t T) method1(Name) string {
//     fmt.Println("Hello")
// }
func main() { 
    t := T{"Ashish"}
    mesure(t)
}