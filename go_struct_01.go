package main

import "fmt"

func main(){
	var subscriber struct {
		name string
		rate int
		actrive bool
	}
	fmt.Printf("%#v\n", subscriber)
	/*
	subscriber := map[string]float64{}
	subscriber["name"]= "kim"
	subscriber["rate"]= 5000
	subscriber["active"]= false
	subsctiber = append(subscriber, 5000)
	subsctiber = append(subscriber,false)
	fmt.Println(subscriber)
	*/

}
