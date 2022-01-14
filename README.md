benchmark go functions

install 
```
go get -u github.com/mhbitarafan/benchmark-functions
```

example
```
import bench "github.com/mhbitarafan/benchmark-functions"

func main() {
  // number of goroutines
  g := 10
  // function runs per goroutine
  c := 1000

  bench.BenchFunc(my_func, g, c)
}

//example output
main.my_func: 131416 runs per second | 7 µs/op | 657 runs per goroutine
```
