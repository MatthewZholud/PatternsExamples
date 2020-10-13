////with goroutine
//
//package main
//
//import (
//	"fmt"
//	"sync"
//)
//
//
//
//
//var lock = &sync.Mutex{}
//
//type single struct {
//}
//
//var singleInstance *single
//
//var message string
//
//func getInstance() *single {
//	if singleInstance == nil {
//		lock.Lock()
//		defer lock.Unlock()
//		if singleInstance == nil{
//			fmt.Println("Creating single instance now.")
//			singleInstance = &single{}
//		} else {
//			fmt.Println("Single instance already created.")
//		}
//	} else {
//		fmt.Println("Single instance already created.")
//	}
//
//	return singleInstance
//}
//
//func main() {
//
//	for i := 0; i < 30; i++ {
//		go getInstance()
//	}
//	fmt.Scanln()
//}
//












// without goroutine
// better for understanding

package main

import (
	"fmt"
	"sync"
)

type singleton map[string]string

var lock = &sync.Mutex{}

var instance singleton

func New() singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			fmt.Println("Creating single instance now.")
			instance = make(singleton)
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return instance
}



func main() {

	s := New()

	s["this"] = "that"

	s2 := New()

	s2["1"] = "2"

	fmt.Println("This is ", s["1"])
	fmt.Println("This is ", s2["this"])
}