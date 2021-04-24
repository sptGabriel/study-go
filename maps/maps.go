package maps

import "fmt"

func Test() {
	score := make(map[string]int)
	score["gabriel"] = 1
	score["gabriel2"] = 1
	score["gabriel3"] = 1
	score["gabriel4"] = 1
	score["gabriel5"] = 1
	score["gabriel6"] = 1
	score["gabriel7"] = 1
	for k, v := range score {
		fmt.Printf("Score of %v is %v \n", k, v)
	}
	val, is_present := score["gabriel"]
	fmt.Println("Value:", val, "Has Value:", is_present)
	fmt.Println("Map length:", len(score))
}
