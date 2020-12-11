package main

import "fmt"

func main(){
	floatSlice := make([]float64, 10)
	boolSlice := make([]bool, 10)
	fmt.Println(floatSlice[8], boolSlice[6])
	myMusic := []string{"aa", "bb", "cc", "dd"}
	mySlice := myMusic[:]
	mySlice[3] = "Sun"
	fmt.Printf("%x\n", &mySlice[3])
	fmt.Printf("%x\n", &myMusic[3])
	myMusic = append(myMusic, "ee")
	fmt.Println(myMusic)
	fmt.Println(mySlice)
	fmt.Printf("%x\n", &mySlice[3])
	fmt.Printf("%x\n", &myMusic[3])

}
