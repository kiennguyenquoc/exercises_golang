//   1. Solve this problem, called the Fizz-Buzz [23] problem:
//   Write a program that prints the numbers from 1 to 100. But for multiples
// of three print “Fizz” instead of the number and for the multiples of five
// print “Buzz”. For numbers which are multiples of both three and five print
// “FizzBuzz”.

package main
import "fmt"

func main() {
  for i := 0; i <= 100; i++ {
    if i%3 == 0 && i%5 == 0 {
      fmt.Printf("%d: FizzBuzz\n", i)
    } else {
      if i%3 == 0 {
        fmt.Printf("%d: Fizz\n", i)
      }  else {
        fmt.Printf("%d: Buzz\n", i)
      }
    }
  }
}
