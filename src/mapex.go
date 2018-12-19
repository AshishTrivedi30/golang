package main
import(
	"fmt"
)
func main()  {
	statePopulation := make(map[string]int)
	statePopulation =map[string] int{
		"California" : 3987543,
		"Texas" : 4356543,
		"New York" : 77655435,

	}

	fmt.Println(statePopulation)
	statePopulation["ohio"]=24356
	fmt.Println(statePopulation)
	delete(statePopulation,"California")
	fmt.Println(statePopulation["ohoo"])
}