package main

import(
	"fmt"
	"sync"
)

var wg sync.WaitGroup


func runner(s string){
	for i:=0;i<3;i++{
		fmt.Println(s)
	}
	defer wg.Done()
}


func main(){
	wg.Add(1)
	go runner("Hello")
	wg.Add(1)
	go runner("There")
	wg.Add(1)
	go runner("Testinf")
	wg.Wait()
}