package main

import (
	"fmt"
	"time"
	// "runtime"
	// "time"
)

// func read(ch <-chan int) {
// 	fmt.Println("read")
// 	for {
// 		v, ok := <-ch
// 		fmt.Printf("val: %d ok:%v\n", v, ok)
// 		if !ok {
// 			return
// 		}
// 	}
// }

// func write(ch chan<- int) {
// 	fmt.Println("write")
// 	for {
// 		ch <- int(time.Now().Nanosecond())
// 	}
// }

func printA(a int) {
	fmt.Println(a)
}

func main() {
	// ch := make(chan int)
	// chw := make(chan int)
	// go read(ch)
	// go write(chw)
	// // defer close(ch)

	// cnt := 0
	// for {
	// 	select {
	// 	case ch <- int(time.Now().Nanosecond()):
	// 		// fmt.Printf("write %d\n", cnt)
	// 		// time.Sleep(time.Second * 2)
	// 	case v, ok := <-chw:
	// 		fmt.Printf("read v: %d ok: %v\n", v, ok)
	// 	default:
	// 		fmt.Println("default")
	// 		if cnt > 10 {
	// 			return
	// 		}
	// 		runtime.Gosched()
	// 		cnt++
	// 		// time.Sleep(time.Second)
	// 	}
	// }
	a := 1
	// go func(a int) {
	// 	fmt.Println(a)
	// }(a)
	go printA(a)
	a = 2
	time.Sleep(time.Second)
	fmt.Println("end")

}
