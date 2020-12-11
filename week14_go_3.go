package main

import(
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)
func main(){
	resp, e := http.Get("https://www.inhatc.ac.kr")
	if e != nil {
		log.Fatal(e)
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(len(body))
}
