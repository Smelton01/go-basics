package main

//Maps array slices loops
import "fmt"

func main() {
	//for loop
	for i := 0; i < 2; i++ {
		fmt.Println("hope we got it right")
	}
	//while loop
	j := 1
	for j <= 5 {
		if j == 3 {
			print(j)
		} else {
			print("*")
		}
		j++
	}
	fmt.Println("")
	//arrays
	nums := []int{0, 1, 4, 9}
	fmt.Println(nums[2:])
	for i, val := range nums {
		fmt.Println("[", i, ":", val, "]")
	}
	//slice :=
	fmt.Println(append(nums[:], 16, 25))
	//maps
	dick := make(map[string]int)
	dick["mary"] = 6
	dick["natalia"] = 9
	fmt.Println(dick["natalia"])
	fmt.Println(dick)

	Titans := map[string]map[string]string{
		"Nightwing": map[string]string{
			"Name":        "Dicky boy",
			"Description": "FKA Robin",
		},
		"Raven": map[string]string{
			"Name":        "Dont even remember",
			"Description": "Weird looking chiq with the cool powers",
		},
	}
	fmt.Println(Titans)
}
