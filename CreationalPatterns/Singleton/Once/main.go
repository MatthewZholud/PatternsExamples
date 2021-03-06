package main


import (
	"fmt"
	"sync"
)

type singleton map[string]string



var (
once sync.Once

instance singleton
)

func  New() singleton {
once.Do(func() {
instance = make(singleton)
})

return instance
}



func main() {

s := New()

s["this"] = "that"

s2 := New()

s2["1"]  = "2"

fmt.Println("This is ", s["1"])
fmt.Println("This is ", s2["this"])
// This is that

}


