package main

import(
	"fmt"
	"sync"
)

func cat() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-ccat
		fmt.Println("cat")
		cdog <- true
	}
}

func dog() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-cdog
		fmt.Println("dog")
		cfish <- true
	}
}

func fish() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		<-cfish
		fmt.Println("fish")
		ccat<-true
	}
}

var wg sync.WaitGroup

var ccat = make(chan bool,1)
var cdog = make(chan bool,1)
var cfish = make(chan bool,1)

func main() {
	ccat <- true
	wg.Add(3)
	go cat()
	go dog()
	go fish()
	wg.Wait()
}
