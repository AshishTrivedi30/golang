package main
import(
	"fmt"
)
type Doctor struct {
	Number int
	ActorName string
	Array [] string
	
}
func  main()  {
	aDoctor :=Doctor {
		Number : 3,
		ActorName :"Salman",
		Array :[]string {
			"Liz Shaw",
			"Jo Grant",
			"Sarah",
		},
	}
	fmt.Println(aDoctor.Array[1])
}