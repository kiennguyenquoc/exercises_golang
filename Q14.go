package main
import "fmt"

func bubblesort(n []int) {
  for i := 0; i < len(n) - 1; i++ {
    for j := i + 1; j < len(n); j++ {
      if n[j] < n[i] {
        n[i], n[j] = n[j], n[i]
      }
    }
  }
}

func prtthem(numbers... int){
  for _, d := range numbers {
    fmt.Printf("%d\n",d)
  }
}

func fibonacci(value int) []float64 {
  x := make([]float64, value)
  x[0], x[1] = 1, 1
  for n := 2; n < value; n++ {
    x[n] = x[n-1] + x[n-2]
  }
  return x
}

func main() {
  n := []int{1, 54, 32, 4, -4, 45, 54353, 435, 2, 44}
  fmt.Printf("unsorted %v\n", n)
  bubblesort(n)
  fmt.Printf("sorted: %v\n", n)
  prtthem(1, 4, 5, 7, 4)
  prtthem(1, 2, 4)
  for _, i := range fibonacci(10) {
    fmt.Printf("%v ", i)
  }
  var aa []float64
  aa = fibonacci(10)
  fmt.Printf("%v", aa)
}
