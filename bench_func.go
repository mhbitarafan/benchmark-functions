package bench

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// g is number of goroutines, c is how many function will run in each goroutine
func BenchFunc(fn func(), g int, c int) {
	fn_name := GetFunctionName(fn)
	done := make(chan int, 1)
	fmt.Printf("%v: ", fn_name)
	start := time.Now()
	for a := 0; a < g; a++ {
		go func() {
			for b := 0; b < c; b++ {
				fn()
			}
			done <- 1
		}()
	}

	for a := 0; a < g; a++ {
		select {
		case <-done:
		}
	}
	elapsed := time.Since(start)
	// run per second
	rps := int(float64(g*c) / elapsed.Seconds())
	// time taken per function run in microseconds
	micro_op := elapsed.Microseconds() / int64(g*c)
	// function runs per goroutines
	rpg := rps / g

	fmt.Printf("%v runs per second | ", rps)
	fmt.Printf("%v µs/op | ", micro_op)
	fmt.Printf("%v runs per goroutine\n", rpg)
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
